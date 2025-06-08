package grpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/98y7tbnb97t/GoMicro/proto/taskpb"
	"github.com/98y7tbnb97t/GoMicro/proto/userpb"
	"github.com/98y7tbnb97t/tasks-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	userId := req.GetUserId()

	_, err := h.userClient.GetUser(ctx, &userpb.UserRequest{Id: userId})
	if err != nil {
		return nil, fmt.Errorf("user %d not found: %w", userId, err)
	}

	t := &task.Task{
		Task:   req.GetTitle(),
		UserID: uint(userId),
		IsDone: false, // Новая задача по умолчанию не выполнена
	}
	if err := h.svc.CreateTask(t); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}
	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: uint32(t.UserID),
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.TaskRequest) (*taskpb.TaskResponse, error) {
	t := &task.Task{}
	if err := h.svc.GetTaskByID(strconv.Itoa(int(req.GetId())), t); err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}
	return &taskpb.TaskResponse{
		Id:     uint32(t.ID),
		Title:  t.Task,
		IsDone: t.IsDone,
		UserId: uint32(t.UserID),
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: uint32(t.UserID),
		})
	}
	return &taskpb.ListTasksResponse{
		Tasks: pbTasks,
	}, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	userId := req.GetUserId()
	tasks, err := h.svc.GetTasksForUser(uint(userId))
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks for user: %w", err)
	}
	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: uint32(t.UserID),
		})
	}
	return &taskpb.ListTasksByUserResponse{
		Tasks: pbTasks,
	}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	t := &task.Task{}
	if err := h.svc.GetTaskByID(strconv.Itoa(int(req.GetId())), t); err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}
	t.Task = req.GetTitle()
	t.IsDone = req.GetIsDone()
	if err := h.svc.UpdateTask(strconv.Itoa(int(req.GetId())), t); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			IsDone: t.IsDone,
			UserId: uint32(t.UserID),
		},
	}, nil
}
func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(strconv.Itoa(int(req.GetId()))); err != nil {
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}
	return &taskpb.DeleteTaskResponse{}, nil
}
