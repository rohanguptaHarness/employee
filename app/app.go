package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App - Core Application configuration
type App struct {
	DB     *sql.DB
	Router *mux.Router
}
