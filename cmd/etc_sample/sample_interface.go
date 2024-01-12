
package etc_sample

import "fmt"

type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
    fmt.Println("Log to console:", message)
}

func processWithLogging(logger Logger) {
    logger.Log("Processing data")
}

type Shape interface {
    Area() float64
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, shape := range shapes {
        area += shape.Area()
    }
    return area
}

type Walker interface {
    Walk()
}

type Dog struct{}

func (d Dog) Walk() {
    fmt.Println("Dog is walking")
}

type Cat struct{}

func (c Cat) Walk() {
    fmt.Println("Cat is walking")
}

func Sample_interface() {
    logger := ConsoleLogger{}
    processWithLogging(logger)

    square := Square{Side: 2}
    circle := Circle{Radius: 3}
    fmt.Println("Total Area:", totalArea(square, circle))

    dog := Dog{}
    cat := Cat{}
    dog.Walk()
    cat.Walk()
}
