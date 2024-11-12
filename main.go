package main

import (
	"fmt"
	"strings"

	"gitlab.com/scabala/zelligo"
)

type Plugin struct {
	tabs     map[string]bool
	testRuns uint32
	config   map[string]string
	modeLog  map[string]int
}

// Load implements zelligo.ZellijPlugin.
func (p *Plugin) Load(configuration map[string]string) error {
	fmt.Println("Hello from Load")
	return nil
}

// Pipe implements zelligo.ZellijPlugin.
func (p *Plugin) Pipe(message zelligo.PipeMessage) (bool, error) {
	fmt.Println("Hello from Pipe")
	return true, nil
}

// Render implements zelligo.ZellijPlugin.
func (p *Plugin) Render(rows uint32, cols uint32) error {
	// Create a padding to push text towards the bottom
	padding := strings.Repeat("\n", int(rows)-2) // Adjust based on desired position
	message := "Hello, World!"

	// Right-align the message within the available width
	padding += fmt.Sprintf("%*s\n", cols-1, message)
	fmt.Print(padding)

	return nil
}

// Update implements zelligo.ZellijPlugin.
func (p *Plugin) Update(event zelligo.Event) (bool, error) {
	fmt.Println("Hello from Update")
	return true, nil
}

func main() {
	fmt.Println("Hello, World!")
}

func init() {
	zelligo.RegisterPlugin(&Plugin{})
}
