package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func main() {
	protogen.Options{
		ParamFunc: flag.Set,
		ImportRewriteFunc: func(path protogen.GoImportPath) protogen.GoImportPath {
			return path
		},
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = SupportedFeatures

		for _, file := range gen.Files {
			// if file does not need to be generated or missing enums then skip
			if !file.Generate || len(file.Enums) == 0 {
				continue
			}

			// else generate file
			generateFile(gen, file)
		}

		return nil
	})
}
