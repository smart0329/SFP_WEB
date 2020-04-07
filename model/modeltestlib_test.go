// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckInt(t *testing.T, got int, expected int) {
	assert.Equal(t, got, expected)
}

func CheckInt64(t *testing.T, got int64, expected int64) {
	assert.Equal(t, got, expected)
}

func CheckString(t *testing.T, got string, expected string) {
	assert.Equal(t, got, expected)
}

func CheckTrue(t *testing.T, test bool) {
	assert.True(t, test)
}

func CheckFalse(t *testing.T, test bool) {
	assert.False(t, test)
}

func CheckBool(t *testing.T, got bool, expected bool) {
	assert.Equal(t, got, expected)
}
