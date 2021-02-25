package parts

type Weapon interface {
	shot()                              //射击
	directionCalibration(direction int) //方向校准
}

func WeaponFactory(code int) Weapon {
	switch code {
	case WeaponDefaultMask:
		return DefaultWeapon{gunshot: 400, firingRate: 2000, bulletRate: 100}
	case MoreGunShot:
		return MoreGunShotWeapon{gunshot: 700, firingRate: 2000, bulletRate: 100}
	case MoreFiringRate:
		return MoreFiringRateWeapon{gunshot: 400, firingRate: 1000, bulletRate: 100}
	case MoreBulletRate:
		return MoreBulletRateWeapon{gunshot: 400, firingRate: 2000, bulletRate: 200}
	case BalanceWeapon:
		return BalanceAllWeapon{gunshot: 500, firingRate: 1700, bulletRate: 150}
	default:
		return NullWeapon{gunshot: 0, firingRate: -1, bulletRate: -1}
	}
}

/*
todo 坦克攻击
*/
func (d DefaultWeapon) shot() {
	panic("implement me")
}

func (d DefaultWeapon) directionCalibration(direction int) {
	panic("implement me")
}

func (d MoreGunShotWeapon) shot() {
	panic("implement me")
}

func (d MoreGunShotWeapon) directionCalibration(direction int) {
	panic("implement me")
}

func (d MoreFiringRateWeapon) shot() {
	panic("implement me")
}

func (d MoreFiringRateWeapon) directionCalibration(direction int) {
	panic("implement me")
}

func (d MoreBulletRateWeapon) shot() {
	panic("implement me")
}

func (d MoreBulletRateWeapon) directionCalibration(direction int) {
	panic("implement me")
}

func (d BalanceAllWeapon) shot() {
	panic("implement me")
}

func (d BalanceAllWeapon) directionCalibration(direction int) {
	panic("implement me")
}

func (d NullWeapon) shot() {
}

func (d NullWeapon) directionCalibration(direction int) {
}
