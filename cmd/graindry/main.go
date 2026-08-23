package main

import (
	"os"

	"github.com/lacsar712/graindry/internal/app"
)

func main() {
	os.Exit(app.RunCLI())
}
