package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalaryIsCalculatedFromHoursTimesHourlysalary(t *testing.T) {
	//arrange
	name := ""
	age := 20
	hourlySalary := 100
	hoursWorked := 5
	//act
	result := CalculateSalary(name, age, hourlySalary, hoursWorked)

	assert.Equal(t, 5001, result)

}

func TestCalculates_correct_name_bonus(t *testing.T) {
	//arrange
	name := "Stefan"
	age := 20
	hourlySalary := 100
	hoursWorked := 5
	//act
	result := CalculateSalary(name, age, hourlySalary, hoursWorked)

	assert.Equal(t, 1500, result)

}

func TestCalculates_correct_age_bonus(t *testing.T) {
	//arrange
	name := ""
	age := 50
	hourlySalary := 100
	hoursWorked := 5
	//act
	result := CalculateSalary(name, age, hourlySalary, hoursWorked)

	assert.Equal(t, 550, result)
}
