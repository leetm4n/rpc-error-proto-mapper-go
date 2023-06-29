package generator

import (
	"fmt"
	"strings"

	rpcerrormapperv1 "github.com/leetm4n/rpc-error-proto-mapper-go/api/proto/rpc/errormapper/v1"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type generator struct {
	enums     []*protogen.Enum
	protoFile *protogen.File
	gen       *protogen.GeneratedFile
}

func (g *generator) Generate() {
	// this is the outline of generated code
	g.addHeaderComment()
	g.addEmptyLine()
	g.addPackageName()
	g.addEmptyLine()
	g.addImports(
		[]string{
			"fmt",
			"google.golang.org/genproto/googleapis/rpc/status",
			"google.golang.org/genproto/googleapis/rpc/code",
			"google.golang.org/genproto/googleapis/rpc/errdetails",
			"google.golang.org/protobuf/proto",
			"google.golang.org/protobuf/types/known/anypb",
		},
		[]string{},
	)
	g.addEncoderFunction()
	g.addEmptyLine()
	g.addDecoderFunction()
	g.addEmptyLine()
	g.addMappers()
	g.addEmptyLine()
}

func (g *generator) addImports(imports []string, initializationImports []string) {
	g.gen.P("import (")

	for _, importValue := range imports {
		g.gen.P(fmt.Sprintf("  \"%s\"", importValue))
	}

	for _, initinitializationImport := range initializationImports {
		g.gen.P(fmt.Sprintf("  _ \"%s\"", initinitializationImport))

	}

	g.gen.P(")")
}

func (g *generator) addPackageName() {
	g.gen.P(fmt.Sprintf("package %s", g.protoFile.GoPackageName))
}

func (g *generator) addMappers() {
	for _, enum := range g.enums {
		g.addErrors(enum)
		g.addConstants(enum)
		g.addEmptyLine()
		g.addDecoderMapper(enum)
		g.addEmptyLine()
		g.addEncoders(enum)
		g.addEmptyLine()
	}
}

func (g *generator) addConstants(enum *protogen.Enum) {
	g.gen.P("const (")

	for _, value := range enum.Values {
		nameWithoutPrefix := stripEnumValuePrefix(value.GoIdent.GoName, enum.GoIdent.GoName)
		name := uppperSnakeCaseToPascalCase(nameWithoutPrefix)
		options := getOptions(value)

		if options == nil {
			continue
		}

		g.gen.P(fmt.Sprintf("  %sCode = %d", name, options.Code))
		g.gen.P(fmt.Sprintf("  %sDomain = \"%s\"", name, options.Domain))
		g.gen.P(fmt.Sprintf("  %sReason = \"%s\"", name, nameWithoutPrefix))
	}

	g.gen.P(")")
}

func getOptions(value *protogen.EnumValue) *rpcerrormapperv1.EnumValueLevelOptions {
	options, ok := proto.GetExtension(
		value.Desc.Options(),
		rpcerrormapperv1.E_Options.TypeDescriptor().Type(),
	).(*rpcerrormapperv1.EnumValueLevelOptions)

	if !ok {
		return nil
	}

	return options
}

func (g *generator) addErrors(enum *protogen.Enum) {
	for _, value := range enum.Values {
		options := getOptions(value)

		if options == nil {
			continue
		}

		name := uppperSnakeCaseToPascalCase(stripEnumValuePrefix(value.GoIdent.GoName, enum.GoIdent.GoName))

		g.gen.P(fmt.Sprintf("type %s struct{", name))
		g.gen.P("  code code.Code")
		g.gen.P("  errorInfoDetails *errdetails.ErrorInfo")
		g.gen.P("  debugInfoDetails *errdetails.DebugInfo")

		switch options.Code {
		case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
			g.gen.P("  retryInfoDetails *errdetails.RetryInfo")

		case int32(code.Code_INVALID_ARGUMENT):
			g.gen.P("  badRequestDetails *errdetails.BadRequest")

		case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
			g.gen.P("  resourceInfoDetails *errdetails.ResourceInfo")

		case int32(code.Code_RESOURCE_EXHAUSTED):
			g.gen.P("  quotaFailureDetails *errdetails.QuotaFailure")

		case int32(code.Code_FAILED_PRECONDITION):
			g.gen.P("  preconditionFailureDetails *errdetails.PreconditionFailure")

		default:

		}

		g.gen.P("}")

		g.addEmptyLine()

		g.gen.P(fmt.Sprintf("func (e %s) Error() string {", name))
		g.gen.P("  return fmt.Sprintf(\"%s: %s\", e.code, e.errorInfoDetails.Reason)")
		g.gen.P("}")

		g.addEmptyLine()

		g.gen.P(fmt.Sprintf("func (e %s) Code() code.Code {", name))
		g.gen.P("  return e.code")
		g.gen.P("}")

		g.addEmptyLine()

		g.gen.P(fmt.Sprintf("func (e %s) ErrorInfoDetails() *errdetails.ErrorInfo {", name))
		g.gen.P("  return e.errorInfoDetails")
		g.gen.P("}")

		g.addEmptyLine()

		switch options.Code {
		case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
			g.gen.P(fmt.Sprintf("func (e %s) RetryInfoDetails() *errdetails.RetryInfo {", name))
			g.gen.P("  return e.retryInfoDetails")
			g.gen.P("}")

			g.addEmptyLine()

		case int32(code.Code_INVALID_ARGUMENT):
			g.gen.P(fmt.Sprintf("func (e %s) BadRequestDetails() *errdetails.BadRequest {", name))
			g.gen.P("  return e.badRequestDetails")
			g.gen.P("}")

			g.addEmptyLine()

		case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
			g.gen.P(fmt.Sprintf("func (e %s) ResourceInfoDetails() *errdetails.ResourceInfo {", name))
			g.gen.P("  return e.resourceInfoDetails")
			g.gen.P("}")

			g.addEmptyLine()

		case int32(code.Code_RESOURCE_EXHAUSTED):
			g.gen.P(fmt.Sprintf("func (e %s) QuotaFailureDetails() *errdetails.QuotaFailure {", name))
			g.gen.P("  return e.quotaFailureDetails")
			g.gen.P("}")

			g.addEmptyLine()

		case int32(code.Code_FAILED_PRECONDITION):
			g.gen.P(fmt.Sprintf("func (e %s) PreconditionFailureDetails() *errdetails.PreconditionFailure {", name))
			g.gen.P("  return e.preconditionFailureDetails")
			g.gen.P("}")

			g.addEmptyLine()

		default:

		}

		g.gen.P(fmt.Sprintf("func (e %s) DebugInfoDetails() *errdetails.DebugInfo {", name))
		g.gen.P("  return e.debugInfoDetails")
		g.gen.P("}")

		g.addEmptyLine()

		g.gen.P(fmt.Sprintf("func New%s(", name))
		g.gen.P("  code code.Code,")
		g.gen.P("  errorInfoDetails *errdetails.ErrorInfo,")
		switch options.Code {
		case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
			g.gen.P("  retryInfoDetails: *errdetails.RetryInfo,")
		case int32(code.Code_INVALID_ARGUMENT):
			g.gen.P("  badRequestDetails *errdetails.BadRequest,")
		case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
			g.gen.P("  resourceInfoDetails *errdetails.ResourceInfo,")
		case int32(code.Code_RESOURCE_EXHAUSTED):
			g.gen.P("  quotaFailureDetails *errdetails.QuotaFailure,")
		case int32(code.Code_FAILED_PRECONDITION):
			g.gen.P("  preconditionFailureDetails *errdetails.PreconditionFailure,")
		default:
		}
		g.gen.P("  debugInfoDetails *errdetails.DebugInfo,")
		g.gen.P(fmt.Sprintf(") %s {", name))
		g.gen.P(fmt.Sprintf("  return %s{", name))
		g.gen.P("    code: code,")
		g.gen.P("    errorInfoDetails: errorInfoDetails,")
		switch options.Code {
		case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
			g.gen.P("    retryInfoDetails: retryInfoDetails,")
		case int32(code.Code_INVALID_ARGUMENT):
			g.gen.P("  badRequestDetails: badRequestDetails,")
		case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
			g.gen.P("  resourceInfoDetails: resourceInfoDetails,")
		case int32(code.Code_RESOURCE_EXHAUSTED):
			g.gen.P("  quotaFailureDetails: quotaFailureDetails,")
		case int32(code.Code_FAILED_PRECONDITION):
			g.gen.P("  preconditionFailureDetails: preconditionFailureDetails,")
		default:
		}
		g.gen.P("    debugInfoDetails: debugInfoDetails,")
		g.gen.P("  }")
		g.gen.P("}")
	}
}

func (g *generator) addEncoderFunction() {
	g.gen.P("func encode(code code.Code, details []*anypb.Any) ([]byte, error) {")
	g.gen.P("  marshalled, err := proto.Marshal(&status.Status{")
	g.gen.P("    Code: int32(code),")
	g.gen.P("    Details: details,")
	g.gen.P("    Message: \"error\",")
	g.gen.P("  })")
	g.gen.P("  if err != nil {")
	g.gen.P("    return nil, err")
	g.gen.P("  }")
	g.gen.P("  return marshalled, nil")
	g.gen.P("}")
}

func (g *generator) addDecoderFunction() {
	g.gen.P("type details struct {")
	g.gen.P("  Code code.Code")
	g.gen.P("  ErrorInfoDetails *errdetails.ErrorInfo")
	g.gen.P("  RetryInfoDetails *errdetails.RetryInfo")
	g.gen.P("  BadRequestDetails *errdetails.BadRequest")
	g.gen.P("  DebugInfoDetails *errdetails.DebugInfo")
	g.gen.P("  QuotaFailureDetails *errdetails.QuotaFailure")
	g.gen.P("  PreconditionFailureDetails *errdetails.PreconditionFailure")
	g.gen.P("  ResourceInfoDetails *errdetails.ResourceInfo")
	g.gen.P("}")

	g.addEmptyLine()

	g.gen.P("func decode(payload []byte) (*details, error) {")
	g.gen.P("  s := &status.Status{}")
	g.gen.P("  err := proto.Unmarshal(payload, s)")
	g.gen.P("  if err != nil {")
	g.gen.P("    return nil, err")
	g.gen.P("  }")
	g.gen.P("  var errorInfoDetails *errdetails.ErrorInfo")
	g.gen.P("  var badRequestDetails *errdetails.BadRequest")
	g.gen.P("  var debugInfoDetails *errdetails.DebugInfo")
	g.gen.P("  var quotaFailureDetails *errdetails.QuotaFailure")
	g.gen.P("  var preconditionFailureDetails *errdetails.PreconditionFailure")
	g.gen.P("  var resourceInfoDetails *errdetails.ResourceInfo")
	g.gen.P("  for _, detail := range s.Details {")
	g.gen.P("    unmarshalled, err := anypb.UnmarshalNew(detail, proto.UnmarshalOptions{})")
	g.gen.P("    if err != nil {")
	g.gen.P("      return nil, err")
	g.gen.P("    }")
	g.gen.P("    switch v := unmarshalled.(type) {")
	g.gen.P("    case *errdetails.ErrorInfo:")
	g.gen.P("      errorInfoDetails = v")
	g.gen.P("    case *errdetails.BadRequest:")
	g.gen.P("      badRequestDetails = v")
	g.gen.P("    case *errdetails.DebugInfo:")
	g.gen.P("      debugInfoDetails = v")
	g.gen.P("    case *errdetails.QuotaFailure:")
	g.gen.P("      quotaFailureDetails = v")
	g.gen.P("    case *errdetails.PreconditionFailure:")
	g.gen.P("      preconditionFailureDetails = v")
	g.gen.P("    case *errdetails.ResourceInfo:")
	g.gen.P("      resourceInfoDetails = v")
	g.gen.P("    default:")
	g.gen.P("      return nil, fmt.Errorf(\"unknown error detail: %s\", v)")
	g.gen.P("    }")
	g.gen.P("  }")
	g.gen.P("  return &details{")
	g.gen.P("    Code: code.Code(s.Code),")
	g.gen.P("    ErrorInfoDetails: errorInfoDetails,")
	g.gen.P("    BadRequestDetails: badRequestDetails,")
	g.gen.P("    DebugInfoDetails: debugInfoDetails,")
	g.gen.P("    QuotaFailureDetails: quotaFailureDetails,")
	g.gen.P("    PreconditionFailureDetails: preconditionFailureDetails,")
	g.gen.P("    ResourceInfoDetails: resourceInfoDetails,")
	g.gen.P("  }, nil")
	g.gen.P("}")

}

func (g *generator) addDecoderMapper(enum *protogen.Enum) {
	g.gen.P(fmt.Sprintf("func %sDecoder(payload []byte) error {", enum.GoIdent.GoName))
	g.gen.P("  details, err := decode(payload)")
	g.gen.P("  if err != nil {")
	g.gen.P("    return err")
	g.gen.P("  }")

	g.gen.P("  switch details.ErrorInfoDetails.Reason {")
	for _, value := range enum.Values {
		options := getOptions(value)

		if options == nil {
			continue
		}

		nameWithoutPrefix := stripEnumValuePrefix(value.GoIdent.GoName, enum.GoIdent.GoName)
		name := uppperSnakeCaseToPascalCase(nameWithoutPrefix)

		g.gen.P(fmt.Sprintf("  case %sReason:", name))

		switch options.Code {
		case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.RetryInfoDetails, details.DebugInfoDetails)", name))
		case int32(code.Code_INVALID_ARGUMENT):
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.BadRequestDetails, details.DebugInfoDetails)", name))
		case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.ResourceInfoDetails, details.DebugInfoDetails)", name))
		case int32(code.Code_RESOURCE_EXHAUSTED):
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.QuotaFailureDetails, details.DebugInfoDetails)", name))
		case int32(code.Code_FAILED_PRECONDITION):
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.PreconditionFailureDetails, details.DebugInfoDetails)", name))
		default:
			g.gen.P(fmt.Sprintf("    return New%s(details.Code, details.ErrorInfoDetails, details.DebugInfoDetails)", name))
		}

	}
	g.gen.P("  default:")
	g.gen.P("    return fmt.Errorf(\"unknown reason: %s\", details.ErrorInfoDetails.Reason)")
	g.gen.P("  }")
	g.gen.P("}")
}

func (g *generator) addEncoders(enum *protogen.Enum) {
	for _, value := range enum.Values {
		options := getOptions(value)

		if options == nil {
			continue
		}

		g.addEncoder(value, enum)
		g.addEmptyLine()
	}
}

func (g *generator) addEncoder(value *protogen.EnumValue, enum *protogen.Enum) {
	name := uppperSnakeCaseToPascalCase(stripEnumValuePrefix(value.GoIdent.GoName, enum.GoIdent.GoName))
	options := getOptions(value)

	g.gen.P(fmt.Sprintf("func %sEncoder(", name))
	g.gen.P("  errorInfoMetadata map[string]string,")

	switch options.Code {
	case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
		g.gen.P("  retryInfoDetails *errdetails.RetryInfo,")
	case int32(code.Code_INVALID_ARGUMENT):
		g.gen.P("  badRequestDetails *errdetails.BadRequest,")
	case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
		g.gen.P("  resourceInfoDetails *errdetails.ResourceInfo,")
	case int32(code.Code_RESOURCE_EXHAUSTED):
		g.gen.P("  quotaFailureDetails *errdetails.QuotaFailure,")
	case int32(code.Code_FAILED_PRECONDITION):
		g.gen.P("  preconditionFailureDetails *errdetails.PreconditionFailure,")
	}

	g.gen.P("debugInfoDetails *errdetails.DebugInfo,")
	g.gen.P(") ([]byte, error) {")
	g.gen.P("  var details []*anypb.Any")
	g.gen.P("  infoDetail, err := anypb.New(&errdetails.ErrorInfo{")
	g.gen.P(fmt.Sprintf("    Reason: %sReason,", name))
	g.gen.P("    Metadata: errorInfoMetadata,")
	g.gen.P(fmt.Sprintf("    Domain: %sDomain,", name))
	g.gen.P("  })")
	g.gen.P("  if err != nil {")
	g.gen.P("    return nil, err")
	g.gen.P("  }")
	g.gen.P("  details = append(details, infoDetail)")
	switch options.Code {
	case int32(code.Code_INVALID_ARGUMENT):
		g.gen.P("  if badRequestDetails != nil {")
		g.gen.P("    detail, err := anypb.New(badRequestDetails)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    details = append(details, detail)")
		g.gen.P("  }")

	case int32(code.Code_RESOURCE_EXHAUSTED):
		g.gen.P("  if quotaFailureDetails != nil {")
		g.gen.P("    detail, err := anypb.New(quotaFailureDetails)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    details = append(details, detail)")
		g.gen.P("  }")

	case int32(code.Code_FAILED_PRECONDITION):
		g.gen.P("  if preconditionFailureDetails != nil {")
		g.gen.P("    detail, err := anypb.New(preconditionFailureDetails)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    details = append(details, detail)")
		g.gen.P("  }")

	case int32(code.Code_DEADLINE_EXCEEDED), int32(code.Code_UNAVAILABLE):
		g.gen.P("  if retryInfoDetails != nil {")
		g.gen.P("    detail, err := anypb.New(retryInfoDetails)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    details = append(details, detail)")
		g.gen.P("  }")

	case int32(code.Code_NOT_FOUND), int32(code.Code_ALREADY_EXISTS):
		g.gen.P("  if resourceInfoDetails != nil {")
		g.gen.P("    detail, err := anypb.New(resourceInfoDetails)")
		g.gen.P("    if err != nil {")
		g.gen.P("      return nil, err")
		g.gen.P("    }")
		g.gen.P("    details = append(details, detail)")
		g.gen.P("  }")

	default:
	}

	g.gen.P("  if debugInfoDetails != nil {")
	g.gen.P("    detail, err := anypb.New(debugInfoDetails)")
	g.gen.P("    if err != nil {")
	g.gen.P("      return nil, err")
	g.gen.P("    }")
	g.gen.P("    details = append(details, detail)")
	g.gen.P("  }")
	g.gen.P(fmt.Sprintf("  return encode(%sCode, details)", name))
	g.gen.P("}")
}

func (g *generator) addHeaderComment() {
	g.gen.P("// Code generated by protoc-gen-rpc-errormapper-go. DO NOT EDIT.")
	g.gen.P(fmt.Sprintf("// source: %s", *g.protoFile.Proto.Name))
}

func (g *generator) addEmptyLine() {
	g.gen.P()
}

func stripEnumValuePrefix(s, prefix string) string {
	return strings.TrimPrefix(s, fmt.Sprintf("%s_", prefix))
}

func uppperSnakeCaseToPascalCase(s string) string {
	words := strings.Split(s, "_")

	var str string
	for _, word := range words {
		str += cases.Title(language.Und).String(word)
	}
	return str
}

func New(
	enums []*protogen.Enum,
	protoFile *protogen.File,
	generatedFile *protogen.GeneratedFile,
) *generator {
	return &generator{
		enums:     enums,
		protoFile: protoFile,
		gen:       generatedFile,
	}
}
