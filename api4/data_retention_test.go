// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"testing"
)

func TestDataRetentionGetPolicy(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	_, resp := th.Client.GetDataRetentionPolicy()
	CheckNotImplementedStatus(t, resp)
}
