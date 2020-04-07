// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package interfaces

import "github.com/mattermost/mattermost-server/v5/model"

type MigrationsJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
