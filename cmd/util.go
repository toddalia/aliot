package cmd

import (
	"fmt"
	"os"
)

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
