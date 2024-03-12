package accounts

type Type int

var Names = map[Type]string{
	Email:   "Email",
	Account: "account",
}

const (
	Email Type = iota
	Account
)

func (typ Type) String() string {
	return Names[typ]
}

// Int returns the int value of the LoginType
func (typ Type) Int() int {
	return int(typ)
}

// UInt returns the int value of the LoginType
func (typ Type) UInt() uint8 {
	return uint8(typ)
}
