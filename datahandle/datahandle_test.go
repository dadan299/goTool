package datahandle_test

import (
	"fmt"
	"github.com/dadan299/gotool/datahandle"
	"testing"
)

func TestJoin(t *testing.T) {

	var a = []int{1,2}

	res := datahandle.JoinInt(a,",")

	fmt.Println(res)
}