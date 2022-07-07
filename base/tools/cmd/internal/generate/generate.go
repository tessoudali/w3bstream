package generate

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-courier/packagesx"
)

func logCost() func() {
	startedAt := time.Now()

	return func() {
		log.Printf("costs %s", color.GreenString("%0.0f ms", float64(time.Now().Sub(startedAt)/time.Millisecond)))
	}
}

type Generator interface {
	Output(cwd string)
}

func Run(cmd string, createGenerator func(pkg *packagesx.Package) Generator) {
	cwd, _ := os.Getwd()
	start := time.Now()
	pkg, err := packagesx.Load(cwd)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Printf("%s %s: costs %s", cmd, pkg.String(),
			color.GreenString("%dms", time.Since(start)/time.Millisecond))
	}()

	g := createGenerator(pkg)
	g.Output(cwd)
}
