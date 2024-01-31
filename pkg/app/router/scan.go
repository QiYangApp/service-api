package router

//https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

import (
	"crypto/md5"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// Names controller Top-level interface used to declare a structure as a controller.
type Names struct {
	Name          string
	FullPackName  string
	ShortPackName string
}

// MethodInfo Api method info
type MethodInfo struct {
	ApiMethodName  string // API method。such as: POST、GET、DELETE、PUT、OPTIONS、PATCH、HEAD
	ApiPath        string // API path
	PackPath       string // API pack
	PackName       string
	PackMethodName string            //
	Annotations    map[string]string // Annotations of the method
}

type Scan struct {
	Apis        map[string]*MethodInfo
	Pkg         string
	ApiRootPath string
	BasePaths   map[string]string
	Names       map[string]Names
	Controller  []Inject
}

func (s *Scan) recursionPkgDir(currentDir, fullPackName, shortPckName string) {

	dirs, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("[%s] read path error, %s", currentDir, err.Error())
	}

	pkgs, err := decorator.ParseDir(token.NewFileSet(), currentDir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("[%s] parse controller directory error, %s", currentDir, err.Error())
	}

	s.recursionPkgFile(pkgs, currentDir, fullPackName, shortPckName)

	for _, dir := range dirs {
		if dir.IsDir() {
			nextDir, err := filepath.Abs(path.Join(currentDir, dir.Name()))
			nextFullPackName := fullPackName + "/" + dir.Name()
			if err != nil {
				log.Fatalf("[%s] red controller directory error, %s", nextDir, err.Error())
			}

			s.recursionPkgDir(nextDir, nextFullPackName, dir.Name())
		}
	}
}

func (s *Scan) recursionPkgFile(pkgs map[string]*dst.Package, currentDir, fullPackName, shortPackName string) {
	commonMethodRegex := regexp.MustCompile("//.*@(BasePath|RootPath|GET|POST|DELETE|PATCH|HEAD|PUT|OPTIONS|ANY).*")
	commonAttrRegex := regexp.MustCompile("[(|,]([^=]+)=\"([^\"]+)[)|\"]")

	for _, p := range pkgs {
		for filePath, file := range p.Files {
			dst.Inspect(file, func(node dst.Node) bool {
				switch t := node.(type) {
				case *dst.GenDecl:
					var match bool
					var structType *dst.StructType
					var spec *dst.TypeSpec
					if spec, match = t.Specs[0].(*dst.TypeSpec); !match {
						return false
					}
					if structType, match = spec.Type.(*dst.StructType); !match {
						return false
					}
					if s.isController(structType.Fields.List) {
						key := fmt.Sprintf("pack_%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s", fullPackName, shortPackName, spec.Name.Name))))
						s.Names[key] = Names{Name: spec.Name.Name, ShortPackName: shortPackName, FullPackName: fullPackName}

						var prefix string
						for _, comment := range t.Decs.Start {
							commentsMethod := commonMethodRegex.FindStringSubmatch(comment)
							if len(commentsMethod) == 0 {
								prefix = "/"
							} else {
								commentsAttr := commonAttrRegex.FindStringSubmatch(commentsMethod[0])
								if len(commentsAttr) == 0 {
									log.Fatalf("[%s]: invalid api definition, example: @GET(path=\"/test\")", currentDir)
								}

								if commentsMethod[1] == "RootPath" && s.ApiRootPath != "" {
									log.Fatalf("api root path repat set, [%s] %s:", s.ApiRootPath, commentsMethod[1])
								}

								if commentsMethod[1] == "BasePath" {
									prefix = commentsAttr[2]
								} else if commentsMethod[1] == "RootPath" {
									s.ApiRootPath = commentsAttr[2]
								}
							}
						}
						s.BasePaths[spec.Name.Name] = prefix
					}
				case *dst.FuncDecl:
					if t.Decs.Start == nil || t.Recv == nil || t.Name.Name == "PostConstruct" {
						return false
					}
					onwer := s.searchFather(t.Recv.List) // Which controller does it belong to
					method := MethodInfo{
						Annotations:    make(map[string]string),
						PackPath:       fullPackName,
						PackName:       shortPackName,
						PackMethodName: onwer,
					}
					for _, comment := range t.Decs.Start {
						commentsMethod := commonMethodRegex.FindStringSubmatch(comment)
						if len(commentsMethod) != 0 {
							if unicode.IsLower(rune(t.Name.Name[0])) {
								log.Fatalf("[%s] %s: invalid method name, name first word must be uppercase", filePath, t.Name.Name)
							}

							method.ApiMethodName = commentsMethod[1]
							commentsAttr := commonAttrRegex.FindAllStringSubmatch(commentsMethod[0], 1)
							if len(commentsAttr) == 0 {
								log.Fatalf("[%s] %s invalid path parameter, Must start with path=", filePath, t.Name.Name)
							}

							for _, attr := range commentsAttr {
								switch attr[1] {
								case "path":
									method.ApiPath = path.Join(s.BasePaths[onwer], attr[2])
									break
								}
							}

							if method.ApiPath == "" {
								log.Fatalf("[%s] %s: invalid api definition, example: @GET(path=\"/test\")", currentDir, t.Name.Name)
							}

							continue
						}

						if strings.HasPrefix(comment, "@") {
							annotationArr := strings.Split(comment, "->")
							var annotationVal string
							if len(annotationArr) == 2 {
								annotationVal = annotationArr[1]
							}
							method.Annotations[annotationArr[0]] = annotationVal
						}
					}
					if method.ApiPath != "" {
						key := fmt.Sprintf("%x", md5.Sum([]byte(fullPackName+"-"+onwer+"-"+t.Name.Name)))
						s.Apis[key] = &method
					}
				}
				return true
			})

		}
	}
}

// Determines whether the current structure is a controller
func (s *Scan) isController(fields []*dst.Field) bool {
	var ok bool
	var selectorExpr *dst.SelectorExpr
	for _, field := range fields {
		selectorExpr, ok = field.Type.(*dst.SelectorExpr)
		if !ok {
			continue
		}
		x := selectorExpr.X.(*dst.Ident)
		sel := selectorExpr.Sel
		if x.Name == "inject" && sel.Name == "Controller" {
			ok = true
			break
		}
	}
	return ok
}

// Query the controller to which the method belongs
func (s *Scan) searchFather(fields []*dst.Field) string {
	for _, field := range fields {
		if f, ok := field.Type.(*dst.StarExpr); ok {
			return f.X.(*dst.Ident).Name
		}
	}

	return ""
}
