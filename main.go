package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/superstarryeyes/bit/ansifonts"
)

var header = func() []string {
	font, err := ansifonts.LoadFont("rasterforge")
	if err != nil {
		log.Fatal(err)
	}

	options := ansifonts.RenderOptions{
		CharSpacing:       1,
		TextColor:         "#d42828",
		ScaleFactor:       0.5,
		GradientDirection: ansifonts.LeftRight,
		GradientColor:     "#ed6d38",
		UseGradient:       true,
	}

	return ansifonts.RenderTextWithOptions("spacelist", font, options)
}()

type Window struct {
	AppName     string `json:"app-name"`
	WindowID    int    `json:"window-id"`
	WindowTitle string `json:"window-title"`
}

type Workspace struct {
	Name    string
	Windows []Window
}

type model struct {
	workspaces  []Workspace
	ready       bool
	selectedIdx int
	viewport    viewport.Model
	filter      textinput.Model
	windowID    string
}

func initModel(windowID string) model {
	ti := textinput.New()
	ti.Placeholder = "Filter by app name..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return model{
		filter:      ti,
		ready:       false,
		selectedIdx: -1,
		windowID:    windowID,
		workspaces:  GetWorkspaces(),
	}
}

func (m model) Init() tea.Cmd {
	if m.windowID == "" {
		return tea.SetWindowTitle("spacelist")
	} else {
		return tea.SetWindowTitle("spacelist-" + m.windowID)
	}
}

func (m model) GetVisibleWindowsCount() int {
	numVisibleWindows := 0
	for _, ws := range m.GetVisibleWindows() {
		numVisibleWindows += len(ws.Windows)
	}
	return numVisibleWindows
}

func (m model) GetSelectedWindow() *Window {
	currIdx := 0
	for _, ws := range m.GetVisibleWindows() {
		for _, window := range ws.Windows {
			if currIdx == m.selectedIdx {
				return &window
			}
			currIdx = currIdx + 1
		}
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()

		if k == "ctrl+c" || k == "esc" {
			return m, tea.Quit
		} else if k == "up" {
			if m.selectedIdx <= 0 {
				m.selectedIdx = m.GetVisibleWindowsCount() - 1
			} else {
				m.selectedIdx = m.selectedIdx - 1
			}
			m.viewport.SetContent(m.ViewportContent())
		} else if k == "down" {
			if m.selectedIdx >= m.GetVisibleWindowsCount()-1 {
				m.selectedIdx = 0
			} else {
				m.selectedIdx = m.selectedIdx + 1
			}
			m.viewport.SetContent(m.ViewportContent())
		} else if k == "enter" {
			window := m.GetSelectedWindow()
			if window != nil {
				window.Focus()
				return m, tea.Quit
			}
		} else if m.ready {
			m.filter, cmd = m.filter.Update(msg)
			m.selectedIdx = -1
			m.viewport.SetContent(m.ViewportContent())
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.ViewportContent())
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	filterBoxStyle := lipgloss.NewStyle().
		MarginTop(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#d42828"))

	filterBox := filterBoxStyle.Render(m.filter.View())

	return lipgloss.JoinVertical(lipgloss.Left, append(header, filterBox)...)
}

func (m model) footerView() string {
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Italic(true)

	separatorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#444444"))

	navText := "Press Enter to focus on window"

	if m.selectedIdx == -1 {
		navText = "Use ↑/↓ to select window"
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		footerStyle.Render(navText),
		separatorStyle.Render(" • "),
		footerStyle.Render("Press Esc or Ctrl+C to quit"),
	)
}

func main() {
	windowID := flag.String("id", "", "Window ID for identification")
	flag.Parse()

	if *windowID != "" {
		fmt.Printf("\033]0;spacelist-%s\007", *windowID)
	}

	state := initModel(*windowID)

	p := tea.NewProgram(
		state,
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
