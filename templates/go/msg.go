package golang

const msgTpl = `
// Hash function
func (m {{ (msgTyp .).Pointer }}) Hash(hasher hash.Hash64) (uint64, error) {
		if m == nil { return 0, nil }
		if hasher == nil { hasher = fnv.New64() }
		{{- if not (eq (len .Fields) 0) }}
		var err error
		{{- end }}
		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ range .Fields }}
					case {{ oneof . }}:
						{{ render . }}
				{{ end }}
			}
		{{ end }}


		return hasher.Sum64(), nil
}

`
