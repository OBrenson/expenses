package dbaccess

import "fmt"

type NoRowsAffected struct {  
	Method string
}

func (e *NoRowsAffected) Error() string {  
	return fmt.Sprintf("Entity was not affetcted by method %s", e.Method)
}

type DbPanicErr struct {
	Method string
	Message string
}

func (e *DbPanicErr) Error() string {  
	return fmt.Sprintf("Panic was catched during %s. Message: %s", e.Method, e.Message)
}