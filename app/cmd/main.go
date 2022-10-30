// Package main ...
package main

import (
	"github.com/Shmyaks/exchange-parser-server/app/core"
)

// @title          Exchange-parser-server
// @version        1.0
// @description    This is a docs for api
// @host           localhost:5050
// @BasePath       /api
func main() {
	core.NewServer().Start()
}
