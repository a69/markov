package markov

type Markov struct {
	State State
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
			a, err := RandomAction(markov.State)
			if err != nil {
				return err
			}
			markov.State = a.Exec(c, markov.State)

		}
	}
}
