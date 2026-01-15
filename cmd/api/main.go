package main

import (
	"log"

	"github.com/husteve07/GoBlogAPI/internal/db"
	"github.com/husteve07/GoBlogAPI/internal/env"
	"github.com/husteve07/GoBlogAPI/internal/store"
	_ "github.com/lib/pq"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdletime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdletime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()

	log.Fatal(app.run(mux))
}
