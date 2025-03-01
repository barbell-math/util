package generatortests

//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestStructValImplsIFace
//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestStructPntrImplsIFace
//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestStructBothImplsIFace

//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestIntValImplsIFace
//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestIntPntrImplsIFace
//go:generate ../../bin/ifaceImplCheck -typeToCheck=TestIntBothImplsIFace

type (
	IFace interface {
		Func1() bool
	}

	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr val
	TestStructValImplsIFace struct{}
	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr pntr
	TestStructPntrImplsIFace struct{}
	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr both
	TestStructBothImplsIFace struct{}

	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr val
	TestIntValImplsIFace int
	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr pntr
	TestIntPntrImplsIFace int
	//gen:ifaceImplCheck ifaceName IFace
	//gen:ifaceImplCheck valOrPntr both
	TestIntBothImplsIFace int
)

func (t TestStructValImplsIFace) Func1() bool   { return false }
func (t *TestStructPntrImplsIFace) Func1() bool { return false }
func (t TestStructBothImplsIFace) Func1() bool  { return false }

func (i TestIntValImplsIFace) Func1() bool   { return false }
func (i *TestIntPntrImplsIFace) Func1() bool { return false }
func (i TestIntBothImplsIFace) Func1() bool  { return false }
