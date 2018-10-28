package controller

import (
	"fmt"
	"time"
)

func noob() {
	g := time.Now()
	t := g.Add(1 * time.Hour)
	fmt.Println(t)
	fmt.Println(g)
	d := g.Sub(t)
	fmt.Println(d)
}
