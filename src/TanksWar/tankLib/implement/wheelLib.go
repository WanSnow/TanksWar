package implement

type Wheel struct {
	code  int //掩码
	speed int //速度
}

const (
	WheelLv1 = 1 << iota
	WheelLv2
	WheelLv3
	WheelLv4
)

func WheelFactory(code int) Wheel {
	switch code {
	case WheelLv1:
		return Wheel{code: WheelLv1, speed: 100}
	case WheelLv2:
		return Wheel{code: WheelLv2, speed: 150}
	case WheelLv3:
		return Wheel{code: WheelLv3, speed: 200}
	case WheelLv4:
		return Wheel{code: WheelLv4, speed: 250}
	default:
		return Wheel{code: 0, speed: 0}
	}
}
