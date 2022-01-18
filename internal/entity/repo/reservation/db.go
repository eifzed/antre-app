package reservation

import (
	"github.com/go-xorm/xorm"
)

type Connection struct {
	Master *xorm.Engine
	Slave  *xorm.Engine
}
type Conn struct {
	DB *Connection
	// Gocrypt *gocrypt.Option
}

func ConnetDB(dataSource string) (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", dataSource)
}

func NewDBConnection(conn *Connection) *Conn {
	return &Conn{
		DB: conn,
	}
}
