package usecase

import "crud-go/internal/domain"

type EmployeeRepository interface {
	GetAll() ([]domain.Employee, error)
	GetByID(id int) (*domain.Employee, error)
	Create(employee *domain.Employee) error
	Update(employee *domain.Employee) error
	Delete(id int) error
}

type EmployeeUseCase struct {
	Repo EmployeeRepository
}

func NewEmployeeUseCase(repo EmployeeRepository) *EmployeeUseCase {
	return &EmployeeUseCase{Repo: repo}
}

func (uc *EmployeeUseCase) GetEmployees() ([]domain.Employee, error) {
	return uc.Repo.GetAll()
}

func (uc *EmployeeUseCase) GetEmployeeByID(id int) (*domain.Employee, error) {
	return uc.Repo.GetByID(id)
}

func (uc *EmployeeUseCase) CreateEmployee(employee *domain.Employee) error {
	return uc.Repo.Create(employee)
}

func (uc *EmployeeUseCase) UpdateEmployee(employee *domain.Employee) error {
	return uc.Repo.Update(employee)
}

func (uc *EmployeeUseCase) DeleteEmployee(id int) error {
	return uc.Repo.Delete(id)
}
