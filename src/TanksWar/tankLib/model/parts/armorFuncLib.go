package parts

type Armor interface {
	BeHit()
}

func ArmorFactory(code int) Armor {
	switch code {
	case ArmorDefaultMask:
		return DefaultArmor{defence: 10, large: 100}
	case MoreDefence:
		return MoreDefenceArmor{defence: 20, large: 100}
	case MoreSize:
		return MoreSizeArmor{defence: 0, large: 50}
	case BalanceArmor:
		return BalanceDefenceAndSizeArmor{defence: 15, large: 75}
	default:
		return NullArmor{defence: 0, large: 100}
	}
}

/*
todo 装甲反馈
*/
func (a DefaultArmor) BeHit() {

}

func (a MoreDefenceArmor) BeHit() {

}

func (a MoreSizeArmor) BeHit() {

}

func (a BalanceDefenceAndSizeArmor) BeHit() {

}

func (a NullArmor) BeHit() {

}
