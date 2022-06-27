{{/* fields */}}

{{define "fieldExpects"}}
    "{{.Spec.Name}}": {
    Ty: def.Type[int]().Ty.(reflect.Type),

    {{if not (kindIs "invalid" .Spec.Default)}}
        HasDefault: true,
    {{end}}
    },
{{end}}

{{define "fieldVars"}}
    {{if not (kindIs "invalid" .Spec.Default)}}
        _{{.Struct.Pkg.Ns}}_def_field_default__{{.Struct.Spec.Name}}__{{.Spec.Name}} int
    {{end}}
{{end}}

{{define "fieldInit"}}
    field_spec__{{.Struct.Spec.Name}}__{{.Spec.Name}} := struct_spec__{{.Struct.Spec.Name}}.Field("{{.Spec.Name}}")
    _ = field_spec__{{.Struct.Spec.Name}}__{{.Spec.Name}}

    {{if not (kindIs "invalid" .Spec.Default)}}
        _{{.Struct.Pkg.Ns}}_def_field_default__{{.Struct.Spec.Name}}__{{.Spec.Name}} = field_spec__{{.Struct.Spec.Name}}__{{.Spec.Name}}.Default().(int)
    {{end}}
{{end}}

{{define "fieldDecls"}}
    {{.Spec.Name}} int
{{end}}

{{define "fieldInitMethod"}}
    {{if not (kindIs "invalid" .Spec.Default)}}
        f.{{.Spec.Name}} = _{{.Struct.Pkg.Ns}}_def_field_default__{{.Struct.Spec.Name}}__{{.Spec.Name}}
    {{end}}
{{end}}


{{/* structs */}}

{{define "structExpects"}}
    "{{.Spec.Name}}": {
    Fields: map[string]def.X_fieldExpect{
    {{range $_, $v := .Fields}}
        {{include "fieldExpects" $v}}
    {{end}}
    },
    NumInits: {{.Spec.Inits | len}},
    },
{{end}}

{{define "structVars"}}
    {{range $_, $v := .Fields -}}
        {{include "fieldVars" $v}}
    {{- end}}

    _{{.Pkg.Ns}}_def_struct_inits__{{.Spec.Name}} [{{.Spec.Inits | len}}]func(*{{.Spec.Name}})
{{end}}

{{define "structInit"}}
    struct_spec__{{.Spec.Name}} := spec.Struct("{{.Spec.Name}}")
    _ = struct_spec__{{.Spec.Name}}

    {{range $_, $v := .Fields -}}
        {{include "fieldInit" $v}}
    {{- end}}

    {{range $i := intRange 0 (.Spec.Inits | len) -}}
        _{{$.Pkg.Ns}}_def_struct_inits__{{$.Spec.Name}}[{{$i}}] = struct_spec__{{$.Spec.Name}}.Inits()[{{$i}}].(func(*{{$.Spec.Name}}))
    {{- end}}
{{end}}

{{define "structDecls"}}
    type {{.Spec.Name}} struct {
    {{range $_, $v := .Fields -}}
        {{include "fieldDecls" $v}}
    {{- end}}
    }

    func (f *{{.Spec.Name}}) init() {
    _{{.Pkg.Ns}}_def_init()

    {{range $_, $v := .Fields -}}
        {{include "fieldInitMethod" $v}}
    {{- end}}

    {{range $i := intRange 0 (.Spec.Inits | len) -}}
        _{{$.Pkg.Ns}}_def_struct_inits__{{$.Spec.Name}}[{{$i}}](f)
    {{end}}
    }
{{end}}


{{/* package */}}

{{define "packageExpects"}}
    var _ = func() any {
    def.X_addPackageDef(def.X_packageExpect{
    Structs: map[string]def.X_structExpect{
    {{range $_, $v := .Structs}}
        {{include "structExpects" $v}}
    {{end}}
    },
    })
    return nil
    }()
{{end}}

{{define "packageVars"}}
    var (
    _{{.Ns}}_def_init_once sync.Once

    {{range $_, $v := .Structs -}}
        {{include "structVars" $v}}
    {{- end}}
    )
{{end}}

{{define "packageInit"}}
    func _{{.Ns}}_def_init() {
    _{{.Ns}}_def_init_once.Do(func() {
    spec := def.X_getPackageSpec()

    {{range $_, $v := .Structs -}}
        {{include "structInit" $v}}
    {{- end}}
    })
    }
{{end}}

{{define "packageDecls"}}
    {{range $_, $v := .Structs -}}
        {{include "structDecls" $v}}
    {{- end}}
{{end}}


{{/* root */}}

{{define "package"}}
    {{include "packageExpects" .}}
    {{include "packageVars" .}}
    {{include "packageInit" .}}
    {{include "packageDecls" .}}
{{end}}