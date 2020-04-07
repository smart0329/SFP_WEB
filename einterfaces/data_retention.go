// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type DataRetentionInterface interface {
	GetPolicy() (*model.DataRetentionPolicy, *model.AppError)
}
