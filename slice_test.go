package concsafe

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestSliceCount(t *testing.T) {
	slice := &Slice{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			slice.Add(fmt.Sprintf("test %d", i))
			wg.Done()
		}(i)

	}

	wg.Wait()

	if slice.Count() != 100 {
		t.Error(fmt.Sprintf("slice.Count() == 100, has %d", slice.Count()))
	}
}

func TestSliceList(t *testing.T) {
	slice := &Slice{}
	var wg sync.WaitGroup
	var sb strings.Builder

	for i := 0; i < 10; i++ {
		wg.Add(1)
		sb.WriteString(fmt.Sprintf("\n%6d | test", i))
		go func(i int) {
			slice.Add("test")
			wg.Done()
		}(i)

	}

	wg.Wait()
	list, err := slice.List()
	test := sb.String()

	if list != test {
		t.Error(fmt.Sprintf("\nwants%s\n has %s", test, list))
	}

	if err != nil {
		t.Error(err)
	}
}
