package cutil

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cuisi521/go-cs-core/tool/util/conv"
	"github.com/cuisi521/go-cs-core/tool/util/crand"
)

func TestConvert(t *testing.T) {
	fmt.Println(conv.Int("1"))
	fmt.Println(conv.Str(2))
	fmt.Println(conv.Int8("3"))
	fmt.Println(conv.Int64("1234567891234567890"))
	b := []byte("123")
	fmt.Println(conv.Str(b))
}

func TestSvq(t *testing.T) {
	s := [5]int{1, 2, 3, 4, 5}
	fmt.Println(s[1])
	s1 := s[1:3]
	fmt.Println(s1)
	s[1] = 20
	s[2] = 30
	fmt.Println(s1)
}

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemType := v.Elem()
	n := elemType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemType)
		return
	}
	fmt.Printf("%s's method set:\n", elemType)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemType.Method(j).Name)
	}
	fmt.Printf("\n")
}

type Interface interface {
	M1()
	M2()
}
type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func TestInterface(ts *testing.T) {
	var t T
	var pt *T
	DumpMethodSet(&t)
	DumpMethodSet(&pt)
	DumpMethodSet((*Interface)(nil))
}

func TestTry(t *testing.T) {
	var ts *string
	err := Try(func() {
		var tv string = "1"
		ts = &tv
		fmt.Println(*ts)
		// panic("err.ro")
	})
	if err != nil {
		fmt.Println("sssssssss")
	}
}

func TestTryCache(t *testing.T) {
	var ts *string
	TryCatch(func() {
		var tv string = "1"
		ts = &tv
		fmt.Println(*ts)
		panic("err.ro")
	}, func(exception error) {
		panic("err.ro")
		if exception != nil {
			fmt.Println("cvv1")
		}

	})

}

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(crand.RS(16))
		}()
	}

}
