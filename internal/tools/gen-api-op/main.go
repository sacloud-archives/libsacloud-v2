package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const destination = "sacloud/zz_api_ops.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-op: ")
}

func main() {
	outputPath := destination
	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), outputPath),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", outputPath)
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-op'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" "encoding/json" "github.com/sacloud/libsacloud-v2/sacloud/naked" "github.com/sacloud/libsacloud-v2/pkg/mapconv" }}
	{{ . }}
{{- end }}
)

{{ range . }}{{ $typeName := .TypeName}}

/************************************************* 
* {{$typeName}}Op
*************************************************/

// {{ .TypeName }}Op implements {{ .TypeName }}API interface
type {{ .TypeName }}Op struct{
	// Client APICaller
    Client APICaller
	// PathSuffix is used when building URL
	PathSuffix string
	// PathName is used when building URL
	PathName string
}

// New{{ $typeName}}Op creates new {{ $typeName}}Op instance
func New{{ $typeName}}Op(client APICaller) {{ $typeName}}API {
	return &{{ $typeName}}Op {
    	Client: client,
		PathSuffix: "{{.GetPathSuffix}}",
		PathName: "{{.GetPathName}}",
	}
}

{{ range .AllOperations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call
func (o *{{ $typeName }}Op) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	url, err := buildURL("{{.GetPathFormat}}", map[string]interface{}{
		"rootURL": SakuraCloudAPIRoot,
		"pathSuffix": o.PathSuffix,
		"pathName": o.PathName,
		{{- range .AllArguments }}
		"{{.Name}}": {{.Name}},
		{{- end }}
	})
	if err != nil {
		return {{ $returnErrStatement }}
	}

	var body interface{}
	{{- $structName := .RequestEnvelopeStructName }}
	{{- range .PassthroughFieldDeciders}} 
	{{- $argName := .ArgName }}
	{
		if {{.ArgName}} == nil {
			{{.ArgName}} = {{.ZeroInitializer}}
		}
		if body == nil {
			body = &{{$structName}}{}
		}
		v := body.(*{{$structName}})
		if err := mapconv.ConvertTo({{.ArgName}}, v); err != nil {
			return {{ $returnErrStatement }}
		}
		body = v
	}
	{{ end }}
	{{ range .MapDestinationDeciders }} 
	{
		if {{.ArgName}} == nil {
			{{.ArgName}} = {{.ZeroInitializer}}
		}
		if body == nil {
			body = &{{$structName}}{}
		}
		v := body.(*{{$structName}})
		n, err := {{.ArgName}}.convertTo()
		if err != nil {
			return {{ $returnErrStatement }}
		}
		v.{{.DestinationFieldName}} = n 
		body = v
	}
	{{ end }}


	{{ if .HasResponseEnvelope -}}
	data, err := o.Client.Do(ctx, "{{.GetMethod}}", url, body)
	{{ else -}}
	_, err = o.Client.Do(ctx, "{{.GetMethod}}", url, body)
	{{ end -}}
	if err != nil {
		return {{ $returnErrStatement }}
	}

	{{ if .HasResponseEnvelope -}}
	nakedResponse := &{{.ResponseEnvelopeStructName}}{}
	if err := json.Unmarshal(data, nakedResponse); err != nil {
		return {{ $returnErrStatement }}
	}

	{{ if .IsResponseSingular -}}
	{{ range $i,$v := .AllResults -}}
	payload{{$i}} := {{$v.ZeroInitializeSourceCode}}
	if err := payload{{$i}}.convertFrom(nakedResponse.{{.SourceField}}); err != nil {
		return {{ $returnErrStatement }}
	}
	{{ end -}}
	{{- else if .IsResponsePlural -}}
	{{ range $i,$v := .AllResults -}}
	var payload{{$i}} []{{$v.GoTypeSourceCode}}
	for _ , v := range nakedResponse.{{.SourceField}} {
		payload := {{$v.ZeroInitializeSourceCode}}
		if err := payload.convertFrom(v); err != nil {
			return {{ $returnErrStatement }}
		}
		payload{{$i}} = append(payload{{$i}}, payload)
	}
	{{ end -}}
	{{ end -}}
	{{ end -}}

	return {{range $i,$v := .AllResults}}payload{{$i}},{{ end }} nil
}
{{ end -}}
{{ end -}}
`
