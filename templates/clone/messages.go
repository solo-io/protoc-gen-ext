package clone

const messageTpl = `
		if h, ok := interface{}({{ .Name }}).(clone.Cloner); ok {
			if !h.Equal({{.TargetName}}) {
				return false
			}
		} else {
			if !proto.Clone({{ .Name }}, {{.TargetName}}) {
				return false
			}
		}
`

const primitiveTmpl = `
		if {{ .Name }} != {{ .TargetName }} {
			return false
		}
`

const bytesTpl = `
		if bytes.Compare({{.Name}}, {{ .TargetName}}) != 0 {
			return false
		}
`

const stringTpl = `
		if strings.Compare({{.Name}}, {{ .TargetName}}) != 0 {
			return false
		}
`

const mapTpl = `
		if len({{ .Name }}) != len({{ .TargetName }}) {
			return false
		}
		for k , v := range {{ .Name }} {
			{{ .InnerTemplates.Value }}
		}
			
`

const repeatedTpl = `
		if len({{ .Name }}) != len({{ .TargetName }}) {
			return false
		}
		for idx, v := range {{ .Name }} {
			{{ .InnerTemplates.Value }}
		}
`
