package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

type App struct {
	Router *mux.Route
	DB *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(addr string) {}