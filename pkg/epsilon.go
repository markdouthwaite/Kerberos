package pkg

import (
	"fmt"
	"math/rand"
)


// EpsilonBandit - the simplest of the bandits. A single tunable parameter, and basic reward averaging.
type EpsilonBandit struct {
	N int
	Counts []int
	Reward []float32
	Epsilon float32
}


// toString - get the string representation of the Bandit.
func (b *EpsilonBandit) ToString() string {
	return fmt.Sprintf("%d %f\n", b.N, b.Epsilon)
}


// getAction - choose an action to perform.
func (b *EpsilonBandit) GetAction() int {
	if rand.Float32() > b.Epsilon {
		return ArgMaxFloat32(b.Reward)
	} else {
		return rand.Intn(b.N)
	}
}


// updateAction - update the internal state of the bandit with new information.
func (b *EpsilonBandit) UpdateAction(action int, reward float32) {
	b.Counts[action] += 1
	k := float32(b.Counts[action])
	b.Reward[action] = ((k - 1) / k) * b.Reward[action] + (1 / k) * reward
}
