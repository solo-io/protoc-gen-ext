package clone

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func register(tpl *template.Template, params pgs.Parameters, common pgs.DebuggerCommon) {
	fns := goSharedFuncs{
		Context:        pgsgo.InitContext(params),
		DebuggerCommon: common,
		tpl:            tpl,
	}

	tpl.Funcs(map[string]interface{}{
		"fullPackageName": fns.fullPackageName,
		"snakeCase":       fns.snakeCase,
		"cmt":             pgs.C80,
		"isBytes":         fns.isBytes,
		"lit":             fns.lit,
		"lookup":          fns.lookup,
		"msgTyp":          fns.msgTyp,
		"name":            fns.Name,
		"oneof":           fns.oneofPointerTypeName,
		"oneofNonPointer": fns.oneofTypeName,
		"pkg":             fns.PackageName,
		"typ":             fns.Type,
		"externalFields":  fns.externalFields,
		"render":          fns.render,
		"oneofRender":     fns.oneofRender,
	})
}

type goSharedFuncs struct {
	pgsgo.Context
	pgs.DebuggerCommon
	tpl *template.Template
}

func (fns goSharedFuncs) accessor(field pgs.Field) string {
	return fmt.Sprintf("m.Get%s()", fns.Name(field))
}

func (fns goSharedFuncs) targetAccessor(field pgs.Field) string {
	return fmt.Sprintf("target.%s", fns.Name(field))
}

func (fns goSharedFuncs) pointerAccessor(field pgs.Field) string {
	return fmt.Sprintf("&m.%s", fns.Name(field))
}

func (fns goSharedFuncs) typeName(field pgs.Field) string {
	return fmt.Sprintf("%s", fns.Type(field))
}

func (fns goSharedFuncs) entityTypeName(f pgs.Field, field pgs.FieldTypeElem) string {
	e := field.Embed()
	t := pgsgo.TypeName(fns.Name(e))
	if fns.ImportPath(e) == fns.ImportPath(f) {
		return fmt.Sprintf("*%s", t.String())
	}

	return pgsgo.TypeName(fmt.Sprintf("*%s.%s", fns.PackageName(field.Embed()), t)).String()
}

func (fns goSharedFuncs) fullPackageName(msg pgs.Message) string {
	return fmt.Sprintf("%s.%s.%s", msg.Package().ProtoName(), fns.ImportPath(msg), fns.Name(msg))
}

func (fns goSharedFuncs) lookup(f pgs.Field, name string) string {
	return fmt.Sprintf(
		"_%s_%s_%s",
		fns.Name(f.Message()),
		fns.Name(f),
		name,
	)
}

func (fns goSharedFuncs) lit(x interface{}) string {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", x)
	case reflect.Uint8:
		return fmt.Sprintf("0x%X", x)
	case reflect.Slice:
		els := make([]string, val.Len())
		for i, l := 0, val.Len(); i < l; i++ {
			els[i] = fns.lit(val.Index(i).Interface())
		}
		return fmt.Sprintf("%T{%s}", val.Interface(), strings.Join(els, ", "))
	default:
		return fmt.Sprint(x)
	}
}

func (fns goSharedFuncs) isBytes(
	f interface {
		ProtoType() pgs.ProtoType
	},
) bool {
	return f.ProtoType() == pgs.BytesT
}

func (fns goSharedFuncs) byteStr(x []byte) string {
	elms := make([]string, len(x))
	for i, b := range x {
		elms[i] = fmt.Sprintf(`\x%X`, b)
	}

	return fmt.Sprintf(`"%s"`, strings.Join(elms, ""))
}

func (fns goSharedFuncs) oneofPointerTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f)).Pointer()
}

func (fns goSharedFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f))
}

func (fns goSharedFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func (fns goSharedFuncs) externalFields(file pgs.File) map[pgs.FilePath]pgs.Name {
	var (
		out []pgs.Entity
	)
	//
	// for _, v := range file.TransitiveImports() {
	// 	if fns.PackageName(v) != fns.PackageName(file) {
	// 		out = append(out, v)
	// 	}
	// }

	for _, msg := range file.AllMessages() {
		for _, fld := range msg.Fields() {
			// if en := fld.Type().Embed(); fld.Type().IsEmbed() && en.Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en) != fns.PackageName(fld) {
			// 	out = append(out, en)
			// }
			// if en := fld.Type().Enum(); fld.Type().IsEnum() && en.Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en) != fns.PackageName(fld) {
			// 	out = append(out, en)
			// }
			if en := fld.Type().Embed(); fld.Type().IsEmbed() {
				out = append(out, en)
			}
			if en := fld.Type().Enum(); fld.Type().IsEnum() {
				out = append(out, en)
			}
			// // Handle repeated
			// if en := fld.Type().Element(); (fld.Type().IsRepeated() || fld.Type().IsMap()) && en.IsEmbed() && en.Embed().Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en.Embed()) != fns.PackageName(fld) {
			// 	out = append(out, en.Embed())
			// }
			//
			// // Handle map
			// if en := fld.Type().Key(); fld.Type().IsMap() && en.IsEmbed() && en.Embed().Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en.Embed()) != fns.PackageName(fld) {
			// 	out = append(out, en.Embed())
			// }
			// Handle repeated
			if en := fld.Type().Element(); (fld.Type().IsRepeated() || fld.Type().IsMap()) && en.IsEmbed() {
				out = append(out, en.Embed())
			}

			// Handle map
			if en := fld.Type().Key(); fld.Type().IsMap() && en.IsEmbed() {
				out = append(out, en.Embed())
			}
		}
	}

	return fns.externalPackages(out, file)
}

func (fns goSharedFuncs) externalPackages(
	entities []pgs.Entity,
	file pgs.File,
) map[pgs.FilePath]pgs.Name {
	out := make(map[pgs.FilePath]pgs.Name, len(entities))

	for _, en := range entities {
		if fns.PackageName(en) == fns.PackageName(file) && en.Package().ProtoName() == file.Package().ProtoName() {
			continue
		}
		fns.Debugf("adding entity with file path (%s), and file name (%s)", fns.ImportPath(en), fns.PackageName(en))
		out[fns.ImportPath(en)] = fns.PackageName(en)
	}

	return out
}

func (fns goSharedFuncs) snakeCase(name string) string {
	return strcase.ToSnake(name)
}
