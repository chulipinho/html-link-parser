package errors

import (
	"fmt"
	"os"
	"strings"
)

func ThrowError(msg ...string) {
	defaultMessage := "Error:"
	message := strings.Join(msg, "")
	fmt.Println(defaultMessage, message)

	os.Exit(1)
}
