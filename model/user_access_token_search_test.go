// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestUserAccessTokenSearchJson(t *testing.T) {
	userAccessTokenSearch := UserAccessTokenSearch{Term: NewId()}
	json := userAccessTokenSearch.ToJson()
	ruserAccessTokenSearch := UserAccessTokenSearchFromJson(strings.NewReader(json))
	require.Equal(t, userAccessTokenSearch.Term, ruserAccessTokenSearch.Term, "Terms do not match")
}
