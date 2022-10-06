package hash

const messageTpl = `
		if h, ok := interface{}({{ .FieldAccessor }}).(safe_hasher.SafeHasher); ok {
			if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
				return 0, err
			}
			if _, err = h.Hash({{ .Hasher }}); err != nil {
				return 0, err 
			}
		} else {
			if fieldValue, err := hashstructure.Hash({{ .FieldAccessor }}, nil); err != nil {
				return 0, err
			} else {
				if _, err = {{ .Hasher }}.Write([]byte("{{ .FieldName }}")); err != nil {
					return 0, err
				}
				var buf [8]byte
				binary.LittleEndian.PutUint64(buf[:], fieldValue)
				if _, err := hasher.Write(buf[:]); err != nil {
					return 0, err
				}
			}
		}
`

const primitiveTmpl = `
		err = binary.Write({{ .Hasher }}, binary.LittleEndian, {{ .FieldAccessor }} )
		if err != nil {return 0, err}
`

const primitiveBoolTmpl = `
		{
			if {{ .FieldAccessor }} {
				_, err = hasher.Write([]byte{1})
			} else {
				_, err = hasher.Write([]byte{0})
			}
			if err != nil {return 0, err}
		}
`
const primitiveFloatTmpl = `
		{
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], math.Float64bits(float64({{ .FieldAccessor }})))
			_, err = hasher.Write(buf[:])
			if err != nil {return 0, err}
		}
`

const primitiveIntTmpl = `
		{
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], uint64({{ .FieldAccessor }}))
			_, err = hasher.Write(buf[:])
			if err != nil {return 0, err}
		}
`

const bytesTpl = `
		if _, err = {{ .Hasher }}.Write({{ .FieldAccessor }}); err != nil {
			return 0, err
		}
`

const stringTpl = `
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
			var buf [8]byte
			binary.LittleEndian.PutUint64(buf[:], result)
			_, err := hasher.Write(buf[:])
			if err != nil {return 0, err}
			
	}
`

const repeatedTpl = `
		for _, v := range {{ .FieldAccessor }} {
			{{ .InnerTemplates.Value }}
		}
`
