package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Record struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Activity string `json:"activity"`
	Name     string `json:"name"`
	Injury   string `json:"injury"`
	Species  string `json:"species"`
	Id       uuid.UUID
}

type RecordRequest struct {
	Date     string `form:"date" binding:"required"`
	Country  string `form:"country" binding:"required"`
	Activity string `form:"activity" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Injury   string `form:"injury" binding:"required"`
	Species  string `form:"species" binding:"required"`
}
type UpdateRecordRequest struct {
	Date     string `form:"date"`
	Country  string `form:"country"`
	Activity string `form:"activity"`
	Name     string `form:"name"`
	Injury   string `form:"injury"`
	Species  string `form:"species"`
}

func setupRouter(data []Record) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./static/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Shark Atatcks",
		})
	})
	r.GET("/shark", func(c *gin.Context) {
		c.HTML(http.StatusOK, "get_shark_template.html", gin.H{"Data": data})
	})
	r.POST("/shark", func(c *gin.Context) {
		var requestBody RecordRequest
		err := c.Bind(&requestBody)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		data = append(data, Record{
			requestBody.Date,
			requestBody.Country,
			requestBody.Activity,
			requestBody.Name,
			requestBody.Injury,
			requestBody.Species,
			uuid.New(),
		})
		c.HTML(http.StatusOK, "get_shark_template.html", gin.H{"Data": data})
	})
	r.PUT("/shark/:id", func(c *gin.Context) {
		id := c.Param("id")
		var requestBody UpdateRecordRequest
		err := c.Bind(&requestBody)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		log.Println(requestBody.Name)
		for record := range data {
			if data[record].Id.String() == id {
				if requestBody.Name != "" {
					data[record].Name = requestBody.Name
				}
				if requestBody.Activity != "" {
					data[record].Activity = requestBody.Activity
				}
				if requestBody.Country != "" {
					data[record].Country = requestBody.Country
				}
				if requestBody.Injury != "" {
					data[record].Injury = requestBody.Injury
				}
				if requestBody.Date != "" {
					data[record].Date = requestBody.Date
				}
				if requestBody.Country != "" {
					data[record].Country = requestBody.Country
				}
				if requestBody.Species != "" {
					data[record].Species = requestBody.Species
				}
			}
		}
		c.HTML(http.StatusOK, "get_shark_template.html", gin.H{"Data": data})
	})
	r.DELETE("/shark/:id", func(c *gin.Context) {
		id := c.Param("id")

		for record := range data {
			if data[record].Id.String() == id {
				data = remove(data, record)
				c.HTML(http.StatusOK, "get_shark_template.html", gin.H{"Data": data})
				return
			}
		}
	})
	return r
}

func remove(slice []Record, s int) []Record {
	return append(slice[:s], slice[s+1:]...)
}

func importData() []Record {
	content, err := os.ReadFile("./global-shark-attack.json")
	if err != nil {
		log.Fatal(err)
	}
	var data []Record

	json_err := json.Unmarshal(content, &data)
	if json_err != nil {
		log.Fatal("Json err", err)
	}
	for record := range data {
		data[record].Id = uuid.New()
	}
	return data
}

func getRandomItems(arr []Record, numItems int) []Record {
	if numItems > len(arr) {
		return nil // Handle case where numItems is greater than array length
	}
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:numItems]
}

func main() {
	data := importData()
	data = getRandomItems(data, 10)

	r := setupRouter(data)
	r.Run(":8080")
}
