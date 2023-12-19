package repository

import "crud-go/internal/domain"

type EmployeeRepository struct {
	Storage *[]domain.Employee
}

func NewEmployeeRepository(storage *[]domain.Employee) *EmployeeRepository {
	return &EmployeeRepository{Storage: storage}
}

func (r *EmployeeRepository) GetAll() ([]domain.Employee, error) {
	return *r.Storage, nil
}

func (r *EmployeeRepository) GetByID(id int) (*domain.Employee, error) {

    for _, employee := range *r.Storage {
		if employee.ID == id {
			return &employee, nil
		}
	}
	return nil, nil
}

func (r *EmployeeRepository) Create(employee *domain.Employee) error {
	// Replace with proper ID generation logic in a real application
	employee.ID = len(*r.Storage) + 1
	*r.Storage = append(*r.Storage, *employee)
	return nil
}

func (r *EmployeeRepository) Update(employee *domain.Employee) error {
	for i, existingEmployee := range *r.Storage {
        if existingEmployee.ID == employee.ID {
			(*r.Storage)[i] = *employee
			return nil
		}
	}
	return nil
}

func (r *EmployeeRepository) Delete(id int) error {
	for i, employee := range *r.Storage {
		if employee.ID == id {
			*r.Storage = append((*r.Storage)[:i], (*r.Storage)[i+1:]...)
			return nil
		}
	}
	return nil
}