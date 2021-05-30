package main

import (
	"fmt"
	"time"

	"github.com/chengjingtao/alauda-ci/pkg"
)

func main() {
	for {
		fmt.Println("Hello world!")
		fmt.Println(pkg.Add(1, 2))
		time.Sleep(time.Second * 1)
	}
}
