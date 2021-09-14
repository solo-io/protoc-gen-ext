package clone

const msgTpl = `
// Clone function
func (m {{ (msgTyp .).Pointer }}) Clone() proto.Message {
		if m == nil {
			return nil
		}
		target := &{{msgTyp .}}{}

		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ $oneOfInterface := name .}}
				{{ range .Fields }}
					case {{ oneof . }}:

						{{ render . }}
				{{ end }}
			}

		{{ end }}


		return target
}

`
