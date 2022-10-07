package main

import (
	"context"
	"fmt"
	"os"

	"github.com/umbe77/dockertui/lib"
	"github.com/umbe77/dockertui/models"

	tea "github.com/charmbracelet/bubbletea"
)



func main() {
	ctx := context.Background()
	containers := lib.RefreshTable(ctx)
	containerModel := models.InitializeContainersModel(containers)
	if err := tea.NewProgram(containerModel, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
