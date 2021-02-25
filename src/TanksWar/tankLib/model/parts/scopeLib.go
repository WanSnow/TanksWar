package parts

type DefaultScope struct {
	//code int //掩码
	//direction int //方向,存入redis
	vision   int //视野
	distance int //视野距离
}

type MoreVisionScope struct {
	//code int //掩码
	//direction int //方向,存入redis
	vision   int //视野
	distance int //视野距离
}

type MoreDistanceScope struct {
	//code int //掩码
	//direction int //方向,存入redis
	vision   int //视野
	distance int //视野距离
}

type BalanceVisionAndDistanceScope struct {
	//code int //掩码
	//direction int //方向,存入redis
	vision   int //视野
	distance int //视野距离
}

type NullScope struct {
	//code int //掩码
	//direction int //方向,存入redis
	vision   int //视野
	distance int //视野距离
}

const (
	ScopeDefaultMask = 1 << iota
	MoreVision
	MoreDistance
	BalanceScope
)
