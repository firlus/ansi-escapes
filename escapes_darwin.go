package escapes

import "os"

// ANSI escape sequences for saving and restoring the cursor position
var (
	CursorSavePosition    string
	CursorRestorePosition string
)

func init() {
	// Apple's terminal uses 7 and 8 rather than s and u
	if os.Getenv("TERM_PROGRAM") == "Apple_Terminal" {
		CursorSavePosition = Esc + "7"
		CursorRestorePosition = Esc + "8"
	} else {
		CursorSavePosition = Esc + "s"
		CursorRestorePosition = Esc + "u"
	}
}
