// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type ChannelMemberHistory struct {
	ChannelId string
	UserId    string
	JoinTime  int64
	LeaveTime *int64
}
