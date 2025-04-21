package clipanic

import (
	"fmt"
	"github.com/swarupdonepudi/karayaml/cmd/karayaml/root"
	"os"
	"runtime"
	"runtime/debug"
)

// Handle displays an emergency error message to the user and a stack trace
// when a panic occurs. For non-panic internal errors, call reportError instead.
//
// finished should be set to false when the handler is deferred and set to true
// as the last statement in the scope. This trick is necessary to avoid catching
// and then discarding a panic(nil).
func Handle(finished *bool) {
	if panicPayload := recover(); !*finished {
		stack := string(debug.Stack())
		fullMsg := fmt.Sprintf("Panic: %s\n\nStack Trace:\n%s", panicPayload, stack)
		reportError(fullMsg)
		os.Exit(1)
	}
}

// reportError prints user-friendly instructions on how to report unexpected
// errors (including those from the backend that are not panics).
func reportError(msg string) {
	fmt.Fprintln(os.Stderr, "================================================================================")
	fmt.Fprintln(os.Stderr, "The KaraYaml CLI encountered an unexpected error.")
	fmt.Fprintln(os.Stderr, "We would appreciate a report: https://github.com/swarupdonepdu/karayaml/issues/")
	fmt.Fprintln(os.Stderr, "Please provide all of the below text in your report.")
	fmt.Fprintln(os.Stderr, "================================================================================")
	fmt.Fprintf(os.Stderr, "CLI Version:       %s\n", root.VersionLabel)
	fmt.Fprintf(os.Stderr, "Go Version:        %s\n", runtime.Version())
	fmt.Fprintf(os.Stderr, "Go Compiler:       %s\n", runtime.Compiler)
	fmt.Fprintf(os.Stderr, "Architecture:      %s\n", runtime.GOARCH)
	fmt.Fprintf(os.Stderr, "Operating System:  %s\n", runtime.GOOS)
	fmt.Fprintf(os.Stderr, "Error Details:     %s\n\n", msg)
}
