{{/* fields */}}

{{define "fieldVars"}}
    _{{.Struct.Pkg.Ns}}_def_field_default__{{.Struct.Spec.Name}}__{{.Spec.Name}} int
{{end}}

{{define "fieldInit"}}
    field_spec__{{.Struct.Spec.Name}}__{{.Spec.Name}} := struct_spec__{{.Struct.Spec.Name}}.Field("{{.Spec.Name}}")
{{end}}


{{/* structs */}}

{{define "structVars"}}
    {{range $_, $v := .Fields -}}
    {{include "fieldVars" $v}}
    {{- end}}

    _{{.Pkg.Ns}}_def_struct_inits__{{.Spec.Name}} [{{.Spec.Inits | len}}]func(*{{.Spec.Name}})
{{end}}

{{define "structInit"}}
    struct_spec__{{.Spec.Name}} := spec.Struct("{{.Spec.Name}}")

    {{range $_, $v := .Fields -}}
    {{include "fieldInit" $v}}
    {{- end}}

    {{range $i := intRange 0 (.Spec.Inits | len)}}
    _{{$.Pkg.Ns}}_def_struct_inits__{{$.Spec.Name}}[{{$i}}] = struct_spec__{{$.Spec.Name}}.Inits()[{{$i}}].(func(*{{$.Spec.Name}}))
    {{end -}}
{{end}}


{{/* package */}}

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
            {{include "structInit" $v | indent 4 " "}}
        {{- end}}
    })
}
{{end}}

{{define "package"}}
{{include "packageVars" .}}
{{include "packageInit" .}}
{{end}}