package main

import (
    "fmt"
    "errors"
    "sort"
)

// Определение структуры для сотрудника компании
type Employee struct {
    ID       int
    Name     string
    Position string
    Salary   float64
}

// Метод для вывода информации о сотруднике
func (e Employee) PrintInfo() {
    fmt.Printf("ID: %d, Имя: %s, Должность: %s, Зарплата: %.2f\n", e.ID, e.Name, e.Position, e.Salary)
}

// Функция для добавления сотрудника в список
func AddEmployee(employees []Employee, id int, name string, position string, salary float64) []Employee {
    newEmployee := Employee{ID: id, Name: name, Position: position, Salary: salary}
    employees = append(employees, newEmployee)
    return employees
}

// Функция для удаления сотрудника по ID
func RemoveEmployee(employees []Employee, id int) ([]Employee, error) {
    for i, emp := range employees {
        if emp.ID == id {
            employees = append(employees[:i], employees[i+1:]...)
            return employees, nil
        }
    }
    return employees, errors.New("сотрудник с таким ID не найден")
}

// Функция для поиска сотрудника по имени
func SearchEmployee(employees []Employee, name string) (*Employee, error) {
    for i, emp := range employees {
        if emp.Name == name {
            return &employees[i], nil
        }
    }
    return nil, errors.New("сотрудник с таким именем не найден")
}

// Функция для вывода всех сотрудников
func PrintAllEmployees(employees []Employee) {
    fmt.Println("\nСписок сотрудников:")
    for _, emp := range employees {
        emp.PrintInfo()
    }
}

// Функция для сортировки сотрудников по зарплате
func SortEmployeesBySalary(employees []Employee) {
    sort.Slice(employees, func(i, j int) bool {
        return employees[i].Salary > employees[j].Salary
    })
}

func main() {
    // Инициализация списка сотрудников
    var employees []Employee
    employees = AddEmployee(employees, 1, "Иван Иванович", "Инженер", 50000.0)
    employees = AddEmployee(employees, 2, "Мария Петровна", "Менеджер", 60000.0)
    employees = AddEmployee(employees, 3, "Сергей Никитич", "Разработчик", 70000.0)

    // Вывод всех сотрудников
    PrintAllEmployees(employees)

    // Сортировка сотрудников по зарплате
    fmt.Println("\nСортировка сотрудников по зарплате:")
    SortEmployeesBySalary(employees)
    PrintAllEmployees(employees)

    // Поиск сотрудника по имени
    fmt.Println("\nПоиск сотрудника по имени 'Мария Петровна':")
    employee, err := SearchEmployee(employees, "Мария Петровна")
    if err != nil {
        fmt.Println(err)
    } else {
        employee.PrintInfo()
    }

    // Удаление сотрудника по ID
    fmt.Println("\nУдаление сотрудника с ID 2:")
    employees, err = RemoveEmployee(employees, 2)
    if err != nil {
        fmt.Println(err)
    } else {
        PrintAllEmployees(employees)
    }

    // Работа с пользователем
    var choice int
    fmt.Println("\nМеню:")
    fmt.Println("1. Добавить сотрудника")
    fmt.Println("2. Удалить сотрудника")
    fmt.Println("3. Найти сотрудника")
    fmt.Println("4. Выйти")
    
    for {
        fmt.Print("Введите ваш выбор: ")
        fmt.Scanln(&choice)
        
        switch choice {
        case 1:
            var id int
            var name string
            var position string
            var salary float64
            fmt.Print("Введите ID: ")
            fmt.Scanln(&id)
            fmt.Print("Введите имя: ")
            fmt.Scanln(&name)
            fmt.Print("Введите должность: ")
            fmt.Scanln(&position)
            fmt.Print("Введите зарплату: ")
            fmt.Scanln(&salary)
            employees = AddEmployee(employees, id, name, position, salary)
            PrintAllEmployees(employees)
        case 2:
            var id int
            fmt.Print("Введите ID сотрудника для удаления: ")
            fmt.Scanln(&id)
            employees, err = RemoveEmployee(employees, id)
            if err != nil {
                fmt.Println(err)
            } else {
                PrintAllEmployees(employees)
            }
        case 3:
            var name string
            fmt.Print("Введите имя сотрудника для поиска: ")
            fmt.Scanln(&name)
            employee, err := SearchEmployee(employees, name)
            if err != nil {
                fmt.Println(err)
            } else {
                employee.PrintInfo()
            }
        case 4:
            fmt.Println("Программа завершена.")
            return
        default:
            fmt.Println("Неверный выбор. Пожалуйста, выберите значение от 1 до 4.")
        }
    }
}
