package main

import (
	"youtube-manager-go/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//Routes
	routes.Init(e)

	//Start server
	e.Logger.Fatal(e.Start(":8080"))
}

//ディレクトリ内のファイル呼び出し方法
