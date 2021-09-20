package clone

const msgTpl = `
// Clone function
func (m {{ (msgTyp .).Pointer }}) Clone() proto.Message {
		var target *{{ name .}}
		if m == nil {
			return target
		}
		target = &{{msgTyp .}}{}

		{{ range .NonOneOfFields }}
			{{ render . }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ $oneOfInterface := name .}}
				{{ range .Fields }}
					case {{ oneof . }}:

						{{ oneofRender . $oneOfInterface}}
				{{ end }}
			}

		{{ end }}


		return target
}

`
