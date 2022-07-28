package merge

const msgTpl = `
// Merge non-nil fields from overrides into m
func (m {{ (msgTyp .).Pointer }}) Merge(overrides {{ (msgTyp .).Pointer }}) {
	if m == nil || overrides == nil { return }

	{{ range .NonOneOfFields }}
		{{ renderField . }}
	{{ end }}

	{{ range .SyntheticOneOfFields }}
	if overrides.{{ name . }} != nil {
		v := overrides.Get{{ name . }}()
		m.{{ name . }} = &v
	}
	{{ end }}

	{{ range .RealOneOfs }}
		{{ renderOneOf . }}
	{{ end }}

}

`
