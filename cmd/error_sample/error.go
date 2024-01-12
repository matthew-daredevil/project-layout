package etc_sample

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

func Error_handle() {
    binding.Validator = new(defaultValidator)

    r := gin.Default()

    r.GET("/codingninjas", func(c *gin.Context) {
        var query struct {
            Ninja string `form:"ninja" json:"ninja" binding:"required"`
        }

        if err := c.ShouldBind(&query); err != nil {
            c.JSON(http.StatusBadRequest, err.Error())
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    r.Run()
}