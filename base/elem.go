package base

import "fmt"

type ElemType int

func (e ElemType) String() string {
	return fmt.Sprintf("%d", e)
}
