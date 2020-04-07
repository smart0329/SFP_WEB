// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package commands

import (
	"testing"
)

func TestVersion(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	th.CheckCommand(t, "version")
}