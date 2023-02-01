package commands

type Action struct {
	Next Command
}

func (a Action) GetType() CommandType {
	return ActionType
}

func (a Action) GetNext() Command {
	return a.Next
}

type Add struct {
	Action
}

type Delete struct {
	Action
}

type Get struct {
	Action
}