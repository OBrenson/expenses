package dbaccess

import "fmt"

type NoRowsAffected struct {  
	EntityName string
	EntityId int
	Method string
}

func (e *NoRowsAffected) Error() string {  
	return fmt.Sprint("Entity %s with id %d was not affetcted by method %s", e.EntityName, e.EntityId, e.Method)
}