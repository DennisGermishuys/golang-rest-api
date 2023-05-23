package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-post/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
    ID      string `json:"id"`
    Text    string `json:"text"`
    Date    string    `json:"day"`
    Reminder bool `json:"reminder"`
}

var Tasks = []Task{
		{ID: "1", Text: "Doctors Appointment", Date: "25 May 2023", Reminder: true},
		{ID: "2", Text: "Go to the gym", Date: "28 June 2023", Reminder: true},
		{ID: "3", Text: "Study", Date: "1 july 2023", Reminder: true},
	}


func handleDelete(c *gin.Context) {

	id := c.Param("id")

	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			c.JSON(http.StatusOK, Tasks)
		}
	}
}

func handleGet(c *gin.Context) {
	file, err := ioutil.ReadFile("db.json")
	if err != nil{

		fmt.Print(err)
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(file, &jsonData)
	if err != nil{

		fmt.Print(err)
	}

	fmt.Println(jsonData)
	c.JSON(http.StatusOK, Tasks)
}

func handleUpdate(c *gin.Context) {

	id := c.Param("id")

	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Reminder = !Tasks[i].Reminder
			c.JSON(http.StatusOK, Tasks)
		}
	}
}

func handleAdd(c *gin.Context) {
	var data Task
	structId := uuid.New().String()

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	data.ID = structId

	Tasks = append(Tasks, data)

	c.JSON(http.StatusOK, Tasks)
	fmt.Print(Tasks)


}

func main() {



	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	router.GET("/tasks", handleGet)
	router.PUT("/tasks/:id", handleUpdate)
	router.POST("/tasks", handleAdd)
	router.DELETE("/tasks/:id", handleDelete)

	router.Run(":8080")

}

