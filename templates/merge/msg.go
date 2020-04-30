package merge

const msgTpl = `
// Merge non-nil fields from m2 into m
func (m {{ (msgTyp .).Pointer }}) Merge(overrides {{ (msgTyp .).Pointer }}) {
	if m == nil || m2 == nil { return }

	{{ range .NonOneOfFields }}
		{{ renderField . }}
	{{ end }}

	{{ range .OneOfs }}
		{{ renderOneOf . }}
	{{ end }}

}

`
