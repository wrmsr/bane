package gogen

//

type Raw struct {
	decl
	stmt
	expr
	Raw string
}

func (r Raw) isNode() {}
