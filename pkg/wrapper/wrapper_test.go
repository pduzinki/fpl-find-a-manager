package wrapper

import (
	"fmt"
	"testing"
)

func TestDummy(t *testing.T) {
	w := NewWrapper()

	manager, err := w.GetManager(43741)
	if err != nil {
		panic(err)
	}

	fmt.Println(manager)
}
