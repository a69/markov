package markov

type State interface {
	String() string
	IsFinal() bool
}
