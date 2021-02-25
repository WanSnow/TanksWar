package parts

type Bullet struct {
	//code     int //掩码
	damage   int //伤害
	boundary int //爆炸范围
}

const (
	BulletDefaultMask = 1 << iota
	MoreDamage
	MoreBoundary
	BalanceBullet
)

func BulletFactory(code int) Bullet {
	switch code {
	case BulletDefaultMask:
		return Bullet{damage: 30, boundary: 50}
	case MoreDamage:
		return Bullet{damage: 50, boundary: 50}
	case MoreBoundary:
		return Bullet{damage: 30, boundary: 100}
	case BalanceBullet:
		return Bullet{damage: 40, boundary: 75}
	default:
		return Bullet{damage: 0, boundary: 0}
	}
}
