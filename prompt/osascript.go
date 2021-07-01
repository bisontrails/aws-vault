package prompt

import (
	"fmt"
	"os/exec"
	"strings"
)

func OSAScriptMfaPrompt(mfaSerial string) (string, error) {
	/* #nosec G204 */
	// command is static
	// user input is sanitized for osascript context using %q
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`
		display dialog %q default answer "" buttons {"OK", "Cancel"} default button 1
        text returned of the result
        return result`,
		mfaPromptMessage(mfaSerial)))

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func init() {
	Methods["osascript"] = OSAScriptMfaPrompt
}
