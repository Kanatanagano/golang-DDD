package main

import (
	"context"

	"github.com/Kanatanagano/api/internal/presenter"
)

// @title ユーザー管理サービスAPI
// @version v0.1.0
// @description ユーザー管理サービスAPIの機能
// @host localhost:8080
func main(){
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}