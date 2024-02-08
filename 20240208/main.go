package main

import (
	"log"
	"net/http"
	handler "testClean/delivery/http"
	"testClean/repository"
	"testClean/usecase"
)

// TODO: ディレクトリ構成
func main() {

	userRepo := repository.NewInMemoryUserRepository()

	userUsecase := usecase.UserUsecase{
		UserRepository: userRepo,
	}

	userHandler := handler.UserHandler{
		UserUsecase: userUsecase,
	}

	// ルーティング
	http.HandleFunc("/user/", userHandler.GetUserByID)

	// サーバ起動
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}
