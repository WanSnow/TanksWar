package parts

import "TanksWar/protocol"

type Scope interface {
	directionCalibration(direction int) //方向校准
	findEnemy(enemy []protocol.Point)   //发现敌人
	findBullet(bullet []protocol.Point) //发现子弹
}

func ScopeFactory(code int) Scope {
	switch code {
	case ScopeDefaultMask:
		return DefaultScope{vision: 90, distance: 200}
	case MoreVision:
		return MoreVisionScope{vision: 150, distance: 200}
	case MoreDistance:
		return MoreDistanceScope{vision: 90, distance: 500}
	case BalanceScope:
		return BalanceVisionAndDistanceScope{vision: 120, distance: 300}
	default:
		return NullScope{vision: 0, distance: 0}
	}
}

/*
TODO 视镜反馈及其校准
*/
func (d DefaultScope) directionCalibration(direction int) {
	panic("implement me")
}

func (d DefaultScope) findEnemy(enemy []protocol.Point) {
	panic("implement me")
}

func (d DefaultScope) findBullet(bullet []protocol.Point) {
	panic("implement me")
}

func (d MoreVisionScope) directionCalibration(direction int) {
	panic("implement me")
}

func (d MoreVisionScope) findEnemy(enemy []protocol.Point) {
	panic("implement me")
}

func (d MoreVisionScope) findBullet(bullet []protocol.Point) {
	panic("implement me")
}

func (d MoreDistanceScope) directionCalibration(direction int) {
	panic("implement me")
}

func (d MoreDistanceScope) findEnemy(enemy []protocol.Point) {
	panic("implement me")
}

func (d MoreDistanceScope) findBullet(bullet []protocol.Point) {
	panic("implement me")
}

func (d BalanceVisionAndDistanceScope) directionCalibration(direction int) {
	panic("implement me")
}

func (d BalanceVisionAndDistanceScope) findEnemy(enemy []protocol.Point) {
	panic("implement me")
}

func (d BalanceVisionAndDistanceScope) findBullet(bullet []protocol.Point) {
	panic("implement me")
}

func (d NullScope) directionCalibration(direction int) {
}

func (d NullScope) findEnemy(enemy []protocol.Point) {
}

func (d NullScope) findBullet(bullet []protocol.Point) {
}
