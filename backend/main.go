package main

import (
	. "foresight-app.v1/backend/apps/libs"
	"foresight-app.v1/backend/apps/route"
)

func init() {
	LoadAutoConfig()
	LoadRedisClient()
}

func main() {

	// go routine call
	//---------------------------------------------------------
	//// scheduler
	//go routines.CreateSchedulerSvc().Start()
	//
	//// mqtt
	//go routines.CreateMqttSvc().Start()
	//---------------------------------------------------------

	router := route.Init()
	router.Logger.Fatal(router.Start(":1323"))
}