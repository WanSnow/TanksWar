package implement

type Scope struct {
	code      int //掩码
	direction int //方向
	vision    int //视野
	distance  int //视野距离
}

const (
	ScopeDefaultMask = 1 << iota
	MoreVision
	MoreDistance
	BalanceScope
)

func ScopeFactory(code int) Scope {
	switch code {
	case ScopeDefaultMask:
		return Scope{code: ScopeDefaultMask, vision: 90, distance: 200}
	case MoreVision:
		return Scope{code: MoreVision, vision: 150, distance: 200}
	case MoreDistance:
		return Scope{code: MoreDistance, vision: 90, distance: 500}
	case BalanceScope:
		return Scope{code: BalanceScope, vision: 120, distance: 300}
	default:
		return Scope{code: 0, vision: 0, distance: 0}
	}
}
