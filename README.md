# spacelist

A beautiful terminal UI for viewing all windows across your Aerospace window manager spaces.

![](./screenshot.png)

## Installation

```bash
brew install magicmark/tap/spacelist
```

## Usage

Simply run:

```bash
spacelist
```

> [!TIP]
> Hotkey spacelist to open in a new terminal in the center of your screen for
> an experience similar to spotlight. WIP but check out `launcher.sh`.

## Roadmap

Checkout out [ROADMAP.md](./ROADMAP.md).

## Features

- Lists all windows organized by workspace
- Real-time filtering by application name
- Beautiful TUI with syntax highlighting using Bubble Tea
- Shows workspace names, application names, and window titles
- Only displays workspaces that contain windows

## How it works

The application:
1. Queries `aerospace list-workspaces --all --json` to get all workspaces
2. For each workspace, runs `aerospace list-windows --workspace <name> --json`
3. Displays results in a filterable TUI using Charm's Bubble Tea library

## Libraries Used

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components (text input)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
