package equal

import (
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/extproto"
	"github.com/solo-io/protoc-gen-ext/templates"
)

const (
	equalName = "equal"
)

type EqualModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Equal() pgs.Module { return &EqualModule{ModuleBase: &pgs.ModuleBase{}} }

func (m *EqualModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *EqualModule) Name() string { return equalName }

func (m *EqualModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	// Process file-level templates
	tpl := templates.Template(m.Parameters())
	for _, f := range targets {
		extension, err := proto.GetExtension(f.Descriptor(), extproto.E_EqualAll)
		if err != nil {
			continue
		}
		hashAll, _ := extension.(bool)
		if !hashAll {
			continue
		}
		m.Push(f.Name().String())
		out := templates.FilePathFor(tpl, equalName)(f, m.ctx, tpl)
		if out != nil {
			m.AddGeneratorTemplateFile(out.String(), tpl, f)
		}

		m.Pop()
	}

	return m.Artifacts()
}

var _ pgs.Module = (*EqualModule)(nil)
