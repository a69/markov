package markov_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/a69/markov"
)

type ActionTest uint

const (
	Action1 ActionTest = 1
	Action2 ActionTest = 2
)

type StateTest uint

const (
	State1 StateTest = 1
	State2 StateTest = 2
	State3 StateTest = 3
)

func (s StateTest) IsFinal() bool {
	return s == State3
}
func (s StateTest) String() string {
	return fmt.Sprintf("%d", s)
}

func (a ActionTest) String() string {
	return fmt.Sprintf("%d", a)
}

func Goto1(state markov.State) markov.State {

	fmt.Println("State:", state.String())
	return State1
}
func Goto2(state markov.State) markov.State {

	fmt.Println("State:", state.String())
	return State2
}
func Goto3(state markov.State) markov.State {

	fmt.Println("State:", state.String())
	return State3
}

func TestMarkov(t *testing.T) {
	m := markov.NewMarcov(map[markov.State]map[markov.Action]markov.TransitionFunc{
		State1: map[markov.Action]markov.TransitionFunc{
			Action1: Goto2,
			Action2: Goto1,
		},
		State2: map[markov.Action]markov.TransitionFunc{
			Action1: Goto1,
			Action2: Goto3,
		},
	},
		map[markov.State]map[markov.Action]float64{
			State1: map[markov.Action]float64{
				Action1: 0.5,
				Action2: 0.5,
			},
			State2: map[markov.Action]float64{
				Action1: 0.5,
				Action2: 0.5,
			},
		},
		State1)
	err := m.Lottery(markov.Context{Context: context.Background()})
	if err != nil {
		t.Error(err)
	}
}
