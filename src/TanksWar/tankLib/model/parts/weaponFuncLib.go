package parts

type WeaponFunc interface {
	shot()                              //射击
	directionCalibration(direction int) //方向校准
}
