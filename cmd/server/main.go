package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todov1 "github.com/shimabukuromeg/connect-go-todo-example/gen/proto/todo/v1" // generated by protoc-gen-go

	"github.com/shimabukuromeg/connect-go-todo-example/gen/proto/todo/v1/todov1connect"
)

type TodoServer struct {
	todos  sync.Map
	nextID int
}

func isEmpty(m *sync.Map) bool {
	isEmpty := true
	m.Range(func(k, v interface{}) bool {
		isEmpty = false
		return false // break after the first item
	})
	return isEmpty
}


func (s *TodoServer) CreateTask(
	ctx context.Context,
	req *connect.Request[todov1.CreateTaskRequest],
) (*connect.Response[todov1.CreateTaskResponse], error) {
	s.nextID++
	id := s.nextID

	newTodo := &todov1.TodoItem{
		Id:     uint64(id),
		Name:   req.Msg.Name,
		Status: req.Msg.Status,
	}

	// TODOを追加
	s.todos.Store(newTodo.Id, newTodo)
	log.Println("TODOを追加")

	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&todov1.CreateTaskResponse{
		Id:     newTodo.Id,
		Name:   newTodo.Name,
		Status: newTodo.Status,
	})

	// TODO一覧
	log.Println("TODO一覧")
	s.todos.Range(func(key, value interface{}) bool {
		log.Printf("key is %d, value is %v", key, value)
		return true
	})

	res.Header().Set("CreateTask-Version", "v1")
	return res, nil
}

func (s *TodoServer) UpdateTaskStatus(
	ctx context.Context,
	req *connect.Request[todov1.UpdateTaskStatusRequest],
) (*connect.Response[todov1.UpdateTaskStatusResponse], error) {
	id := req.Msg.Id
	log.Println("Request headers: ", req.Header())

	// 更新対象のtodoを取得
	t, ok := s.todos.Load(id)
	if !ok {
		fmt.Printf("id %d のtodoはありません", id)
	}

	todo := t.(*todov1.TodoItem)

	updateTodo := &todov1.TodoItem{
		Id:     id,
		Name:   todo.Name,
		Status: req.Msg.Status,
	}

	// TODOを更新
	s.todos.Store(id, updateTodo)
	log.Println("TODOを更新")

	res := connect.NewResponse(&todov1.UpdateTaskStatusResponse{
		Id:     updateTodo.Id,
		Status: todov1.Status_STATUS_DONE,
	})

	// TODO一覧
	log.Println("TODO一覧")
	if isEmpty(&s.todos) {
		log.Println("todoはありません")
	} else {
		s.todos.Range(func(key, value interface{}) bool {
			log.Printf("key is %d, value is %v", key, value)
			return true
		})
	}

	res.Header().Set("UpdateTaskStatus-Version", "v1")
	return res, nil
}

func (s *TodoServer) DeleteTask(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTaskRequest],
) (*connect.Response[todov1.DeleteTaskResponse], error) {
	id := req.Msg.Id
	log.Println("Request headers: ", req.Header())

	// 削除対象のtodoを取得
	_, ok := s.todos.Load(id)
	if !ok {
		fmt.Printf("id %d のtodoはありません", id)
	}

	// TODOを削除
	s.todos.Delete(id)
	log.Println("TODOを更新")

	res := connect.NewResponse(&todov1.DeleteTaskResponse{
		Id: id,
	})

	// TODO一覧
	log.Println("TODO一覧")
	s.todos.Range(func(key, value interface{}) bool {
		log.Printf("key is %d, value is %v", key, value)
		return true
	})

	res.Header().Set("DeleteTask-Version", "v1")
	return res, nil
}


func main() {
    todoServer := &TodoServer{}	
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"todo.v1.TodoService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	path, handler := todov1connect.NewToDoServiceHandler(todoServer)
	mux.Handle(path, handler)
	log.Fatal(http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))
}