package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var ctx = context.Background()

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()

			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.table.SetRows(refreshTable(ctx))
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func refreshTable(ctx context.Context) []table.Row {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	var rows []table.Row
	for _, container := range containers {
		rows = append(rows, table.Row{
			container.ID[:10],
			strings.Join(container.Names, ", "),
			container.State,
			container.Image,
		})
	}
	return rows
}

func main() {

	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Names", Width: 30},
		{Title: "State", Width: 12},
		{Title: "Image", Width: 50},
	}

	rows := refreshTable(ctx)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBackground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error runnig: ", err)
		os.Exit(1)
	}
}
