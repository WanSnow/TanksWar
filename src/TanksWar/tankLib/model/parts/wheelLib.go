package parts

type Lv1Wheel struct {
	//code  int //掩码
	speed int //速度
}

type Lv2Wheel struct {
	//code  int //掩码
	speed int //速度
}

type Lv3Wheel struct {
	//code  int //掩码
	speed int //速度
}

type Lv4Wheel struct {
	//code  int //掩码
	speed int //速度
}

type NullWheel struct {
	//code  int //掩码
	speed int //速度
}

const (
	WheelLv1 = 1 << iota
	WheelLv2
	WheelLv3
	WheelLv4
)
