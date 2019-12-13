package golang

const primitiveTmpl = `
		err = binary.Write(hasher, binary.LittleEndian,  {{ accessor . }} )
		if err != nil {return 0, err}
`

const bytesTpl = `

`

const stringTpl = `
`
