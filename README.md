# spacelist

A beautiful terminal UI for viewing all windows across your Aerospace window manager spaces.

![](./screenshot.png)

## Features

- Lists all windows organized by workspace
- Real-time filtering by application name
- Beautiful TUI with syntax highlighting using Bubble Tea
- Shows workspace names, application names, and window titles
- Only displays workspaces that contain windows

## Installation

### Homebrew (recommended)

```bash
brew install magicmark/tap/spacelist
```

### From source

```bash
# Build and install
make install

# Or manually
go build -o spacelist
cp spacelist /usr/local/bin/spacelist
```

## Usage

Simply run:

```bash
spacelist
```

## FAQs

### **`Apple could not verify...`**

Run this to launch spacelist from a CLI wrapper:

```bash
xattr -d com.apple.quarantine /opt/homebrew/bin/spacelist
```

### Controls

- Type to filter windows by application name (case-insensitive)
- `Esc` or `Ctrl+C` to quit
- `Enter` to focus the selected window and quit spacelist

## How it works

The application:
1. Queries `aerospace list-workspaces --all --json` to get all workspaces
2. For each workspace, runs `aerospace list-windows --workspace <name> --json`
3. Displays results in a filterable TUI using Charm's Bubble Tea library

## Libraries Used

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components (text input)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
