// Copyright (c) 2019-present Smart.Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

type MetricsInterface interface {
	StartServer()
	StopServer()

	IncrementPostCreate()
	IncrementWebhookPost()
	IncrementPostSentEmail()
	IncrementPostSentPush()
	IncrementPostBroadcast()
	IncrementPostFileAttachment(count int)

	IncrementHttpRequest()
	IncrementHttpError()
	ObserveHttpRequestDuration(elapsed float64)

	IncrementClusterRequest()
	ObserveClusterRequestDuration(elapsed float64)
	IncrementClusterEventType(eventType string)

	IncrementLogin()
	IncrementLoginFail()

	IncrementEtagHitCounter(route string)
	IncrementEtagMissCounter(route string)

	IncrementMemCacheHitCounter(cacheName string)
	IncrementMemCacheMissCounter(cacheName string)
	IncrementMemCacheInvalidationCounter(cacheName string)
	IncrementMemCacheMissCounterSession()
	IncrementMemCacheHitCounterSession()
	IncrementMemCacheInvalidationCounterSession()

	IncrementWebsocketEvent(eventType string)
	IncrementWebSocketBroadcast(eventType string)

	AddMemCacheHitCounter(cacheName string, amount float64)
	AddMemCacheMissCounter(cacheName string, amount float64)

	IncrementPostsSearchCounter()
	ObservePostsSearchDuration(elapsed float64)
	ObserveStoreMethodDuration(method string, success string, elapsed float64)
	ObserveApiEndpointDuration(endpoint string, elapsed float64)
}
