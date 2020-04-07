package pkg


type Bandit interface {
	getAction() int
	updateAction(int, float32)
}


type StringBandit interface {
	toString() string
}

