// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"context"

	"github.com/mattermost/mattermost-server/v5/model"
)

type MessageExportInterface interface {
	StartSynchronizeJob(ctx context.Context, exportFromTimestamp int64) (*model.Job, *model.AppError)
	RunExport(format string, since int64) *model.AppError
}
