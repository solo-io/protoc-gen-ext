package merge

const msgTpl = `
// Merge non-nil fields from overrides into m
func (m {{ (msgTyp .).Pointer }}) Merge(overrides {{ (msgTyp .).Pointer }}) {
	if m == nil || overrides == nil { return }

	{{ range .NonOneOfFields }}
		{{ renderField . }}
	{{ end }}

	{{ range .OneOfs }}
		{{ renderOneOf . }}
	{{ end }}

}

`
