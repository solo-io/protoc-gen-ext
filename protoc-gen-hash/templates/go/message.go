package golang

const messageTpl = `
		
		if h, ok := interface{}({{ accessor . }}).(interface{ Hash(hasher hash.Hash64) (uint64, error) }); ok {
			if val, err = h.Hash(hasher); err != nil {
				return  0, err 
			} else {
				err = binary.Write(hasher, binary.LittleEndian, val)
				if err != nil {
					return 0, err
				}
			}
		} else {
			if val, err = hashstructure.Hash(h, nil); err != nil {
				return 0, err
			} else {
				err = binary.Write(hasher, binary.LittleEndian, val)
				if err != nil {
					return 0, err
				}
			}
		}
`
