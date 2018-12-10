package controller

import (
	"fmt"
    "testing"
)
func TestParse(t *testing.T) {
	m, err := Parse("config")
	if err != nil {
		panic(err)
	}
	for _, i := range m {
		fmt.Println(i.ClusterName)
		fmt.Println(i.Godhcp)
		fmt.Println(i.Govirt)
		fmt.Println(i.Storage)
	}
}
