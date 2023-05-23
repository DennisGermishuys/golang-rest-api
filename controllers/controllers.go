package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
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


func HandleDelete(c *gin.Context) {

	id := c.Param("id")

	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			c.JSON(http.StatusOK, Tasks)
		}
	}
}

func HandleGet(c *gin.Context) {

	//TODO
	// Read data from db.json file instead of the struct for api

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

func HandleUpdate(c *gin.Context) {

	id := c.Param("id")

	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Reminder = !Tasks[i].Reminder
			c.JSON(http.StatusOK, Tasks)
		}
	}
}

func HandleAdd(c *gin.Context) {

	var data Task
	structId := uuid.New().String()

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	data.ID = structId

	Tasks = append(Tasks, data)

	c.JSON(http.StatusOK, Tasks)
}
