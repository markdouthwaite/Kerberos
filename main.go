package main

import (
	"fmt"
	"github.com/markdouthwaite/Kerberos/kerberos"
)

func main() {
	counts := []int{0, 0, 0, 0, 0}
	reward := []float32{0, 0, 0, 0, 0}

	b := kerberos.EpsilonBandit{
		N:       5,
		Counts:  counts,
		Reward:  reward,
		Epsilon: 0.2,
	}

	b.UpdateAction(1, 1.0)
	fmt.Printf("%f\n", b.Reward)
	fmt.Printf("%f\n", kerberos.MaxFloat32(b.Reward))
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", b.GetAction())
	}

}
