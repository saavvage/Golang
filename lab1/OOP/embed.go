package main
import "fmt"

type Employee struct {
    Name string
    ID   int
}

type Manager struct {
    Employee   
    Department string
}

func (e Employee) Work() {
    fmt.Printf("Employee Name: %s, Employee ID: %d is working.\n", e.Name, e.ID)
}

func main() {
    manager := Manager{
        Employee: Employee{
            Name: "Artem",
            ID:   22,
        },
        Department: "Engineering",
    }

    manager.Work()  
}
