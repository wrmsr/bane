{{define "structVars"}}
    _{{$.Pkg.Ns}}_def_struct_inits__{{.Spec.Name}} [1]func(*{{.Spec.Name}})
{{end}}

var (
    _{{$.Ns}}_def_init_once sync.Once

    {{range $_, $v := .Structs}}
        {{include "structVars" $v}}
    {{end}}

    _{{$.Ns}}_def_field_default__Foo__baz int

    _{{$.Ns}}_def_struct_inits__Foo [1]func(*Foo)
)

func _{{$.Ns}}_def_init() {
    _{{$.Ns}}_def_init_once.Do(func() {
        spec := def.X_getPackageSpec()

        struct_spec__Foo := spec.Struct("Foo")

        field_spec__Foo__baz := struct_spec__Foo.Field("baz")
        _{{$.Ns}}_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

        _{{$.Ns}}_def_struct_inits__Foo = make([]func(*Foo), len(struct_spec__Foo.Inits()))
        for i, f := range struct_spec__Foo.Inits() {
            _{{$.Ns}}_def_struct_inits__Foo[i] = f.(func(*Foo))
        }
    })
}

