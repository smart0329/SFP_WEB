// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type MfaInterface interface {
	GenerateSecret(user *model.User) (string, []byte, *model.AppError)
	Activate(user *model.User, token string) *model.AppError
	Deactivate(userId string) *model.AppError
	ValidateToken(secret, token string) (bool, *model.AppError)
}