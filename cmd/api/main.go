package main

import "github.com/qori-aziz-kyc/wallet-backend/cmd/api/server"

func main() {
	app, err := server.NewApp()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
