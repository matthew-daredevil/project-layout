package error_sample

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gopkg.in/go-playground/validator.v9"
)

func main() {
    binding.Validator = new(defaultValidator)

    r := gin.Default()

    r.GET("/codingninjas", func(c *gin.Context) {
        var query struct {
            Ninja string `form:"ninja" json:"ninja" binding:"required"`
        }

        if err := c.ShouldBind(&query); err != nil {
            for _, fieldErr := range err.(validator.ValidationErrors) {
                c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
                return // exit on the first error
            }
        }

        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.Run()
}