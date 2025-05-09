package main

import (
	tech "TechShop"
	"TechShop/pkg/handler"
	"TechShop/pkg/redis"
	"TechShop/pkg/repository"
	"TechShop/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error initializing configs", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error initializing configs", err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	red := redis.InitRedis()
	if err != nil {
		logrus.Fatalf("failed to initializate a db: %s", err.Error())
	}
	repos := repository.NewRepository(db, red)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(tech.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
			logrus.Fatal("Error in cmd main - ", err)
		}
	}()
	logrus.Println("Todo App Shuting down")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("To do App Shutting DOWN")
	if err := srv.Shutdown(); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
