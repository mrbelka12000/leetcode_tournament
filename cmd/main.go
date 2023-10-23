package main

import (
	"log"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/client/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/handler"
	"github.com/mrbelka12000/leetcode_tournament/internal/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/service"
	"github.com/mrbelka12000/leetcode_tournament/internal/usecase"
	"github.com/mrbelka12000/leetcode_tournament/pkg/config"
	"github.com/mrbelka12000/leetcode_tournament/pkg/database/postgres"
	"github.com/mrbelka12000/leetcode_tournament/pkg/ratelimiter"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("get config: %v", err)
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}
	defer db.Close()

	leetcodeClient := leetcode.New(cfg.LeetCodeApiURL)
	rateLimiter := ratelimiter.New(5, 25)

	repo := repo.New(db)
	cr := service.New(repo, leetcodeClient)
	uc := usecase.New(cr)
	deliv := handler.New(uc, rateLimiter)
	r := deliv.InitRoutes()

	log.Println("starting on port: ", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Printf("run server error: %v \n", err)
		return
	}
}
