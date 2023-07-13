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
	if field.Type().IsEmbed() {
		return fns.importableTypeName(field, field.Type().Embed())
	} else if field.Type().IsMap() || field.Type().IsRepeated() {
		// We only care about embedded elements

		// Replace the pacakge name
		if field.Type().Element().IsEmbed() {
			typeName := fmt.Sprintf("%s", fns.Type(field))
			e := field.Type().Element().Embed()
			properImportPath := fns.packageName(e)
			typeName = strings.ReplaceAll(typeName, fns.PackageName(e).String(), properImportPath)
			return typeName
		}
	}

	fns.Debugf("getting type for embedded standard field (%s)", field.Name())
	return fmt.Sprintf("%s", fns.Type(field))
}

func (fns goSharedFuncs) entityTypeName(f pgs.Field, field pgs.FieldTypeElem) string {
	e := field.Embed()
	return fns.importableTypeName(f, e)
}

func (fns goSharedFuncs) importableTypeName(f pgs.Field, e pgs.Entity) string {
	t := pgsgo.TypeName(fns.Name(e))

	if fns.ImportPath(e) == fns.ImportPath(f) {
		return fmt.Sprintf("*%s", t.String())
	}

	return fmt.Sprintf("*%s.%s", fns.packageName(e), t)
}

func (fns goSharedFuncs) fieldPackageName(field pgs.Field) string {
	if field.Type().IsMap() || field.Type().IsRepeated() {
		if field.Type().Element().IsEmbed() {
			e := field.Type().Element().Embed()
			return fns.packageName(e)
		}
	}

	return fns.packageName(field.Type().Embed())

}

func (fns goSharedFuncs) packageName(e pgs.Entity) string {
	importName := fns.ImportPath(e).String()
	importName = strings.ReplaceAll(importName, "/", "_")
	importName = strings.ReplaceAll(importName, ".", "_")
	importName = strings.ReplaceAll(importName, "-", "_")
	fns.Debugf("packageName (%s)", importName)
	return importName
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

func (fns goSharedFuncs) externalFields(file pgs.File) map[pgs.FilePath]string {
	var (
		out []pgs.Entity
	)

	for _, msg := range file.AllMessages() {
		for _, fld := range msg.Fields() {
			if en := fld.Type().Embed(); fld.Type().IsEmbed() {
				out = append(out, en)
			}
			if en := fld.Type().Enum(); fld.Type().IsEnum() {
				out = append(out, en)
			}
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
) map[pgs.FilePath]string {
	out := make(map[pgs.FilePath]string, len(entities))

	for _, en := range entities {
		if fns.PackageName(en) == fns.PackageName(file) && en.Package().ProtoName() == file.Package().ProtoName() {
			continue
		}
		fns.Debugf("adding entity with file path (%s), and file name (%s)", fns.ImportPath(en), fns.PackageName(en))
		out[fns.ImportPath(en)] = fns.packageName(en)
	}

	return out
}

func (fns goSharedFuncs) snakeCase(name string) string {
	return strcase.ToSnake(name)
}

func (fns goSharedFuncs) contains(strings []string, s string) bool {
	for _, str := range strings {
		if str == s {
			return true
		}
	}
	return false
}
