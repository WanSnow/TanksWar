package parts

type Engine struct {
	code       int //掩码
	efficiency int //效率
}

type EngineFunc interface {
	cycleInstruction() //周期指令
}
