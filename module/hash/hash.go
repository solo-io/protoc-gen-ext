package hash

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/protoc-gen-ext/extproto/gogoproto"
	"github.com/solo-io/protoc-gen-ext/templates"
)

const (
	hasherName = "hash"
)

type HashModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Hash() pgs.Module { return &HashModule{ModuleBase: &pgs.ModuleBase{}} }

func (m *HashModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *HashModule) Name() string { return hasherName }

func (m *HashModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	// Process file-level templates
	tpl := templates.Template(m.Parameters())

	for _, f := range targets {
		var hashAll bool
		_, err := f.Extension(extproto.E_HashAll, &hashAll)
		if err != nil {
			m.Debugf("error getting hash extension, %s", err.Error())
			continue
		}
		if !hashAll {
			m.Debugf("Skipping hash functions for %s", f.Name())
			continue
		}
		m.Push(f.Name().String())
		out := templates.FilePathFor(tpl, hasherName)(f, m.ctx, tpl)
		if out != nil {
			m.AddGeneratorTemplateFile(out.String(), tpl, f)
		}

		m.Pop()
	}

	return m.Artifacts()
}

var _ pgs.Module = (*HashModule)(nil)
