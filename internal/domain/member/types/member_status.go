package types

type MemberStatus int

const (
	ACTIVE MemberStatus = iota
	INACTIVE
	WITHDRAW
)

func (s MemberStatus) String() string {
	switch s {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	case WITHDRAW:
		return "withdraw"
	}

	return ""
}
