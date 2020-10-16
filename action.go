package markov

type Action interface {
	String() string
	Exec(Context, State) State
}
