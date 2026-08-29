# Requirements

## MVP Scope

### Main Components (Architecture)

1. GUI Client - Minimal ui, text editor with lsp support. Choosed not native UI Candidates: (Flutter, Wails, Fyne, GioUI)
2. Backend - pure go (data structure, file system operation, etc.)
3. gRPC + Proto communication
4. File structure: Clean architecture

### Functional Requirements

1. Open file, Show contents of file.
2. Unicode support.
3. Write to the file anywhere (prepend, append, write in the middle).
4. Integrate with OS (file picker, light/dark theme).
5. Color schemes.
6. Simple auto-complete (word-based).
7. CUA keybinds (Ctrl+A,Z,X,C,V,W,T,F,S).
8. Tabs.
9. Use key positions for shortcuts.
10. Config file.
11. GUI to edit config file.
12. Undo/redo.

### Non-function Requirements

#### Minimal system reqirements

1. 2 core CPU processor
2. 100mb RAM

#### Load Time

Startup time < 20ms

#### App Size

< 1gb

## Refinement of the scope

1. Use fyne to make a golang gui
2. Without gRPC
3. Editor with textarea, basic functions (save, edit, open)
4. As text buffer use primitive methods.
