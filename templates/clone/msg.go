package clone

const msgTpl = `
// Equal function
func (m {{ (msgTyp .).Pointer }}) Equal(that interface{}) bool {
		if that == nil {
			return m == nil
		}
	
		target, ok := that.({{ (msgTyp .).Pointer }})
		if !ok {
			that2, ok := that.({{msgTyp .}})
			if ok {
				target = &that2
			} else {
				return false
			}
		}
		if target == nil {
			return m == nil
		} else if m == nil {
			return false
		}

		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ $oneOfInterface := name .}}
				{{ range .Fields }}
					case {{ oneof . }}:
						if _, ok := target.{{ $oneOfInterface }}.({{ oneof . }}); !ok {
							return false
						}

						{{ render . }}
				{{ end }}
					default:
						// m is nil but target is not nil
						if m.{{ name . }} != target.{{ name . }} {
							return false
						}
			}

		{{ end }}


		return true
}

`
