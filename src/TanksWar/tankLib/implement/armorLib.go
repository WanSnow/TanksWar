package implement

type Armor struct {
	code    int //掩码
	defence int //防御力
	large   int //大小
}

const (
	DefaultMask = 1 << iota
	MoreDefence
	MoreSize
	Balance
)
