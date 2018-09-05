package template

var (
	Gen = `// Package {{ .Name }} provides access to the {{ .Title }}.{{ if .ReplacementPackage }}
//
// This package is DEPRECATED. Use package {{ .ReplacementPackage }} instead.{{ end }}{{ if .DocumentationLink }}
//
// See {{ .DocumentationLink }}{{ end }}
//
// Usage example:
//
//   import "{{ package .Name .Version }}"
//   ...
//   {{ .Name }}Service, err := {{ .Name }}.New(oauthHttpClient)
package {{ .Name }} // import "{{ package .Name .Version }}"

import ({{range .ImportPackage}}
	{{ if .Alias }}{{ .Alias }} {{ end }}"{{ .Package }}"{{ end }}
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "{{ .ID }}"
const apiName = "{{ .Name }}"
const apiVersion = "{{ .Version }}"
const basePath = "{{ base_url }}"{{ if .Auth.OAuth2Scopes }}

// OAuth2 scopes used by this API.
const ({{ range .Auth.OAuth2Scopes }}{{ if .Description }}
	{{ asComment "" .Description }}{{ end }}
	{{ scopeIdentifierFromURL .URL }} = "{{ .URL }}"{{ end }}
){{ end }}

func New(client *http.Client) (*{{ service_type }}, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &{{ service_type }}{client: client, BasePath: basePath}{{ range .Resources }}
	s.{{ resourceGoField . nil }} = New{{ resourceGoType . }}(s){{ end }}
	return s, nil
}

type {{ service_type }} struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment{{ range .Resources }}
	{{ resourceGoField . nil }} *{{ resourceGoType . }}{{ end }}
}

func (s *{{ service_type }}) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}


{{define "tmpl_resource"}}
	func New{{ resourceGoType . }}(s *{{ service_type }}) *{{ resourceGoType . }} {
		rs := &{{ resourceGoType . }}{s: s}
		return rs
	}

	type {{ resourceGoType . }} struct {
		s *{{ service_type }}{{ range $k, $v := .Resources }}
	{{ resourceGoField $v . }} *{{ resourceGoType $v }}{{ end }}
	}

	{{ range .Resources }}{{ template "tmpl_resource" .Resources }}{{ end }}
{{ end }}


{{ range .Resources }}{{ template "tmpl_resource" . }}{{ end }}



{{define "tmpl_schema"}}
{{ if .Variant }}// TODO:: Variant {{ end }}
// {{ .Description }}
{{ end }}



{{ range .Schemas }}

{{ if eq .Kind 1 }}
{{ template "tmpl_schema" . }}
{{ end }}

{{ end }}

`
)
