package env

import (
	"fmt"
	"io/ioutil"
	"os"
)

// GetVar will check for the given value in the environment
// If not found in environment, it will check [envvar]_FILE and if present read its contents as the value.
// Else, it will return the default value.
func GetVar(envvar string, def string) string {
	v := os.Getenv(envvar)
	if v == "" {
		// Let's try to look in the file
		fvar := envvar + "_FILE"
		fv := os.Getenv(fvar)
		if fv != "" {
			// Try to read the value from the given file
			if buf, err := ioutil.ReadFile(fv); err == nil {
				v = string(buf)
			} else {
				fmt.Printf("Failed to read file: [%s]\n", fv)
			}
		}
	}

	if v == "" {
		fmt.Printf("Missing env variable: [%s], using default value: [%s]\n", envvar, def)
		v = def
	}

	return v
}
