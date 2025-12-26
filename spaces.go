package main

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Workspace represents an aerospace workspace
type ListWorkspaceOutput struct {
	Name string `json:"workspace"`
}

func (w Window) Focus() {
	cmd := exec.Command("aerospace", "focus", "--window-id", strconv.Itoa(w.WindowID))
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatal(fmt.Sprint(err) + string(output))
	}
}

func GetWorkspaces() []Workspace {
	cmd := exec.Command("aerospace", "list-workspaces", "--all", "--json")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var _workspaces []ListWorkspaceOutput
	if err := json.Unmarshal(output, &_workspaces); err != nil {
		log.Fatal(err)
	}

	// Get windows for each workspace
	var workspaces []Workspace
	for _, ws := range _workspaces {
		cmd := exec.Command("aerospace", "list-windows", "--workspace", ws.Name, "--json")
		output, err := cmd.Output()
		if err != nil {
			continue
		}

		var windows []Window
		if err := json.Unmarshal(output, &windows); err != nil {
			continue
		}

		if len(windows) > 0 {
			workspaces = append(workspaces, Workspace{
				Name:    ws.Name,
				Windows: windows,
			})
		}
	}

	return workspaces
}

func (m model) GetVisibleWindows() []Workspace {
	filterText := strings.ToLower(strings.TrimSpace(m.filter.Value()))

	var workspaces []Workspace

	for _, ws := range m.workspaces {
		_workspace := Workspace{Name: ws.Name}

		for _, win := range ws.Windows {
			if strings.TrimSpace(filterText) == "" ||
				strings.TrimSpace(filterText) != "" &&
					strings.Contains(strings.ToLower(win.AppName), filterText) {
				_workspace.Windows = append(_workspace.Windows, win)
			}
		}

		if len(_workspace.Windows) > 0 {
			workspaces = append(workspaces, _workspace)
		}
	}

	return workspaces
}

func (m model) ViewportContent() string {
	workspaces := m.GetVisibleWindows()

	if len(workspaces) == 0 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Render("No windows match your filter.")
	}

	workspaceStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFD700")).
		MarginTop(1)

	// appStyle := lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("#61AFEF"))

	appNameStyle := lipgloss.NewStyle().Width(20).Foreground(lipgloss.Color("#61AFEF")).MarginLeft(2)
	appNameStyleSelected := appNameStyle.Foreground(lipgloss.Color("212")).Bold(true)

	windowTitleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	windowTitleStyleSelected := windowTitleStyle.Foreground(lipgloss.Color("212")).Bold(true)

	var content strings.Builder

	visibleLineIdx := 0

	for _, ws := range workspaces {
		var workspaceOutput strings.Builder
		for _, win := range ws.Windows {

			appName := appNameStyle.Render(win.AppName)
			windowTitle := windowTitleStyle.Render(win.WindowTitle)

			if visibleLineIdx == m.selectedIdx {
				appName = appNameStyleSelected.Render(win.AppName)
				windowTitle = windowTitleStyleSelected.Render(win.WindowTitle)
			}

			lineStr := lipgloss.JoinHorizontal(0, appName, windowTitle)
			workspaceOutput.WriteString(lineStr)
			workspaceOutput.WriteString("\n")

			visibleLineIdx = visibleLineIdx + 1
		}

		content.WriteString(workspaceStyle.Render(fmt.Sprintf("Workspace %s", ws.Name)))
		content.WriteString("\n")
		content.WriteString(workspaceOutput.String())
	}

	return content.String()
}
