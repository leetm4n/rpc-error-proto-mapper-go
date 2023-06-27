package main

import (
	rpcerrormapperv1 "github.com/leetm4n/rpc-error-proto-mapper-go/api/proto/rpc/errormapper/v1"
	"github.com/leetm4n/rpc-error-proto-mapper-go/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func generateFile(gen *protogen.Plugin, protoFile *protogen.File) {
	generationTargetEnums := getGenerationTargetEnums(protoFile.Enums)
	// if no Enums are generation target then skip
	if len(generationTargetEnums) == 0 {
		return
	}

	filename := protoFile.GeneratedFilenamePrefix + ".pb.rpcerrormapper.go"
	generatedFile := gen.NewGeneratedFile(filename, protoFile.GoImportPath)

	generator := generator.New(
		generationTargetEnums,
		protoFile,
		generatedFile,
	)

	generator.Generate()
}

func getGenerationTargetEnums(enums []*protogen.Enum) []*protogen.Enum {
	targetEnums := []*protogen.Enum{}

	for _, enum := range enums {
		if isEnumGenerationTarget(enum) {
			targetEnums = append(targetEnums, enum)
		}
	}

	return targetEnums
}

func isEnumGenerationTarget(enum *protogen.Enum) bool {
	isTarget, ok := proto.GetExtension(
		enum.Desc.Options(),
		rpcerrormapperv1.E_IsGenerationTarget.TypeDescriptor().Type(),
	).(bool)

	if !ok {
		return false
	}

	return isTarget
}
