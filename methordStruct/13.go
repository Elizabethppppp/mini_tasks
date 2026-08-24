package methordStruct

type Account struct {
	owner string
	start int
}

func NewAccount(owner string, start int) (*Account, bool) {
	if start < 0 || owner == "" {
		return nil, false
	}
	return &Account{owner: owner, start: start}, true
}
