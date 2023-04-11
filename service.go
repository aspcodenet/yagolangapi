package main

func CalculateSalary(name string, age int, hourlySalary int, hoursWorked int) int {
	if age >= 50 {
		hourlySalary = int(float32(hourlySalary) * 1.1)
	}
	salary := hourlySalary * hoursWorked
	if name == "Stefan" {
		salary += 1000
	}
	return salary
}
