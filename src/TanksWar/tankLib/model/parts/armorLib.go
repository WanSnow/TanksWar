package parts

type DefaultArmor struct {
	//code    int //掩码
	defence int //防御力
	large   int //大小
}

type MoreDefenceArmor struct {
	//code    int //掩码
	defence int //防御力
	large   int //大小
}

type MoreSizeArmor struct {
	//code    int //掩码
	defence int //防御力
	large   int //大小
}

type BalanceDefenceAndSizeArmor struct {
	//code    int //掩码
	defence int //防御力
	large   int //大小
}

type NullArmor struct {
	//code    int //掩码
	defence int //防御力
	large   int //大小
}

const (
	ArmorDefaultMask = 1 << iota
	MoreDefence
	MoreSize
	BalanceArmor
)
