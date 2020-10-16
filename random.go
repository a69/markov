package markov

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomAction(s State) (Action, error) {
	random := rand.Float64()
	sum := 0.0
	for action, prob := range s.GetProbability() {
		if prob == 0 {
			continue
		}
		sum += prob
		if sum > random {
			return action, nil
		}
	}
	return nil, errors.New("action not find")
}
