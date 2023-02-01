package commands

type CommandType string

const (
	ActionType CommandType = "Action"
	SubjectType CommandType = "Subject"
	ValueType CommandType = "Value"
)

type Command interface {
	GetType() CommandType
	GetNext() Command
}

func CheckCommandsChain(head Command) error {
	if head.GetType() != ActionType {
		return &CommandChainError{err: "Commands chain must start from Action"}
	}

	isSub := true
	for h := head.GetNext(); h != nil; {
		if isSub {
			if h.GetType() != SubjectType {
				return &CommandChainError{err: "After Action must Go Subject"}
			}
		} else {
			if h.GetType() != SubjectType {
				return &CommandChainError{err: "After Subject must Go Value"}
			}
		}
		isSub = !isSub
	}
	if !isSub {
		return &CommandChainError{err: "Last Command must be Value"}
	}
	return nil
}

type CommandChainError struct {
	err string
}

func (c *CommandChainError) Error() string {
	return c.err
}