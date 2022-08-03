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
	fmt.Printf("Main = %s (%s)\n", info.Main.Version, info.Main.Sum)
	for _, m := range info.Deps {
		fmt.Printf("Deps[%s] = %s (%s)\n", m.Path, m.Version, m.Sum)
	}
	fmt.Println("Settings:")
	for _, s := range info.Settings {
		fmt.Printf("\t%s=%s\n", s.Key, s.Value)
	}
}
