package boxes

import "ptypes/ptr/prim"

const (
	UINT = "uint"
	INT  = "int"
)

func (box *IntBox) Int() {

}

func (box *IntBox) Uint() uint {
	conversion, err := box.ToUintOfWidth(1)
	if err != nil {
		return 0
	}
	return conversion.(uint)
}

func (box *IntBox) ToUintOfWidth(width int) (interface{}, error) {
	var blockSize string
	if width == 1 {
		blockSize = ""
	} else {
		blockSize = string(rune(width))
	}
	builtinName := UINT + blockSize
	p := prim.Builtin(builtinName)
	if p == nil {
		return nil, makeError("there is no builtin for the given name " + builtinName)
	}
	conversion := p.Convert(box.Value)
	return conversion, nil
}
