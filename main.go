package main

import (
	"SnakeLadderGame/model"
	"SnakeLadderGame/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
