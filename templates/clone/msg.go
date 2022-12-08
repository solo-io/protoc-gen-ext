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

		{{ range .SyntheticOneOfFields }}
		if m.{{ name . }} != nil {
			v := m.Get{{ name . }}()
			target.{{ name . }} = &v
		}
		{{ end }}

		{{ range .RealOneOfs }}
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
