package controller

import (
	"fmt"
	"testing"
)

type intStruct struct {
	val int
}

func TestMapAndSlice(t *testing.T) {
	m := map[string]*intStruct{}
	s := []*intStruct{}
	a := intStruct{2}

	m["a"] = &a
	s = append(s, &a)
	b := m["a"]
	b.val++
	fmt.Println(m["a"], s[0])
}
