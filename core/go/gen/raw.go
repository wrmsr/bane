package gen

//

type Raw struct {
	decl
	stmt
	expr
	type_
	Raw string
}

func (r Raw) isNode() {}
