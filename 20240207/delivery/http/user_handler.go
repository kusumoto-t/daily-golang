package http

import (
	"encoding/json"
	"net/http"

	"github.com/kusumoto-t/cleanarchitecture/domain"
	"github.com/kusumoto-t/cleanarchitecture/usecase"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// リクエストからIDを取得
	id := r.URL.Query().Get("id")

	// ユースケースを使ってユーザーを取得
	user, err := h.UserUsecase.GetUserByID(id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) SaverUser(w http.ResponseWriter, r *http.Request) {
	// リクエストからユーザー情報を取得
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ユーザーを保存
	if err := h.UserUsecase.SaveUser(&user); err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	w.WriteHeader(http.StatusCreated)
}