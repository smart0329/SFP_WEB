// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type ElasticsearchIndexerInterface interface {
	MakeWorker() model.Worker
}

type ElasticsearchAggregatorInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
