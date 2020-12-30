package markov

//State interface
type State interface {
	String() string
	IsFinal() bool
}
