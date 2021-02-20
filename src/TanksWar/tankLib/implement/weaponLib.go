package implement

type Weapon struct {
	code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

const (
	WeaponDefaultMask = 1 << iota
	MoreGunShot
	MoreFiringRate
	MoreBulletRate
	BalanceWeapon
)

func WeaponFactory(code int) Weapon {
	switch code {
	case WeaponDefaultMask:
		return Weapon{code: WeaponDefaultMask, gunshot: 400, firingRate: 2000, bulletRate: 100}
	case MoreGunShot:
		return Weapon{code: MoreGunShot, gunshot: 700, firingRate: 2000, bulletRate: 100}
	case MoreFiringRate:
		return Weapon{code: MoreFiringRate, gunshot: 400, firingRate: 1000, bulletRate: 100}
	case MoreBulletRate:
		return Weapon{code: MoreBulletRate, gunshot: 400, firingRate: 2000, bulletRate: 200}
	case BalanceWeapon:
		return Weapon{code: MoreBulletRate, gunshot: 500, firingRate: 1700, bulletRate: 150}
	default:
		return Weapon{code: 0, gunshot: 0, firingRate: -1, bulletRate: -1}
	}
}
