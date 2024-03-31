package config

import (
	"database/sql"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/redis/go-redis/v9"
	_ "go.uber.org/automaxprocs"
	"log"
	"sync"
	"template/adapter"
	"template/adapter/cache"
	"template/constant"
	"template/internal/api"
	"template/internal/repository/transaction"
	"template/internal/repository/user"
	transactionUcase "template/internal/usecase/transaction"
	userUcase "template/internal/usecase/user"
)

func InitService(cfg Config, server string) {
	var (
		wg  sync.WaitGroup
		cch cache.Cache
		db  *sql.DB
		err error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()

		switch cfg.Cache.Driver {
		case constant.CacheRedis:
			redisClient := redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
			})
			cch = cache.NewRedisCache(redisClient)
		case constant.CacheMemcache:
			client := memcache.New(fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port))
			cch = cache.NewMemcacheCache(client)
		}

	}()

	go func() {
		defer wg.Done()

		db, err = adapter.NewDatabase(cfg.DB)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

	// repository
	userRepo := user.NewUserRepository(db, cch)
	transactionRepo := transaction.NewTransactionRepository(db, cch)

	// usecase
	userUsecase := userUcase.NewUserUsecase(userRepo, transactionRepo)
	transactionUcase := transactionUcase.NewTransactionsUsecase(transactionRepo, userRepo)

	switch server {
	case constant.ServerGRPC:

	case constant.ServerRest:
		api.Run(cfg, userUsecase, transactionUcase)
	}

}
