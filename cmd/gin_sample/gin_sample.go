package gin_sample

import (
    "fmt"
    "log"
    "net/http"
    // "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    r := gin.Default()

    // Basic routing
    r.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Gin!")
    })

    // URL parameters
    r.GET("/hello/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello, %s!", name)
    })

    // Query parameters
    r.GET("/search", func(c *gin.Context) {
        query := c.DefaultQuery("q", "none")
        c.String(http.StatusOK, "Searched for: %s", query)
    })

    // JSON response
    r.GET("/profile", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "name": "John",
            "age":  25,
        })
    })

    // Grouping routes
    api := r.Group("/api")
    {
        api.GET("/tasks", getTasks)
        api.POST("/tasks", addTask)
    }

    // Middleware
    r.Use(RequestDuration())

    // File upload
    r.POST("/upload", func(c *gin.Context) {
        file, _ := c.FormFile("file")
        c.SaveUploadedFile(file, file.Filename)
        c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
    })

    // Custom error handling
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"message": "Not Found, sorry!"})
    })

    // Redirects
    r.GET("/old-route", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/new-route")
    })

    // Asynchronous handlers
    r.GET("/async", func(c *gin.Context) {
        copyContext := c.Copy()
        go func() {
            time.Sleep(5 * time.Second)
            log.Println("Done! in path " + copyContext.Request.URL.Path)
        }()
        c.String(http.StatusOK, "Async task started")
    })

    // CORS middleware
    r.Use(CORSMiddleware())
	return r
}

func getTasks(c *gin.Context) {
    // Logic to get tasks
}

func addTask(c *gin.Context) {
    // Logic to add a task
}

func RequestDuration() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)
        log.Println("Request took", duration)
    }
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Next()
    }
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}