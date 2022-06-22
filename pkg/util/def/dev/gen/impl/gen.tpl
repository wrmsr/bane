{{$Ns := .Ns}}

{{define "structVars"}}
{{end}}

var (
    _{{$Ns}}_def_init_once sync.Once

    {{range $k, $v := .Pkg.Structs}}
        {{$k}}
    {{end}}

    _{{$Ns}}_def_field_default__Foo__baz int

    _{{$Ns}}_def_struct_inits__Foo [1]func(*Foo)
)

func _{{$Ns}}_def_init() {
    _{{$Ns}}_def_init_once.Do(func() {
        spec := def.X_getPackageSpec()

        struct_spec__Foo := spec.Struct("Foo")

        field_spec__Foo__baz := struct_spec__Foo.Field("baz")
        _{{$Ns}}_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

        _{{$Ns}}_def_struct_inits__Foo = make([]func(*Foo), len(struct_spec__Foo.Inits()))
        for i, f := range struct_spec__Foo.Inits() {
            _{{$Ns}}_def_struct_inits__Foo[i] = f.(func(*Foo))
        }
    })
}

