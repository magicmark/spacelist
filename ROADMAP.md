# Roadmap

### Adaptive Light/Dark Mode
- Implement automatic light/dark mode switching using `lipgloss.AdaptiveColor`
- Reference: https://pkg.go.dev/github.com/charmbracelet/lipgloss#readme-adaptive-colors

### Fix Launcher Centering Logic
- Clean up and fix `launcher.sh` centering logic

### Enhanced Filtering
- Extend filter to support window titles (not just app names)
- Add toggle to switch between app name / window title / both
- Display filter mode in the UI

### Visual Feedback
- Add visual indicators for selected windows
- Consider using symbols (▶, •, ▸) to mark selected item
- Highlight the entire row with background color, not just the text color

### Dynamic Updates
- Add refresh capability
- `r` key to refresh workspace/window list
- Optional auto-refresh with configurable interval
- Show "Refreshing..." indicator during updates

### Sorting Options
- Add sort modes for window list:
  - By app name (alphabetical)
  - By window title (alphabetical)
  - Toggle sort direction (ascending/descending)

### Configuration File
- Support config file (`~/.config/spacelist/config.yml` or similar)
  - Customizable key bindings
  - Color scheme overrides
  - Default filter mode
  - Auto-refresh settings
  - Launch options

### Empty State Improvements
- Better empty state handling:
  - Show helpful message when no workspaces have windows
  - Display ASCII art or logo
  - Show helpful keyboard shortcuts

## Platform Support

### Cross-Platform Launcher
- Make launcher.sh more portable
  - Currently tied to ghostty.app
  - Support other terminals (iTerm2, Kitty, Alacritty, etc.)
  - Better macOS/Linux compatibility

### Window Manager Support
- Explore support for other tiling window managers?
  - i3/Sway
  - yabai
  - Amethyst
  - Abstract window manager interface for portability
