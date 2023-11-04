package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"

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
	log := zerolog.New(os.Stdout)
	cfg, err := config.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("get config")
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("connect to postgres")
	}
	defer db.Close()

	s := gocron.NewScheduler(time.UTC)

	leetcodeClient := leetcode.New(cfg.LeetCodeApiURL)
	rateLimiter := ratelimiter.New(5, 25)

	repo := repo.New(db)
	cr := service.New(repo, leetcodeClient)
	uc := usecase.New(cr, log)
	deliv := handler.New(uc, rateLimiter, log)

	r := deliv.InitRoutes()

	go runCronJobs(context.Background(), log, s, uc)

	log.Info().Msgf("starting on port: %v", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, r); err != nil {
		log.Err(err).Msg("run server error")
		return
	}
}

func runCronJobs(ctx context.Context, log zerolog.Logger, s *gocron.Scheduler, uc *usecase.UseCase) {
	s.Every(5).Minute().Do(func() {
		err := uc.UsrScoreUpdate(ctx)
		if err != nil {
			log.Err(err).Send()
			return
		}
	})

	s.Every(1).Minute().Do(func() {
		err := uc.EventFinish(ctx)
		if err != nil {
			log.Err(err).Send()
			return
		}
	})

	s.StartBlocking()
}
