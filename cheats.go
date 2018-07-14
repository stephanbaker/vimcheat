package main

import "fmt"

// CheatGroup describes a named set of cheats
type CheatGroup struct {
	Name   string
	Cheats []Cheat
}

// Cheat describes a command
type Cheat struct {
	Command     string
	Description string
}

func (c Cheat) String() string {
	return fmt.Sprintf("%s - %s", c.Command, c.Description)
}

// Cheat Groups
var global = CheatGroup{
	Name: "Global",
	Cheats: []Cheat{
		Cheat{":help keyword", "open help for keyword"},
		Cheat{":saveas file", "save file as"},
		Cheat{":close", "close current pane"},
		Cheat{"K", "open man page for word under the cursor"},
	},
}

var cursorMovement = CheatGroup{
	Name: "Cursor movement",
	Cheats: []Cheat{
		Cheat{"h", "move cursor left"},
		Cheat{"j", "move cursor down"},
		Cheat{"k", "move cursor up"},
		Cheat{"l", "move cursor right"},
		Cheat{"H", "move to the top of the screen"},
		Cheat{"M", "move to the middle of the screen"},
		Cheat{"L", "move to the bottom of the screen"},
		Cheat{"w", "jump forwards to the start of a word"},
		Cheat{"W", "jump forwards to the start of a word (words can contain punctuation)"},
		Cheat{"e", "jump forwards to the end of a word"},
		Cheat{"E", "jump forwards to the end of a word (words can contain punctuation)"},
		Cheat{"b", "jump backwards to the start of a word"},
		Cheat{"B", "jump backwards to the start of a word (words can contain punctuation)"},
		Cheat{"%", "move to matching character (default supported pairs: '()', '{}', '[]' - use :h matchpairs in vim for more info)"},
		Cheat{"0", "jump to the start of the line"},
		Cheat{"^", "jump to the first non-blank character of the line"},
		Cheat{"$", "jump to the end of the line"},
		Cheat{"g_", "jump to the last non-blank character of the line"},
		Cheat{"gg", "go to the first line of the document"},
		Cheat{"G", "go to the last line of the document"},
		Cheat{"5G", "go to line 5"},
		Cheat{"fx", "jump to the next occurrence of character x"},
		Cheat{"tx", "jump to before next occurrence of character x"},
		Cheat{"Fx", "jump to previous occurrence of character x"},
		Cheat{"Tx", "jump to after previous occurrence of character x"},
		Cheat{";", "repeat previous f, t, F, or T movement"},
		Cheat{",", "repeat previous f, t, F, or T movement backwards"},
		Cheat{"}", "jump to the next paragraph (or function/block, when editing code)"},
		Cheat{"{", "jump to the previous paragraph (or function/block, when editing code)"},
		Cheat{"zz", "center cursor on screen"},
		Cheat{"Ctrl + z", "move screen down one line (without moving cursor)"},
		Cheat{"Ctrl + y", "move screen up one line (without moving cursor)"},
		Cheat{"Ctrl + b", "move back one full screen"},
		Cheat{"Ctrl + f", "move forward one full screen"},
		Cheat{"Ctrl + d", "move forward half a screen"},
		Cheat{"Ctrl + u", "move back half a screen"},
	},
}

var insertMode = CheatGroup{
	Name: "Insert mode - inserting/appending text",
	Cheats: []Cheat{
		Cheat{"i", "insert before the cursor"},
		Cheat{"I", "insert at the beginning of the line"},
		Cheat{"a", "insert (append after the cursor"},
		Cheat{"A", "insert (append) at the end of the line"},
		Cheat{"o", "append (open) a new line below the current line"},
		Cheat{"O", "append (open) a new line above the current line"},
		Cheat{"ea", "insert (append) at the end of the word"},
		Cheat{"Esc", "exit insert mode"},
	},
}

var workingWithMultipleFiles = CheatGroup{
	Name: "Working with multiple files",
	Cheats: []Cheat{
		Cheat{":e file", "edit a file in a new buffer"},
		Cheat{":bnext", "go to the next buffer"},
		Cheat{":bn", "go to the next buffer"},
		Cheat{":bp", "go to the previous buffer"},
		Cheat{":bd", "delete a buffer (close a file)"},
		Cheat{":ls", "list all open buffers"},
		Cheat{":sp file", "open a file in a new buffer and split window"},
		Cheat{":vsp fail", "open a file in a buffer and vertically split window"},
		Cheat{"Ctrl + ws", "split window"},
		Cheat{"Ctrl + ww", "switch window"},
		Cheat{"wq", "quit a window"},
		Cheat{"wv", "split window vertically"},
		Cheat{"Ctrl + wh", "move cursor to the left window (vertical split)"},
		Cheat{"Ctrl + wl", "move cursor to the right window (vertical split)"},
		Cheat{"Ctrl + wj", "move cursor to the window below (horizontal split)"},
		Cheat{"Ctrl + wk", "move cursor to the window above (horizontal split)"},
	},
}

var editing = CheatGroup{
	Name: "Editing",
	Cheats: []Cheat{
		Cheat{"r", "replace a single character"},
		Cheat{"J", "join line below to the current one with one space in between"},
		Cheat{"gJ", "join line below to the current one without space in between"},
		Cheat{"gwip", "reflow paragraph"},
		Cheat{"cc", "change (replace) entire line"},
		Cheat{"c$", "change (replace) to the end of the line"},
		Cheat{"ciw", "change (replace) entire word"},
		Cheat{"cw", "change (replace) to the end of the word"},
		Cheat{"s", "delete character and substitute text"},
		Cheat{"S", "delete line and substitute text (same as cc)"},
		Cheat{"xp", "transpose two letters (delete and paste)"},
		Cheat{"u", "undo"},
		Cheat{"Ctrl + r", "redo"},
		Cheat{".", "repeat last command"},
	},
}

var markingText = CheatGroup{
	Name: "Marking text (visual mode)",
	Cheats: []Cheat{
		Cheat{"v", "start visual mode, mark lines, then do a command (like y-yank)"},
		Cheat{"V", "start linewise visual mode"},
		Cheat{"o", "move to other end of marked area"},
		Cheat{"Ctrl + v", "start visual block mode"},
		Cheat{"O", "move to other corner of block"},
		Cheat{"aw", "mark a word"},
		Cheat{"ab", "a block with ()"},
		Cheat{"aB", "a block with {}"},
		Cheat{"ib", "inner block with ()"},
		Cheat{"iB", "inner block with {}"},
		Cheat{"Esc", "exit visual mode"},
	},
}

var visualCommands = CheatGroup{
	Name: "Visual commands",
	Cheats: []Cheat{
		Cheat{">", "shift text right"},
		Cheat{"<", "shift text left"},
		Cheat{"y", "yank (copy) marked text"},
		Cheat{"d", "delete marked text"},
		Cheat{"~", "switch case"},
	},
}

var registers = CheatGroup{
	Name: "Registers",
	Cheats: []Cheat{
		Cheat{":reg", "show registers content"},
		Cheat{"\"xy", "yank into register x"},
		Cheat{"\"xp", "paste contents of register x"},
	},
}

var marks = CheatGroup{
	Name: "Marks",
	Cheats: []Cheat{
		Cheat{":marks", "list of marks"},
		Cheat{"ma", "set current position for mark A"},
		Cheat{"`a", "jump to position of mark A"},
		Cheat{"y`a", "yank text to position of mark A"},
	},
}

var macros = CheatGroup{
	Name: "Macros",
	Cheats: []Cheat{
		Cheat{"qa", "record macro a"},
		Cheat{"q", "stop recording macro"},
		Cheat{"@a", "run macro a"},
		Cheat{"@@", "rerun last run macro"},
	},
}

var cutAndPaste = CheatGroup{
	Name: "Cut and paste",
	Cheats: []Cheat{
		Cheat{"yy", "yank (copy) a line"},
		Cheat{"2yy", "yank (copy) 2 lines"},
		Cheat{"yw", "yank (copy) the characters of the word from the cursor position to the start of the next word"},
		Cheat{"y$", "yank (copy) to end of line"},
		Cheat{"p", "put (paste) the clipboard after cursor"},
		Cheat{"P", "put (paste) before cursor"},
		Cheat{"dd", "delete (cut) a line"},
		Cheat{"2dd", "delete (cut) 2 lines"},
		Cheat{"dw", "delete (cut) the characters of the word from the cursor position to the start of the next word"},
		Cheat{"D", "delete (cut) to the end of the line"},
		Cheat{"d$", "delete (cut) to the end of the line"},
		Cheat{"x", "delete (cut) character"},
	},
}

var exiting = CheatGroup{
	Name: "Exiting",
	Cheats: []Cheat{
		Cheat{":w", "write (save) the file, but don't exit"},
		Cheat{":w !sudo tee %", "write out the current file using sudo"},
		Cheat{":wq or :x or ZZ", "write (save) and quit"},
		Cheat{":q", "quit (fails if there are unsaved changes)"},
		Cheat{":q! or ZQ", "quit and throw away unsaved changes"},
		Cheat{":wqa", "write (save) and quit on all tabs"},
	},
}

var searchAndReplace = CheatGroup{
	Name: "Search and replace",
	Cheats: []Cheat{
		Cheat{"/pattern", "search for pattern"},
		Cheat{"?pattern", "search backward for pattern"},
		Cheat{"\\vpattern", "'very magic' pattern: non-alphanumeric characters are interpreted as special regex symbols (no escaping needed)"},
		Cheat{"n", "repeat search in same direction"},
		Cheat{"N", "repeat search in opposite direction"},
		Cheat{":%s/old/new/g", "replace all old with new throughout file"},
		Cheat{":%s/old/new/gc", "replace all old with new throughout file with confirmations"},
		Cheat{":noh", "remove highlighting of search matches"},
	},
}

var searchInMultipleFiles = CheatGroup{
	Name: "Search in multiple files",
	Cheats: []Cheat{
		Cheat{":vimgrep /pattern/ {file}", "search for pattern in multiple files (e.g. :vimgrep /foo/ **/*"},
		Cheat{":cn", "jump to the next match"},
		Cheat{":cp", "jump to the previous match"},
		Cheat{":copen", "open a window containing the list of matches"},
	},
}
