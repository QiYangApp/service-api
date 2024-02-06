package router

//https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

import (
	"github.com/dave/jennifer/jen"
	"path"
	"path/filepath"
)

// GenerateApi All controller information and Api information for the current project is recorded here
func GenerateApi(scan *Scan, storageDir string) {
	if len(scan.Apis) == 0 {
		return
	}
	newFile := jen.NewFile("main")
	newFile.HeaderComment("// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!")
	newFile.HeaderComment("// All controller information and Api information for the current project is recorded here\n")
	newFile.ImportName("app/router", "router")

	for key, pkg := range scan.Names {
		newFile.ImportAlias(pkg.FullPackName, key)
	}

	var registerCode []jen.Code
	for _, pkg := range scan.Names {
		registerCode = append(registerCode, jen.Op("&").Qual(pkg.FullPackName, pkg.Name).Values())
	}
	newFile.Func().Id("init").Params().Block(
		jen.Qual("app/router", "Register").Call(registerCode...),
		jen.Qual("app/router", "Apis").Op("=").Map(jen.String()).Op("*").
			Qual("app/router", "MethodInfo").
			Values(jen.DictFunc(func(dict jen.Dict) {
				for k, methodInfo := range scan.Apis {
					dict[jen.Lit(k)] = jen.Block(jen.Dict{
						jen.Id("PackName"):       jen.Lit(methodInfo.PackName),
						jen.Id("PackPath"):       jen.Lit(methodInfo.PackPath),
						jen.Id("PackMethodName"): jen.Lit(methodInfo.PackMethodName),
						jen.Id("ApiMethodName"):  jen.Lit(methodInfo.ApiMethodName),
						jen.Id("ApiPath"):        jen.Lit(path.Join(scan.ApiRootPath, methodInfo.ApiPath)),
						jen.Id("Annotations"): jen.Map(jen.String()).String().Values(jen.DictFunc(func(dict jen.Dict) {
							for k, v := range methodInfo.Annotations {
								dict[jen.Lit(k)] = jen.Lit(v)
							}
						})),
					})
				}
			})),
	)

	err := newFile.Save(filepath.Join(storageDir, "init.go"))
	if err != nil {
		panic(err)
	}

}
