package implement

type Armor struct {
	code    int //掩码
	defence int //防御力
	large   int //大小
}

const (
	ArmorDefaultMask = 1 << iota
	MoreDefence
	MoreSize
	BalanceArmor
)

func ArmorFactory(code int) Armor {
	switch code {
	case ArmorDefaultMask:
		return Armor{code: ArmorDefaultMask, defence: 10, large: 100}
	case MoreDefence:
		return Armor{code: MoreDefence, defence: 20, large: 100}
	case MoreSize:
		return Armor{code: MoreSize, defence: 0, large: 50}
	case BalanceArmor:
		return Armor{code: BalanceArmor, defence: 15, large: 75}
	default:
		return Armor{code: 0, defence: 0, large: 100}
	}
}
