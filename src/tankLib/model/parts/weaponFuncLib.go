package parts

type Weapon struct {
	code       int //掩码
	direction  int //方向
	gunshot    int //射程
	firingRate int //射速
	bulletRate int //炮弹飞行速度
}

type WeaponFunc interface {
	shot()                              //射击
	directionCalibration(direction int) //方向校准
}
