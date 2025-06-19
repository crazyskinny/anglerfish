package anglerfish

type StateMachine interface {
	IsStateChanged() bool
	StateEnter(state string) error
	StateExit(state string) error
}

type DefaultStateMachineImpl struct {
}
