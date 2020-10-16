package markov

type State interface {
	String() string
	IsFinal() bool
	GetProbability() map[Action]float64
}
