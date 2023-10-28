package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"

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

	s := gocron.NewScheduler(time.UTC)

	leetcodeClient := leetcode.New(cfg.LeetCodeApiURL)
	rateLimiter := ratelimiter.New(5, 25)

	repo := repo.New(db)
	cr := service.New(repo, leetcodeClient)
	uc := usecase.New(cr)
	deliv := handler.New(uc, rateLimiter)
	r := deliv.InitRoutes()

	go runCronJobs(context.Background(), s, uc)

	log.Println("starting on port: ", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Printf("run server error: %v \n", err)
		return
	}
}

func runCronJobs(ctx context.Context, s *gocron.Scheduler, uc *usecase.UseCase) {
	s.Every(5).Minute().Do(func() {
		err := uc.UsrScoreUpdate(ctx)
		if err != nil {
			log.Println(err)
			return
		}
	})

	s.Every(1).Minute().Do(func() {
		err := uc.EventFinish(ctx)
		if err != nil {
			log.Println(err)
			return
		}
	})

	s.StartBlocking()
}
