# **🚀 Go Task Processing Workflow**

This Go program demonstrates a simple task processing workflow using goroutines and channels. It initializes a task queue, creates multiple worker goroutines to process tasks concurrently, sends a specified number of tasks to the workers, and waits for all tasks to be completed before printing a confirmation message. The use of a `WaitGroup` ensures that the main program waits for all workers to finish processing before exiting.

---

## **✨ Features**

- ⚡ Utilizes **goroutines** for concurrent task processing.
- 📌 Implements **channels** for efficient task distribution.
- 🔄 Uses **sync.WaitGroup** to synchronize worker completion.
- ⏳ Simulates real-world task processing with **time.Sleep**.

---

## **📜 Code Overview**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, tasksChannel chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for taskId := range tasksChannel {
		executeTask(workerId, taskId)
	}
}

func executeTask(workerId int, taskId int) {
	fmt.Printf("Worker %d processing task %d\n", workerId, taskId)
	time.Sleep(time.Second)
}

func main() {
	const totalWorkers = 3
	const totalRequestsAllowed = 10
	tasksChannel := make(chan int, totalRequestsAllowed)
	var waitGroup sync.WaitGroup

	for workerIndex := 1; workerIndex <= totalWorkers; workerIndex++ {
		waitGroup.Add(1)
		go worker(workerIndex, tasksChannel, &waitGroup)
	}

	for taskIndex := 1; taskIndex <= totalRequestsAllowed; taskIndex++ {
		tasksChannel <- taskIndex
	}

	close(tasksChannel)
	waitGroup.Wait()

	fmt.Println("All tasks processed.")
}
```

---

## **🔍 How It Works**

1. **⚙️ Goroutines**: The main program creates a pool of workers (goroutines), each of which processes tasks from the channel.
2. **📦 Task Distribution**: Tasks are distributed across the workers through the channel, processed in parallel.
3. **🛠️ Synchronization**: `sync.WaitGroup` ensures the program waits for all workers to finish before exiting.

---

## **🔄 Workflow Summary**

1. **📌 Initialize Task Queue**: Create a buffered channel to hold tasks.
2. **🚀 Create Workers**: Launch `totalWorkers` worker goroutines.
3. **📤 Send Tasks**: Send `totalRequestsAllowed` tasks to the channel.
4. **🔒 Close the Channel**: Indicate that no more tasks will be added.
5. **⚡ Workers Process Tasks**: Workers read from the channel and process tasks concurrently.
6. **⏳ Wait for Completion**: The program waits for all workers to finish.
7. **✅ Final Message**: A confirmation message is printed after all tasks are processed.

---

## **📌 Business and Impact Value**

- 🏢 **Business Efficiency**: Helps optimize task processing, reducing bottlenecks in concurrent workflows.
- 💡 **Scalability**: Can be extended to handle a higher number of tasks with minimal modifications.
- ⏳ **Time Optimization**: Reduces processing time by leveraging concurrent workers.
- 🚀 **Practical Use Cases**: Ideal for job queues, background processing, and distributed task execution.

---

## **📌 Usage**

### **🔧 Prerequisites**

Ensure you have Go installed on your machine. You can verify this with:

```sh
$ go version
```

### **▶️ Running the Program**

To execute the program, run:

```sh
$ go run main.go
```

---

## **📜 License**

This project is licensed under the MIT License.

