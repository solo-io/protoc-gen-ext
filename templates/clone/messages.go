package clone

const messageTpl = `
		if h, ok := interface{}({{ .Name }}).(clone.Cloner); ok {
			{{ .TargetName }} = h.Clone().({{.TypeName}})
		} else {
			{{ .TargetName }} = proto.Clone({{.Name}}).({{.TypeName}})
		}
`

const oneofMessageTpl = `
		if h, ok := interface{}({{ .Name }}).(clone.Cloner); ok {
			{{ .TargetName }} = h.Clone().({{.TypeName}})
		} else {
			{{ .TargetName }} = proto.Clone({{.Name}}).({{.TypeName}})
		}
`

const primitiveTmpl = `
		{{ .TargetName }} = {{ .Name }}
`
const stringTpl = primitiveTmpl

const bytesTpl = `
		{{ .TargetName }} = make([]byte, len({{ .Name }}))
		copy({{ .TargetName }}, {{ .Name }})
`

const mapTpl = `
		for k , v := range {{ .Name }} {
			{{ .InnerTemplates.Value }}
		}
`

const repeatedTpl = `
		for idx, v := range {{ .Name }} {
			{{ .InnerTemplates.Value }}
		}
`
