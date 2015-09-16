package contracts

import (
	"time"
)

type Choosable struct {
	Id     int
	Owner  int
	Voters []int
	Type   ChoosableType

	DataType ChoosableDataType
	Data     interface{}
}

//generic choosables
type FreeText struct {
	Text string
}

type ExactDate struct {
	Date time.Time
}
