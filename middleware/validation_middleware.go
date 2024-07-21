package middleware

import (
	"encoding/json"
	"net/http"
	"spycat/models"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ValidBreeds []string
	once        sync.Once
	validate    *validator.Validate
)

func init() {
	validate = validator.New()
}

func fetchValidBreeds() {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		panic("failed to fetch breeds from TheCatAPI")
	}
	defer resp.Body.Close()

	var breeds []struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		panic("failed to decode breeds response")
	}

	for _, breed := range breeds {
		ValidBreeds = append(ValidBreeds, breed.Name)
	}
}

func CatValidator() gin.HandlerFunc {
	once.Do(func() {
		fetchValidBreeds()
	})

	return func(c *gin.Context) {
		path := c.FullPath()
		if path != "/cats" && !strings.HasPrefix(path, "/cats/") {
			c.Next()
			return
		}

		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete {
			c.Next()
			return
		}

		var cat models.Cat

		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			c.Abort()
			return
		}

		if err := validate.Struct(cat); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationErrors {
				errors[e.Field()] = e.Error()
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			c.Abort()
			return
		}

		if c.Request.Method == http.MethodPut && cat.Breed == "" {
			c.Set("cat", &cat)
			c.Next()
			return
		}

		valid := false
		for _, b := range ValidBreeds {
			if b == cat.Breed {
				valid = true
				break
			}
		}
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breed"})
			c.Abort()
			return
		}

		c.Set("cat", &cat)
		c.Next()
	}
}
