package main

import (
	"context"
	"fmt"
	"os/exec"
)

var globalPercentage string

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	cmd := exec.Command("cat", "/proc/ram_202001570")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("Error : ", err)
	}
	return fmt.Sprintf(string(output))

}
