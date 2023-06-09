package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
    ID      string `json:"id"`
    Text    string `json:"text"`
    Date    string    `json:"day"`
    Reminder bool `json:"reminder"`
}


//TODO
//Remove when done with testing
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

	collection := c.Param("collection")
	var jsonData map[string]interface{}

	//TODO
	// will need error handling if collection does not exist in file

	file, err := ioutil.ReadFile("db.json")
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Please create a db.json file at the project route"})
	}

	err = json.Unmarshal(file, &jsonData)
	if err != nil{

		fmt.Print(err)
	}

	c.JSON(http.StatusOK, jsonData[collection])
}

func HandleUpdate(c *gin.Context) {
	var jsonData map[string]interface{}
	var itemToUpdate map[string]interface{}

	// Grab query params
	collection := c.Param("collection")
	id := c.Param("id")

	file, err := ioutil.ReadFile("db.json")
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Please create a db.json file at the project route"})
		return
	}

	err = json.Unmarshal(file, &jsonData)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing JSON data"})
		return
	}	

	collectionData, ok := jsonData[collection].([]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Collection not found. Please ensure that it is inside of db.json"})
		return
	}

	for _, field := range collectionData {
		fieldMap, ok := field.(map[string]interface{})
		if !ok {
			continue
		}

		if fieldItemID, idOk := fieldMap["id"].(float64); idOk && strconv.Itoa(int(fieldItemID)) == id {
			itemToUpdate = fieldMap
			break
		}
	}

	if itemToUpdate != nil {
		currentValue := itemToUpdate["reminder"].(bool)
		itemToUpdate["reminder"] = !currentValue

		updatedData, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ioutil.WriteFile("db.json", updatedData, os.ModePerm)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully written to file"})
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
