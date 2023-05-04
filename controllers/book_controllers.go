package controllers

import (
	"github.com/gin-gonic/gin"
	"go-mongo/controllers/requests/book"
	"go-mongo/services"
	"net/http"
)

func CreateBook(c *gin.Context) {

	request := &requests.CreateRequest{}

	//Parse request from Context and validate fields too
	if err := requests.Parse(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// Create Service
	err := services.CreateBookService(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book Created",
	})
}
func CreateManyBooks(c *gin.Context) {

	request := &[]requests.CreateRequest{}

	//Parse All request from Context and validate fields too

	if err := requests.Parse(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := services.CreateManyBookService(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Books Created",
	})
}
func UpdateBook(c *gin.Context) {
	request := &requests.UpdateRequest{}

	id := c.Param("id")

	//Parse request from Context and validate fields too
	if err := requests.Parse(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//Update Service
	err := services.UpdateBookService(request, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book Updated",
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	result, err := services.GetBookByIdService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book retrieved",
		"data":    *result,
	})
}

func ListAllBooks(c *gin.Context) {

	results, err := services.ListBookService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Books retrieved",
		"data": gin.H{
			"count": len(*results),
			"books": *results,
		},
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := services.DeleteBookService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book Deleted",
		"data":    *result,
	})
}

func DeleteMultipleBooks(c *gin.Context) {

	request := &requests.DeleteBookRequest{}
	//Parse request from Context and validate fields too
	if err := requests.Parse(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := services.DeleteMultipleBookService(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Books Deleted",
	})

}
func CreateUpdateBook(c *gin.Context) {

	request := &requests.CreateUpdateBookRequest{}

	//Parse request from Context and validate fields too
	if err := requests.Parse(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//Create or Update Book into the DB
	err := services.CreateUpdateBookService(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Upsert Successfully",
	})
}
