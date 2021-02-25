package parts

import "TanksWar/protocol"

type Wheel interface {
	move(point protocol.Point)
}

func WheelFactory(code int) Wheel {
	switch code {
	case WheelLv1:
		return Lv1Wheel{speed: 100}
	case WheelLv2:
		return Lv2Wheel{speed: 150}
	case WheelLv3:
		return Lv3Wheel{speed: 200}
	case WheelLv4:
		return Lv4Wheel{speed: 250}
	default:
		return NullWheel{speed: 0}
	}
}

/*
todo 坦克移动
*/
func (w Lv1Wheel) move(point protocol.Point) {
	panic("implement me")
}

func (w Lv2Wheel) move(point protocol.Point) {
	panic("implement me")
}

func (w Lv3Wheel) move(point protocol.Point) {
	panic("implement me")
}

func (w Lv4Wheel) move(point protocol.Point) {
	panic("implement me")
}

func (w NullWheel) move(point protocol.Point) {
	panic("implement me")
}
