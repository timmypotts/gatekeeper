// Expose references to router and database that we will use
package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Adding Initialize method to connect to DB
func (a *App) Initialize(user, password, dbname string) {}

// Run method starts application'
func (a *App) Run(addr string) {}
