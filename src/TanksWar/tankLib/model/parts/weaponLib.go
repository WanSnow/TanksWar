package parts

type DefaultWeapon struct {
	//code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type MoreGunShotWeapon struct {
	//code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type MoreFiringRateWeapon struct {
	//code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type MoreBulletRateWeapon struct {
	//code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type BalanceAllWeapon struct {
	//code int //掩码
	//direction  int //方向存入redis
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type NullWeapon struct {
	//code int //掩码
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
