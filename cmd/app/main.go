package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"crud-go/internal/handler"
	"crud-go/internal/domain"
	"crud-go/internal/usecases"
	"crud-go/internal/repository"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// In-memory storage (replace with a database in a real application)
	var employees []domain.Employee

	// Initialize UseCase and Handler
	employeeRepo := repository.NewEmployeeRepository(&employees)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeUseCase)

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/employees", employeeHandler.GetEmployees)
	r.Get("/employees/{id}", employeeHandler.GetEmployeeByID)
	r.Post("/employees", employeeHandler.CreateEmployee)
	r.Put("/employees/{id}", employeeHandler.UpdateEmployee)
	r.Delete("/employees/{id}", employeeHandler.DeleteEmployee)

	http.ListenAndServe(":8080", r)
}

// func NewEmployeeRepository(storage *[]domain.Employee) *EmployeeRepository {
// 	return &EmployeeRepository{Storage: storage}
// }
