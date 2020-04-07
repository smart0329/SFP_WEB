// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type DataRetentionJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
