package golang

const msgTpl = `
// Hash function
func (m {{ (msgTyp .).Pointer }}) Hash(hasher hash.Hash64) (uint64, error) {
		if m == nil { return 0, nil }
		if hasher == nil { hasher = fnv.New64() }
		hasher.Reset()
		var err error
		var val uint64

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
