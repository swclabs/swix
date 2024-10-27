/**
 * app folder representing the delevery layer in clean architecture
 * you can use this folder to define any configuration settings or
 * operation, start-up applications

 * Package app implement api server for swipe application

 * You can use _Server to connect to specific service adapters.
 * use fx Framework (uber-go/fx) to create your own adapters
 * with dependency injection pattern.

 * See the example below.

Example:

package main

import (
	"swclabs/swipex/app"
	_ "swclabs/swipex/app/init"
	"swclabs/swipex/internal/apis"
)

func main() {
	app := app.Builder(apis.NewApp)
	_ = app.Run()
}
*/

package app

import (
	_ "swclabs/swipex/app/init" // init package deps, like docs, migration
)
