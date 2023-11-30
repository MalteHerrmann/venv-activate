package main

import (
	"os"
)

var (
	// VENVDIR is the directory, where the venv environments are stored.
	VENVDIR = os.Getenv("VENV_DIR")
)

func main() {
	envs := getAvailableEnvs(VENVDIR)
	choice := selectEnv(envs)
	activateEnv(choice)
}
