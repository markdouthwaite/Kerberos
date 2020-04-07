package kerberos


type Bandit interface {
	GetAction() int
	UpdateAction(int, float32)
}


type StringBandit interface {
	ToString() string
}

