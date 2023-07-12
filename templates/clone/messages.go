package clone

const messageTpl = `
		if h, ok := interface{}({{ .Name }}).(clone.Cloner); ok {
			{{ .TargetName }} = h.Clone().({{.TypeName}})
		} else {
			{{ .TargetName }} = proto.Clone({{.Name}}).({{.TypeName}})
		}
`

const messageGogoTpl = `
		{{ .TargetName }} = {{ .Name }}.DeepCopy()
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
		if {{ .Name }} != nil {
			{{ .TargetName }} = make([]byte, len({{ .Name }}))
			copy({{ .TargetName }}, {{ .Name }})
		}
`

const oneofBytesTpl = `
		if {{ .Name }} != nil {
			newArr := make([]byte, len({{ .Name }}))
			copy(newArr, {{ .Name }})
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: newArr,
			}
		} else {
			target.{{ .OneOfInterface }} = &{{ oneofNonPointer .Field }}{
				{{ name .Field }}: nil,
			}
		}
`

const mapTpl = `
		if {{ .Name }} != nil {
			{{ .TargetName }} = make({{.TypeName}}, len({{ .Name }}))
			for k , v := range {{ .Name }} {
				{{ .InnerTemplates.Value }}
			}
		}
`

const repeatedTpl = `
		if {{ .Name }} != nil {
			{{ .TargetName }} = make({{.TypeName}}, len({{ .Name }}))
			for idx, v := range {{ .Name }} {
				{{ .InnerTemplates.Value }}
			}	
		}
`
