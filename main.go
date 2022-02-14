package main

import (
	"github.com/PurotoApp/libpuroto/ginHelper"
	"github.com/PurotoApp/libpuroto/logHelper"
	"github.com/PurotoApp/libpuroto/mongoHelper"
	"github.com/PurotoApp/meltdown/internal/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	// create DB connection
	client, err := mongoHelper.ConnectDB(mongoHelper.GetDBUri())
	logHelper.ErrorFatal("MongoDB", err)
	// create collections

	// TODO: move into redis
	collSession := client.Database("authfox").Collection("session")
	// authfox needs to updated
	collProfiles := client.Database("meltdown").Collection("profiles")

	// test the connection
	logHelper.ErrorFatal("MongoDB", mongoHelper.TestDBConnection(client))
	// close connection on program exit
	// TODO: execute on CTRL+C
	defer func() {
		logHelper.ErrorFatal("MongoDB", mongoHelper.DisconnectDB(client))
	}()

	// create router
	router := gin.Default()

	// configure gin
	ginHelper.ConfigRouter(router)

	// set routes
	endpoints.SetRoutes(router, collSession, collProfiles)

	// start
	router.Run("localhost:3621")
}
