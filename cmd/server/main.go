package main

import (
	"fmt"
	"log"

	db "github.com/98y7tbnb97t/tasks-service/internal/database"
	"github.com/98y7tbnb97t/tasks-service/internal/task"
	grpc "github.com/98y7tbnb97t/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Инициализация БД
	db.InitDB()

	// 2. Репозиторий и сервис задач
	repo := task.NewRepository(db.DB)
	svc := task.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := grpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	fmt.Println("сервер запустился на порту: localhost:50052")

	// 4. Запуск gRPC Tasks-сервиса
	if err := grpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
