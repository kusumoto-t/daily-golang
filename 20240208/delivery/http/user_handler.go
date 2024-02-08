package http

import (
	"encoding/json"
	"net/http"

	"testClean/usecase"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// リクエストからIDを取得
	userID := r.URL.Path[len("/user/"):]

	// ユースケースを使ってユーザーを取得
	user, err := h.UserUsecase.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}