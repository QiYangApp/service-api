package main

//https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/jennifer/jen"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"service-api/src/core/inject"
	"strings"
	"unicode"
)

type Container struct {
	ApiRootPath string
	Apis        map[string]*inject.MethodInfo
	BasePaths   map[string]string
	Names       []string
	Pkg         string
}

// Parse project controllers and API methods
// Run with go generate
func main() {
	controllerDir := "./src/api/controller"
	fullPackName := "service-api/src/api/controller"
	shortPckName := "controller"
	if len(os.Args) > 1 {
		controllerDir = os.Args[1]
	}
	controllerAbs, err := filepath.Abs(controllerDir)
	if err != nil {
		log.Fatalf("[%s] get controller directory abstract path error, %s", controllerDir, err.Error())
	}

	container := Container{
		BasePaths: make(map[string]string),
		Names:     make([]string, 0),
		Apis:      make(map[string]*inject.MethodInfo),
	}

	recursionPkgDir(container, controllerAbs, fullPackName, shortPckName)

	recordProjectControllerAndApi(container, controllerAbs, fullPackName)
}

func recursionPkgDir(container Container, currentDir, fullPackName, shortPckName string) {

	dirs, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("[%s] read path error, %s", currentDir, err.Error())
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			nextDir, err := filepath.Abs(path.Join(currentDir, dir.Name()))
			nextFullPackName := fullPackName + "/" + dir.Name()
			if err != nil {
				log.Fatalf("[%s] red controller directory error, %s", nextDir, err.Error())
			}

			recursionPkgDir(container, nextDir, nextFullPackName, dir.Name())
		}
	}

	pkgs, err := decorator.ParseDir(token.NewFileSet(), currentDir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("[%s] parse controller directory error, %s", currentDir, err.Error())
	}

	recursionPkgFile(pkgs, container, currentDir, fullPackName, shortPckName)
}

func recursionPkgFile(pkgs map[string]*dst.Package, container Container, currentDir, fullPackName, shortPackName string) {
	commonMethodRegex := regexp.MustCompile("//.*@(BasePath|RootPath|GET|POST|DELETE|PATCH|HEAD|PUT|OPTIONS).*")
	commonAttrRegex := regexp.MustCompile("[(|,]([^=]+)=\"([^\"]+)[)|\"]")
	haser := md5.New()

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
					if isController(structType.Fields.List) {
						container.Names = append(container.Names, spec.Name.Name)
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

								if commentsMethod[1] == "RootPath" && container.ApiRootPath != "" {
									log.Fatalf("api root path repat set, [%s] %s:", container.ApiRootPath, commentsMethod[1])
								}

								if commentsMethod[1] == "BasePath" {
									prefix = commentsAttr[2]
								} else if commentsMethod[1] == "RootPath" {
									container.ApiRootPath = commentsAttr[2]
								}
							}
						}
						container.BasePaths[spec.Name.Name] = prefix
					}
				case *dst.FuncDecl:
					if t.Decs.Start == nil || t.Recv == nil || t.Name.Name == "PostConstruct" {
						return false
					}
					onwer := searchFather(t.Recv.List) // Which controller does it belong to
					method := inject.MethodInfo{
						Annotations: make(map[string]string),
						PackPath:    fullPackName,
						PackName:    shortPackName,
						ServiceName: onwer,
					}
					for _, comment := range t.Decs.Start {
						commentsMethod := commonMethodRegex.FindStringSubmatch(comment)
						if len(commentsMethod) != 0 {
							if unicode.IsLower(rune(t.Name.Name[0])) {
								log.Fatalf("[%s] %s: invalid method name, name first word must be uppercase", filePath, t.Name.Name)
							}

							method.MethodName = commentsMethod[1]
							commentsAttr := commonAttrRegex.FindAllStringSubmatch(commentsMethod[0], 1)
							if len(commentsAttr) == 0 {
								log.Fatalf("[%s] %s invalid path parameter, Must start with path=", filePath, t.Name.Name)
							}

							for _, attr := range commentsAttr {
								switch attr[1] {
								case "path":
									method.ApiPath = path.Join(container.ApiRootPath, container.BasePaths[onwer], attr[2])
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
						haser.Write([]byte(fullPackName + "/" + onwer + "/" + t.Name.Name))
						key := hex.EncodeToString(haser.Sum(nil))
						container.Apis[key] = &method
					}
				}
				return true
			})

		}
	}
}

// Determines whether the current structure is a controller
func isController(fields []*dst.Field) bool {
	var ok bool
	var selectorExpr *dst.SelectorExpr
	for _, field := range fields {
		selectorExpr, ok = field.Type.(*dst.SelectorExpr)
		if !ok {
			continue
		}
		x := selectorExpr.X.(*dst.Ident)
		sel := selectorExpr.Sel
		if x.Name == "mvc" && sel.Name == "Controller" {
			ok = true
			break
		}
	}
	return ok
}

// Query the controller to which the method belongs
func searchFather(fields []*dst.Field) string {
	for _, field := range fields {
		if f, ok := field.Type.(*dst.StarExpr); ok {
			return f.X.(*dst.Ident).Name
		}
	}
	return ""
}

// Remove comment prefixes such as // @BasePath("/") -> @BasePath("/")
func removePrefix(text string) string {
	text = strings.ReplaceAll(text, " ", "")
	if strings.HasPrefix(text, "//") {
		return text[2:]
	}
	return text
}

// All controller information and Api information for the current project is recorded here
func recordProjectControllerAndApi(container Container, dir, controllerPackName string) {
	if len(container.Apis) == 0 {
		return
	}
	newFile := jen.NewFile("controller")
	newFile.HeaderComment("// ⚠️⛔ Auto generate code by gin-plus framework, Do not edit!!!")
	newFile.HeaderComment("// All controller information and Api information for the current project is recorded here\n")
	newFile.ImportName("service-api/src/core/inject", "inject")

	var registerCode []jen.Code
	for _, controllerName := range container.Names {
		registerCode = append(registerCode, jen.Id(fmt.Sprintf("&%s{}", controllerName)))
	}
	newFile.Func().Id("init").Params().Block(
		jen.Qual("service-api/src/core/inject", "Register").Call(registerCode...),
		jen.Qual("service-api/src/core/inject", "Apis").Op("=").Map(jen.String()).Op("*").
			Qual("service-api/src/core/inject", "MethodInfo").
			Values(jen.DictFunc(func(dict jen.Dict) {
				for k, methodInfo := range container.Apis {
					dict[jen.Lit(k)] = jen.Block(jen.Dict{
						jen.Id("PackName"):    jen.Lit(methodInfo.PackName),
						jen.Id("PackPath"):    jen.Lit(methodInfo.PackPath),
						jen.Id("ServiceName"): jen.Lit(methodInfo.ServiceName),
						jen.Id("MethodName"):  jen.Lit(methodInfo.MethodName),
						jen.Id("ApiPath"):     jen.Lit(methodInfo.ApiPath),
						jen.Id("Annotations"): jen.Map(jen.String()).String().Values(jen.DictFunc(func(dict jen.Dict) {
							for k, v := range methodInfo.Annotations {
								dict[jen.Lit(k)] = jen.Lit(v)
							}
						})),
					})
				}
			})),
	)
	err := newFile.Save(filepath.Join(dir, "controller_init.go"))
	if err != nil {
		panic(err)
	}

}
