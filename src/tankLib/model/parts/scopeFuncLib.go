package parts

type Scope struct {
	code      int //掩码
	direction int //方向
	vision    int //视野
	distance  int //视野距离
}

type ScopeFunc interface {
	directionCalibration(direction int) //方向校准
}
