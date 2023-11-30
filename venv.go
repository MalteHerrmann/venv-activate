package main

import (
	"fmt"
	"github.com/nexidian/gocliselect"
	"log"
	"os"
	"os/exec"
)

// venv holds all necessary information about the given virtual environment.
//
// TODO: Implement and use this struct.
type venv struct {
	Name string
}

// getAvailableEnvs scans the given directory and looks for any venv environments
// that can be activated.
func getAvailableEnvs(folder string) []string {
	entries, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal("could not read directory contents", err)
	}

	var envs []string
	for _, entry := range entries {
		if entry.IsDir() {
			envs = append(envs, entry.Name())
		}
	}

	fmt.Println("Found environments: ", envs)
	return envs
}

// selectEnv lets the user select the environment they want to activate.
func selectEnv(envs []string) string {
	menu := gocliselect.NewMenu("Choose which environment to activate")
	for _, env := range envs {
		menu.AddItem(env, env)
	}
	return menu.Display()

}

// activateEnv activates the given environment.
//
// FIXME: This is not actually activating the environment??
func activateEnv(env string) {
	fmt.Printf("Activating %s\n", env)
	cmd := exec.Command("bash", "-s", fmt.Sprintf("source %s/%s/bin/activate", VENVDIR, env))
	err := cmd.Run()
	if err != nil {
		log.Fatal("could not execute command", err)
	}
}
