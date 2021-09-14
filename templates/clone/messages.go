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
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: h.Clone().({{.TypeName}}),
			}
		} else {
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: proto.Clone({{ .Name }}).({{.TypeName}}),
			}
		}
`

const primitiveTmpl = `
		{{ .TargetName }} = {{ .Name }}
`

const stringTpl = primitiveTmpl

const oneofPrimitiveTmpl = `
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: {{ .Name }},
			}
`
const oneofStringTpl = oneofPrimitiveTmpl

const bytesTpl = `
		{{ .TargetName }} = make([]byte, len({{ .Name }}))
		copy({{ .TargetName }}, {{ .Name }})
`

const oneofBytesTpl = `
		{
			newArr := make([]byte, len({{ .Name }}))
			copy(newArr, {{ .Name }})
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: newArr,
			}
		}
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
