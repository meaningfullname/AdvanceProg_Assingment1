package main

import (
	"fmt"
	"strconv"
)

type FulltimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID           uint64
	Name         string
	Hourly_rated uint64
	Hours_Worked float32
}

type Comp struct {
	Employee map[uint64]Employee
}

var comp Comp

type Employee interface {
	GetDetails() string
}

func (fte FulltimeEmployee) GetDetails() string {
	return "Full-Time Employee: ID=" + strconv.FormatUint(fte.ID, 10) +
		", Name=" + fte.Name +
		", Salary=" + strconv.FormatUint(uint64(fte.Salary), 10)
}

func (pte PartTimeEmployee) GetDetails() string {
	totalEarnings := float32(pte.Hourly_rated) * pte.Hours_Worked
	return "Part-Time Employee: ID=" + strconv.FormatUint(pte.ID, 10) +
		", Name=" + pte.Name +
		", Hourly Rate=" + strconv.FormatUint(pte.Hourly_rated, 10) +
		", Hours Worked=" + fmt.Sprintf("%.1f", pte.Hours_Worked) +
		", Total Earnings=" + fmt.Sprintf("%.2f", totalEarnings)
}

func AddEmployeers(id uint64, emp Employee) {
	if comp.Employee == nil {
		comp.Employee = make(map[uint64]Employee)
	}
	comp.Employee[id] = emp
}

func ListEmployeers() {
	if len(comp.Employee) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for id, emp := range comp.Employee {
		fmt.Printf("Employee ID: %d\n", id)
		fmt.Println(emp.GetDetails())
		fmt.Println("---------------------------")
	}
}

func main() {
	comp.Employee = make(map[uint64]Employee)
	var i int
	var id uint64 = 1

	for {
		fmt.Println("Welcome! Input a number to choose functions:")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Show a list of employees")
		fmt.Println("3. Exit")
		fmt.Scan(&i)

		switch i {
		case 1:
			var ch int
			fmt.Println("Which employee do you want to add? Enter 1 for Part-Time, 2 for Full-Time:")
			fmt.Scan(&ch)

			if ch == 1 {
				var name string
				var workh float64
				var hourlyRate uint64

				fmt.Print("Enter Name, Hours Worked, Hourly Rate (space-separated): ")
				fmt.Scan(&name, &workh, &hourlyRate)

				emp := PartTimeEmployee{
					ID:           id,
					Name:         name,
					Hours_Worked: float32(workh),
					Hourly_rated: hourlyRate,
				}
				AddEmployeers(id, emp)
				id++
			} else if ch == 2 {
				var name string
				var salary uint64

				fmt.Print("Enter Name and Salary (space-separated): ")
				fmt.Scan(&name, &salary)

				emp := FulltimeEmployee{
					ID:     id,
					Name:   name,
					Salary: uint32(salary),
				}
				AddEmployeers(id, emp)
				id++
			} else {
				fmt.Println("Invalid choice, please try again.")
			}

		case 2:
			ListEmployeers()

		case 3:
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid input. Please enter a number between 1 and 3.")
		}
	}
}
