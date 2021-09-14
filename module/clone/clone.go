package clone

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/solo-io/protoc-gen-ext/extproto"
	"github.com/solo-io/protoc-gen-ext/templates"
	"github.com/solo-io/protoc-gen-ext/templates/clone"
)

const (
	cloneName = "clone"
)

type CloneModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Clone() pgs.Module { return &CloneModule{ModuleBase: &pgs.ModuleBase{}} }

func (m *CloneModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *CloneModule) Name() string { return cloneName }

func (m *CloneModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	// Process file-level templates
	tpl := templates.Template(m.Parameters(), clone.Register)
	for _, f := range targets {
		var cloneAll bool
		_, err := f.Extension(extproto.E_CloneAll, &cloneAll)
		if err != nil {
			m.Debugf("error getting clone extension, %s", err.Error())
			continue
		}
		if !cloneAll {
			m.Debugf("Skipping clone functions for %s", f.Name())
			continue
		}
		m.Push(f.Name().String())
		out := templates.FilePathFor(tpl, cloneName)(f, m.ctx, tpl)
		if out != nil {
			m.AddGeneratorTemplateFile(out.String(), tpl, f)
		}

		m.Pop()
	}

	return m.Artifacts()
}

var _ pgs.Module = (*CloneModule)(nil)
