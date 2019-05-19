package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const destination = "sacloud/stub/zz_api_stubs.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-stub: ")
}

func main() {
	schema.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-stub'; DO NOT EDIT

package stub

import (
{{- range .ImportStatements "context" "log" }}
	{{ . }}
{{- end }}
)

{{ range . }} {{ $typeName := .TypeName }}

/************************************************* 
* {{ $typeName }}Stub
*************************************************/

{{ range .AllOperations }}
// {{ $typeName }}{{.MethodName}}Result is expected values of the {{ .MethodName }} operation
type {{ $typeName }}{{.MethodName}}Result struct {
	{{ range .StubFieldDefines -}}
	{{ . }}
	{{ end -}}	
	Err error
}
{{ end -}}

// {{ $typeName }}Stub is for trace {{ $typeName }}Op operations
type {{ $typeName }}Stub struct {
{{ range .AllOperations -}}
	{{.MethodName}}Result *{{ $typeName }}{{.MethodName}}Result 
{{ end -}}
}

// New{{ $typeName}}Stub creates new {{ $typeName}}Stub instance
func New{{ $typeName}}Stub(caller sacloud.APICaller) sacloud.{{$typeName}}API {
	return &{{ $typeName}}Stub{}
}

{{ range .AllOperations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call with trace log
func (s *{{ $typeName }}Stub) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	if s.{{$operationName}}Result == nil {
		log.Fatal("{{$typeName}}Stub.{{$operationName}}Result is not set")
	}
	{{.StubReturnStatement "s"}}
}
{{- end -}}

{{ end }}
`
