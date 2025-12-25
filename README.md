# Spacelist (sl)

A beautiful terminal UI for viewing all windows across your Aerospace window manager spaces.

## Features

- Lists all windows organized by workspace
- Real-time filtering by application name
- Beautiful TUI with syntax highlighting using Bubble Tea
- Shows workspace names, application names, and window titles
- Only displays workspaces that contain windows

## Installation

```bash
# Build and install
make install

# Or manually
go build -o sl
cp sl /usr/local/bin/sl
```

## Usage

Simply run:

```bash
sl
```

### Controls

- Type to filter windows by application name (case-insensitive)
- `Esc` or `Ctrl+C` to quit

## Requirements

- Go 1.21 or later
- [Aerospace](https://github.com/nikitabobko/AeroSpace) window manager

## How it works

The application:
1. Queries `aerospace list-workspaces --all --json` to get all workspaces
2. For each workspace, runs `aerospace list-windows --workspace <name> --json`
3. Displays results in a filterable TUI using Charm's Bubble Tea library

## Libraries Used

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components (text input)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
# spacelist
