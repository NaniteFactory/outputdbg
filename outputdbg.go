package outputdbg

import (
	"fmt"
	"log"

	outputdebug "github.com/zetamatta/go-outputdebug"
)

// LogPrintln println()s its args to both your console & debug view.
func LogPrintln(args ...interface{}) {
	outputdebug.String(fmt.Sprintln(args...))
	log.Println(args...)
}
