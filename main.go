package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"systementor.se/yagolangapi/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
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

func textToImage(c *gin.Context) {
	textContent := strings.Trim(c.Query("text"), " ")
	fontName := strings.Trim(c.Query("fontName"), " ")

	fontSize, err := strconv.ParseFloat(strings.Trim(c.Query("fontSize"), " "), 64)
	if err != nil {
		fontSize = 32
	}

	fontColorHex := strings.Trim(c.Query("fontColor"), " ")
	fontColorHex = strings.ToLower(fontColorHex)

	backgroundColorHex := strings.Trim(c.Param("backgroundColor"), " ")
	backgroundColorHex = strings.ToLower(backgroundColorHex)

	b, err := generateImage(textContent, fontColorHex, backgroundColorHex, fontName, fontSize)
	if err != nil {
		log.Println(err)
		return
	}

	remoteFileService := NewRemoteFile(config.S3.AwsKey,
		config.S3.AwsSecret,
		config.S3.HostBase, config.S3.Urlgenerator)

	uniqueFileName := uuid.New().String() + "-n1.png"
	url := remoteFileService.Upload(config.S3.BucketName,
		uniqueFileName,
		"image/png", b)
	//os.WriteFile("images/n1.png", b, 0777)

	c.String(200, url)
}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
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
	router.GET("/api/employees", employeesJson)
	router.GET("/api/addemployee", addEmployee)
	router.GET("/api/addmanyemployees", addManyEmployees)
	router.GET("/api/texttoimage", textToImage)
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
