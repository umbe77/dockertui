package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/api/types"
)

var (
	listStyle = lipgloss.NewStyle().
		Margin(2, 2).
		Padding(2, 2).
		BorderForeground(lipgloss.Color("5")).
		BorderBackground(lipgloss.Color("0")).
		BorderStyle(lipgloss.RoundedBorder())
)

type containerItem struct {
	id     string
	names  string
	image  string
	state  string
	status string
}

func (i containerItem) Title() string { return i.id }
func (i containerItem) Description() string {
	return lipgloss.JoinHorizontal(lipgloss.Left,
		i.names, "\t",
		i.image, "\t",
		i.status, "\t",
		i.state)
}
func (i containerItem) FilterValue() string { return i.image }

type ContainersModel struct {
	list list.Model
}

func InitializeContainersModel(containers []types.Container) ContainersModel {
	var items = []list.Item{}

	for _, item := range containers {
		items = append(items, containerItem{
			id:     item.ID[:10],
			names:  strings.Join(item.Names, ", "),
			image:  item.Image,
			state:  item.State,
			status: item.Status,
		})
	}

	m := ContainersModel{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}

	m.list.Title = "Containers"

	return m
}

func (m ContainersModel) Init() tea.Cmd {
	return nil
}

func (m ContainersModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, (msg.Height-v)/2)
	case tea.KeyMsg:

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ContainersModel) View() string {
	return listStyle.Render(m.list.View())
}
