package parts

type Wheel struct {
	code  int //掩码
	speed int //速度
}

type WheelFunc interface {
	move(x int, y int)
}
