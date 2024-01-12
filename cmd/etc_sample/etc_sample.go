package etc_sample

import (
    "fmt"
    // "io"
    "log"
    "os"
    "sync"
    "time"
)

// 에러 처리
func ReadFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err
    }
    return data, nil
}

// 교착 상태를 방지하려면 채널에서 데이터를 보내고 받는 데 별도의 고루틴을 사용
func GoroutineChannelUsage() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
    fmt.Println("GoroutineChannelUsage:", <-ch)
}

// Speaker interface and Person struct demonstrate correct interface usage
type Speaker interface {
    Speak() string
}

type Person struct{}

func (p Person) Speak() string {
    return "Hello"
}

// ProcessFile demonstrates using defer for resource management
func ProcessFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Perform some operations with file
    fmt.Println("File processed:", filename)
}

// Increment demonstrates using concurrency primitives
var counter int
var mu sync.Mutex

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

func main() {
    // Error handling example
    _, err := ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
    }

    // Goroutine and channel usage example
    GoroutineChannelUsage()

    // Interface usage example
    bob := Person{}
    fmt.Println("Interface example:", bob.Speak())

    // Defer for resource management example
    ProcessFile("test.txt")

    // Concurrency primitives example
    for i := 0; i < 1000; i++ {
        go increment()
    }
    time.Sleep(time.Second) // Wait for goroutines to finish
    fmt.Println("Final counter value:", counter)
}