package merge

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/protoc-gen-ext/extproto/gogoproto"
	"github.com/solo-io/protoc-gen-ext/templates"
	"github.com/solo-io/protoc-gen-ext/templates/merge"
)

const (
	mergerName = "merge"
)

type MergeModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Merge() pgs.Module { return &MergeModule{ModuleBase: &pgs.ModuleBase{}} }

func (m *MergeModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *MergeModule) Name() string { return mergerName }

func (m *MergeModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	// Process file-level templates
	tpl := templates.Template(m.Parameters(), merge.Register)

	for _, f := range targets {
		var mergeAll bool
		_, err := f.Extension(extproto.E_MergeAll, &mergeAll)
		if err != nil {
			m.Debugf("error getting merge extension, %s", err.Error())
			continue
		}
		if !mergeAll {
			m.Debugf("Skipping merge functions for %s", f.Name())
			continue
		}
		m.Push(f.Name().String())
		out := templates.FilePathFor(tpl, mergerName)(f, m.ctx, tpl)
		if out != nil {
			m.AddGeneratorTemplateFile(out.String(), tpl, f)
		}

		m.Pop()
	}

	return m.Artifacts()
}

var _ pgs.Module = (*MergeModule)(nil)
