// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/store/storetest"
)

func TestSchemeStore(t *testing.T) {
	StoreTest(t, storetest.TestSchemeStore)
}
