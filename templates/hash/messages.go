package hash

const messageTpl = `
		if h, ok := interface{}({{ .FieldAccessor }}).(safe_hasher.SafeHasher); ok {
			if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
				return 0, err
			}
			if _, err = h.Hash({{ .Hasher }}); err != nil {
				return  0, err 
			}
		} else {
			if fieldValue, err := hashstructure.Hash({{ .FieldAccessor }}, nil); err != nil {
				return 0, err
			} else {
				if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
					return 0, err
				}
				if err := binary.Write({{ .Hasher }}, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}
`

const primitiveTmpl = `
		{{ if .UniqueHash }}
		if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
			return 0, err
		}
		{{ end }}
		err = binary.Write({{ .Hasher }}, binary.LittleEndian,  {{ .FieldAccessor }} )
		if err != nil {return 0, err}
`

const bytesTpl = `
		{{ if .UniqueHash }}
		if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
			return 0, err
		}
		{{ end }}
		if _, err = {{ .Hasher }}.Write({{ .FieldAccessor }}); err != nil {
			return 0, err
		}
`

const stringTpl = `
		{{ if .UniqueHash }}
		if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
			return 0, err
		}
		{{ end }}
		if _, err = {{ .Hasher }}.Write([]byte({{ .FieldAccessor }})); err != nil {
			return 0, err
		}
`

const mapTpl = `
	{
			var result uint64
			innerHash := fnv.New64()
			for k , v := range {{ .FieldAccessor }} {
				innerHash.Reset()
	
				{{ .InnerTemplates.Value }}
	
	
				{{ .InnerTemplates.Key }}

				result = result ^ innerHash.Sum64()
			}
			err = binary.Write(hasher, binary.LittleEndian, result)
			if err != nil {return 0, err}
			
	}
`

const repeatedTpl = `
		{{ if .UniqueHash }}
		if _, err = hasher.Write([]byte("{{ .FieldName }}")); err != nil {
			return 0, err
		}
		for i, v := range {{ .FieldAccessor }} {
			if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
				return 0, err
			}
		{{ else }}
		for _, v := range {{ .FieldAccessor }} {
		{{ end }}
			{{ .InnerTemplates.Value }}
		}
`
