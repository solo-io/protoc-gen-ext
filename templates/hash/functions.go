package hash

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func register(tpl *template.Template, params pgs.Parameters) {
	fns := goSharedFuncs{
		Context: pgsgo.InitContext(params),
		tpl:     tpl,
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
		"oneof":           fns.oneofTypeName,
		"pkg":             fns.PackageName,
		"typ":             fns.Type,
		"externalEnums":   fns.externalEnums,
		"enumPackages":    fns.enumPackages,
		"render":          fns.render,
	})
}

type goSharedFuncs struct {
	pgsgo.Context
	tpl *template.Template
}

func (fns goSharedFuncs) accessor(field pgs.Field) string {
	return fmt.Sprintf("m.Get%s()", fns.Name(field))
}

func (fns goSharedFuncs) pointerAccessor(field pgs.Field) string {
	return fmt.Sprintf("&m.%s", fns.Name(field))
}

func (fns goSharedFuncs) fieldName(field pgs.Field) string {
	return fmt.Sprintf("%s", fns.Name(field))
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

func (fns goSharedFuncs) isBytes(f interface {
	ProtoType() pgs.ProtoType
}) bool {
	return f.ProtoType() == pgs.BytesT
}

func (fns goSharedFuncs) byteStr(x []byte) string {
	elms := make([]string, len(x))
	for i, b := range x {
		elms[i] = fmt.Sprintf(`\x%X`, b)
	}

	return fmt.Sprintf(`"%s"`, strings.Join(elms, ""))
}

func (fns goSharedFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f)).Pointer()
}

func (fns goSharedFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func (fns goSharedFuncs) externalEnums(file pgs.File) []pgs.Enum {
	var out []pgs.Enum

	for _, msg := range file.AllMessages() {
		for _, fld := range msg.Fields() {
			if en := fld.Type().Enum(); fld.Type().IsEnum() && en.Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en) != fns.PackageName(fld) {
				out = append(out, en)
			}
		}
	}

	return out
}

func (fns goSharedFuncs) enumPackages(enums []pgs.Enum) map[pgs.FilePath]pgs.Name {
	out := make(map[pgs.FilePath]pgs.Name, len(enums))

	for _, en := range enums {
		out[fns.ImportPath(en)] = fns.PackageName(en)
	}

	return out
}

func (fns goSharedFuncs) snakeCase(name string) string {
	return strcase.ToSnake(name)
}
