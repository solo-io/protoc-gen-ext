package clone

const msgTpl = `
// Clone function
func (m {{ (msgTyp .).Pointer }}) Clone() proto.Message {
		if m == nil {
			return nil
		}
		target := (msgTyp .).Pointer

		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ $oneOfInterface := name .}}
				{{ range .Fields }}
					case {{ oneof . }}:
						if _, ok := target.{{ $oneOfInterface }}.({{ oneof . }}); !ok {
							return target
						}

						{{ render . }}
				{{ end }}
					default:
						// m is nil but target is not nil
						if m.{{ name . }} != target.{{ name . }} {
							return target
						}
			}

		{{ end }}


		return target
}

`
