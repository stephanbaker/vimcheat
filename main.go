package main

import "fmt"

func main() {
	groups := []CheatGroup{
		global,
		cursorMovement,
		insertMode,
		workingWithMultipleFiles,
		editing,
		markingText,
		visualCommands,
		registers,
		marks,
		macros,
		cutAndPaste,
		exiting,
		searchAndReplace,
		searchInMultipleFiles,
	}

	for _, g := range groups {
		for _, c := range g.Cheats {
			fmt.Println(c)
		}
	}
}
