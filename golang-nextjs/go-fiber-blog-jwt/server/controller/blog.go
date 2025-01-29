package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/database"
	"github.com/neerajbg/go-gin-auth/model"
)

// Blog list
func BlogList(c *gin.Context) {
	context := gin.H{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	// Sleep to add some delay in API response
	time.Sleep(time.Millisecond * 1500)

	db := database.DBConn
	var records []model.Blog
	db.Find(&records)

	context["blog_records"] = records

	c.JSON(200, context)
}

// Blog detail page
func BlogDetail(c *gin.Context) {
	id := c.Param("id")
	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		c.JSON(404, gin.H{"msg": "Record not Found."})
		return
	}

	c.JSON(200, gin.H{
		"statusText": "Ok",
		"msg":        "Blog Detail",
		"record":     record,
	})
}

// Add a Blog into Database
func BlogCreate(c *gin.Context) {
	response := gin.H{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	var record model.Blog

	// Handle multipart form
	if err := c.ShouldBind(&record); err != nil {
		log.Println("Error in parsing request:", err)
		response["statusText"] = ""
		response["msg"] = "Something went wrong"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// File upload handling
	file, err := c.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		log.Println("Error in file upload:", err)
	}

	if file != nil && file.Size > 0 {
		filename := "./static/uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			log.Println("Error in file uploading:", err)
			response["statusText"] = ""
			response["msg"] = "Error uploading file"
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		record.Image = filename
	}

	result := database.DBConn.Create(&record)
	if result.Error != nil {
		log.Println("Error in saving data:", result.Error)
		response["statusText"] = ""
		response["msg"] = "Something went wrong"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response["msg"] = "Record is saved successfully"
	response["data"] = record
	c.JSON(http.StatusCreated, response)
}

// Update a Blog
func BlogUpdate(c *gin.Context) {
	id := c.Param("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		c.JSON(400, gin.H{"msg": "Record not Found."})
		return
	}

	if err := c.ShouldBind(&record); err != nil {
		log.Println("Error in parsing request.")
		c.JSON(400, gin.H{"msg": "Something went wrong."})
		return
	}

	// File upload
	file, err := c.FormFile("file")
	if err == nil && file.Size > 0 {
		filename := "./static/uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, filename); err != nil {
			log.Println("Error in file uploading...", err)
		}
		record.Image = filename
	}

	result := database.DBConn.Save(record)
	if result.Error != nil {
		log.Println("Error in saving data.")
		c.JSON(400, gin.H{"msg": "Error in saving data."})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "Record updated successfully.",
		"data": record,
	})
}

// Delete a Blog
func BlogDelete(c *gin.Context) {
	id := c.Param("id")
	var record model.Blog

	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")
		c.JSON(404, gin.H{"msg": "Record not Found."})
		return
	}

	// Remove image
	if record.Image != "" {
		err := os.Remove(record.Image)
		if err != nil {
			log.Println("Error in deleting file.", err)
		}
	}

	result := database.DBConn.Delete(&record)
	if result.Error != nil {
		c.JSON(400, gin.H{"msg": "Something went wrong."})
		return
	}

	c.JSON(200, gin.H{"statusText": "Ok", "msg": "Record deleted successfully."})
}
