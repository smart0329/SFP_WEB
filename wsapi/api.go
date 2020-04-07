// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package wsapi

import (
	"github.com/mattermost/mattermost-server/v5/app"
)

type API struct {
	App    *app.App
	Router *app.WebSocketRouter
}

func Init(a *app.App, router *app.WebSocketRouter) {
	api := &API{
		App:    a,
		Router: router,
	}

	api.InitUser()
	api.InitSystem()
	api.InitStatus()

	a.HubStart()
}
