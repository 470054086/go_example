package Inc

type SexType int

const (
	SexMan   SexType = 1
	SexWoman SexType = 2
)

func (s SexType) String() string {
	switch s {
	case SexMan:
		return "男"
	case SexWoman:
		return "女"
	default:
		panic("Sex UnKNOWN")
	}
}
