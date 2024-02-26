package main

import (
	"errors"
	"net/http"
	"strconv"

	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"systementor.se/yagolangapi/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
	//added a comment
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "Hej Golang"})
}

// HTML
// JSON

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func addEmployee(c *gin.Context) {

	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})

}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}

}

func apiEmployee(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)
}

func apiEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeUpdateById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		if err := c.BindJSON(&employee); err != nil {
			return
		}
		employee.Id, _ = strconv.Atoi(id)
		data.DB.Save(&employee)
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeDeleteById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	err := data.DB.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		data.DB.Delete(&employee)
		c.IndentedJSON(http.StatusNoContent, employee)
	}
}
func apiEmployeeAdd(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	err := data.DB.Create(&employee).Error
	if err != nil {

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, employee)
	}
}

var config Config

func main() {
	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	readConfig(&config)

	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/start", start)
	router.GET("/api/employee", apiEmployee)
	router.GET("/api/employee/:id", apiEmployeeById)
	router.PUT("/api/employee/:id", apiEmployeeUpdateById)
	router.DELETE("/api/employee/:id", apiEmployeeDeleteById)
	router.POST("/api/employee", apiEmployeeAdd)

	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.Run(":8080")

	// e := data.Employee{
	// 	Age:  1,
	// 	City: "Strefabn",
	// 	Namn: "wddsa",
	// }

	// if e.IsCool() {
	// 	fmt.Printf("Namn is cool:%s\n", e.Namn)
	// } else {
	// 	fmt.Printf("Namn:%s\n", e.Namn)
	// }

	// fmt.Println("Hello")
	// t := tabby.New()
	// t.AddHeader("Namn", "Age", "City")
	// t.AddLine("Stefan", "50", "Stockholm")
	// t.AddLine("Oliver", "14", "Stockholm")
	// t.Print()
}
