package builder

//

type Raw struct {
	stmt
	expr
	Raw string
}

func (r Raw) isNode() {}
