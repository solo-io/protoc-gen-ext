package golang

const messageTpl = `
		if h, ok := interface{}({{ .Name }}).(interface{ Hash({{ .Hasher }} hash.Hash64) (uint64, error) }); ok {
			if _, err = h.Hash({{ .Hasher }}); err != nil {
				return  0, err 
			}
		} else {
			if val, err := hashstructure.Hash({{ .Name }}, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write({{ .Hasher }}, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}
`

const primitiveTmpl = `
		err = binary.Write({{ .Hasher }}, binary.LittleEndian,  {{ .Name }} )
		if err != nil {return 0, err}
`

const bytesTpl = `
		if _, err = {{ .Hasher }}.Write({{ .Name }}); err != nil { 
			return 0, err 
		}
`

const stringTpl = `
		if _, err = {{ .Hasher }}.Write([]byte({{ .Name }})); err != nil { 
			return 0, err 
		}
`

const mapTpl = `
		innerHash := fnv.Hash64()
		for k , v := range {{ .Name }} {
			innerHash.Reset()
			
			{{  }}

			
			err = binary.Write(hasher, binary.LittleEndian, innerHash.Sum64())
			if err != nil {return 0, err}
			err = binary.Write(hasher, binary.LittleEndian, innerHash.Sum64())
		}
`

const repeatedTpl = `
		for _, v := range {{ .Name }} {
			{{ .InnerTemplates.Value }}
		}
`