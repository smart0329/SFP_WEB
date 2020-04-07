// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"fmt"
	"os"
)

func CommandPrintln(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func CommandPrintErrorln(a ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stderr, a...)
}

func CommandPrettyPrintln(a ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stdout, a...)
}
