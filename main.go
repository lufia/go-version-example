package main

import (
	"fmt"
	"log"
	"runtime/debug"

	_ "github.com/pkg/errors"
)

func main() {
	log.SetFlags(0)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("can't get BuildInfo")
	}
	fmt.Printf("Main = %s\n", info.Main.Version)
	for _, m := range info.Deps {
		fmt.Printf("Deps[%s] = %+v\n", m.Path, m.Version)
	}
}
