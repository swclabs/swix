/**
 * boot folder representing the delevery layer in clean architecture
 * you can use this folder to define any configuration settings or
 * operation, start-up applications

 * Package boot implement api server for swipe application

 * You can use _Server to connect to specific service adapters.
 * use fx Framework (uber-go/fx) to create your own adapters
 * with dependency injection pattern.

 * See the example below.

Example:

package main

import (
	"swclabs/swix/boot"
	_ "swclabs/swix/boot/init"
	"swclabs/swix/internal/apis"
)

func main() {
	app := boot.App(apis.NewAPIServer)
	_ = app.Run()
}
*/

package boot

import (
	_ "swclabs/swix/boot/init" // init package deps, like docs, migration
)
