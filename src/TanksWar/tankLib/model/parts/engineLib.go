package parts

type Lv1Engine struct {
	//code       int //掩码
	efficiency int //效率
}

type Lv2Engine struct {
	//code       int //掩码
	efficiency int //效率
}

type Lv3Engine struct {
	//code       int //掩码
	efficiency int //效率
}

type Lv4Engine struct {
	//code       int //掩码
	efficiency int //效率
}

type NullEngine struct {
	//code       int //掩码
	efficiency int //效率
}

const (
	EngineLv1 = 1 << iota
	EngineLv2
	EngineLv3
	EngineLv4
)
