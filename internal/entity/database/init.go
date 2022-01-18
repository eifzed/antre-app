package reservation

import (
	"github.com/go-xorm/xorm"
)

type Connection struct {
	Master *xorm.Engine
	Slave  *xorm.Engine
}

func ConnetDB(dataSource string) (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", dataSource)
}
