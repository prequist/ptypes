package ptypes

const (
	UINT = "uint"
	INT  = "int"
)

func (box IntBox) Int() *int {
	conversion, err := box.IntOfWidth(1)
	if err != nil {
		return nil
	}
	return conversion.(*int)
}

func (box *IntBox) IntOfWidth(width int) (interface{}, error) {
	parent := Box(*box)
	builtinName := INT + determine(width == 1, "", string(rune(width))).(string)
	p := Builtin(builtinName)
	if p == nil {
		return nil, makeError("there is no builtin for the given name " + builtinName)
	}
	if _, err := parent.CanAssign(*p); !box.HasErasure && err != nil {
		return nil, err
	}
	conversion := p.Convert(Box(*box))
	return conversion, nil
}

func (box *IntBox) Uint() *uint {
	conversion, err := box.UintOfWidth(1)
	if err != nil {
		return nil
	}
	return conversion.(*uint)
}

func (box *IntBox) UintOfWidth(width int) (interface{}, error) {
	parent := Box(*box)
	builtinName := UINT + determine(width == 1, "", string(rune(width))).(string)
	p := Builtin(builtinName)
	if p == nil {
		return nil, makeError("there is no builtin for the given name " + builtinName)
	}
	if _, err := parent.CanAssign(*p); !box.HasErasure && err != nil {
		return nil, err
	}
	conversion := p.Convert(Box(*box))
	return conversion, nil
}

func determine(condition bool, option interface{}, other interface{}) interface{} {
	if condition {
		return option
	} else {
		return other
	}
}
