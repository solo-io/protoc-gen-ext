package hash

const msgTpl = `
// Hash function
func (m {{ (msgTyp .).Pointer }}) Hash(hasher hash.Hash64) (uint64, error) {
		if m == nil { return 0, nil }
		if hasher == nil { hasher = fnv.New64() }
		var err error
		if _, err = hasher.Write([]byte("{{ fullPackageName . }}")); err != nil { 
			return 0, err 
		}

		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .SyntheticOneOfFields }}
			{{ render . }}
		{{ end }}
		{{ range .RealOneOfs }}
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
