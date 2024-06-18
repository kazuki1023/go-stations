package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	// HealthzHandler の登録
	healthzHandler := handler.NewHealthzHandler()
	mux.Handle("/healthz", healthzHandler)

	//  TODOHandlerの作成
	todoService := service.NewTODOService(todoDB)
	todoHandler := handler.NewTODOHandler(todoService)

	// エンドポイントの登録
	mux.Handle("/todos", todoHandler)

	return mux
}
