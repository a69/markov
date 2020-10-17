package markov

// TransitionFunc transition function
type TransitionFunc func(state State) (newState State)

//Markov struct
type Markov struct {
	State State
	//StateTransitionTable
	TransitionFuncs map[State]map[Action]TransitionFunc

	//StateTransitionProbabilityTable
	Probabilities map[State]map[Action]float64
}

func NewMarcov(TransitionFuncs map[State]map[Action]TransitionFunc, Probabilities map[State]map[Action]float64, initState State) *Markov {
	return &Markov{
		State:           initState,
		TransitionFuncs: TransitionFuncs,
		Probabilities:   Probabilities,
	}
}

func (markov Markov) Lottery(c Context) error {
	for {
		select {
		case <-c.Done():
			return c.Err()
		default:
			if markov.State.IsFinal() {
				return nil
			}
			a, err := RandomAction(markov.Probabilities[markov.State])
			if err != nil {
				return err
			}
			markov.State = markov.TransitionFuncs[markov.State][a](markov.State)

		}
	}
}
