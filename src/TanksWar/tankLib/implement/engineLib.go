package implement

type Engine struct {
	code       int //掩码
	efficiency int //效率
}

const (
	EngineLv1 = 1 << iota
	EngineLv2
	EngineLv3
	EngineLv4
)

func EngineFactory(code int) Engine {
	switch code {
	case EngineLv1:
		return Engine{code: EngineLv1, efficiency: 1000}
	case EngineLv2:
		return Engine{code: EngineLv2, efficiency: 800}
	case EngineLv3:
		return Engine{code: EngineLv3, efficiency: 500}
	case EngineLv4:
		return Engine{code: EngineLv4, efficiency: 200}
	default:
		return Engine{code: 0, efficiency: -1}
	}
}
