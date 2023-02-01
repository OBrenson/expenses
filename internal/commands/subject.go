package commands

type Subject struct {
	Next Command
}

func (s Subject) GetType() CommandType {
	return SubjectType
}

func (s Subject) GetNext() Command {
	return s.Next
}

type Sum struct {
	Subject
}

type Period struct {
	Subject
}