package main

import (
	"fmt"  //! For printing output.
	"sync" //! To use synchronization primitives, specifically WaitGroup for ensuring all goroutines complete before the program exits.
	"time" //! To simulate task processing time with Sleep.
)

// ! Worker function, This function simulates the task processing done by each worker.
// ! workerId: A unique identifier for the worker.
// ! tasksChannel: The channel from which the worker will fetch tasks.
// ! waitGroup: A pointer to the sync.WaitGroup, used to signal when the worker has completed all its tasks.
func worker(workerId int, tasksChannel chan int, waitGroup *sync.WaitGroup) {
	//! This ensures that wg.Done() is called when the worker function exits. It's used to signal that a goroutine has finished its work, and the WaitGroup can proceed.
	defer waitGroup.Done()
	//! This loop reads tasks from the tasks channel until it's closed. Each task is processed by the worker.
	for taskId := range tasksChannel {
		executeTask(workerId, taskId)
	}
}

func executeTask(workerId int, taskId int) {
	fmt.Printf("Worker %d processing task %d\n", workerId, taskId)
	//! simulates a task that takes 1 second to process.
	time.Sleep(time.Second)
}

func main() {

	//! Defines the number of workers in the pool (3 in this case).
	const totalWorkers = 3
	//! The total number of tasks to be processed (10 tasks in this case).
	const totalRequestsAllowed = 10
	//! Creates a buffered channel that can hold up to 10 tasksChannel. A buffered channel is useful here because it allows tasksChannel to be queued while the workers are still processing others.
	tasksChannel := make(chan int, totalRequestsAllowed) //! Buffered Channel
	//! Initializes a WaitGroup that will be used to wait for all workers to finish processing their tasks.
	var waitGroup sync.WaitGroup

	//! Start workers
	//! Loops over the number of workers (3 workers in this case).
	for workerIndex := 1; workerIndex <= totalWorkers; workerIndex++ {
		//! This increments the WaitGroup counter by 1, indicating that there is one more goroutine to wait for.
		waitGroup.Add(1)
		//! Launches a new goroutine that calls the worker function. The worker is assigned an ID (workerIndex), the task channel, and the WaitGroup.
		go worker(workerIndex, tasksChannel, &waitGroup)
	}

	//! Send tasks to the task queue
	//! Sends 10 tasks into the tasks channel.
	for taskIndex := 1; taskIndex <= totalRequestsAllowed; taskIndex++ {
		tasksChannel <- taskIndex
	}

	//! Closes the tasks channel once all tasks have been sent. This signals the workers that no more tasks are coming, and they should stop processing when the channel is empty.
	close(tasksChannel)

	//! This blocks the main program from exiting until all workers have completed their tasks. Each worker signals completion using wg.Done().
	waitGroup.Wait()

	//! Once all tasks are processed and all workers finish their work, the program prints a confirmation message.
	fmt.Println("All tasks processed.")
}

//? How It Works:-
//! Goroutines: The main program creates a pool of workers (goroutines), each of which processes tasks from the channel.
//! Task Distribution: Tasks are distributed across the workers through the channel, and the workers process them in parallel. Since the channel is buffered, the program can queue tasks even if all workers are busy, avoiding unnecessary blocking.
//! Synchronization: The sync.WaitGroup ensures that the main program waits for all workers to finish processing before exiting. This prevents the program from exiting prematurely.

//? Complete Workflow Summary:-
//! Initialize Task Queue: A buffered channel (tasks) is created to hold the tasks to be processed by workers.
//! Create Workers: Launch numWorkers worker goroutines. Each worker pulls tasks from the tasks channel and processes them.
//! Send Tasks: The main program sends numReq tasks to the tasks channel.
//! Close the Channel: The main program closes the tasks channel to indicate no more tasks will be added.
//! Workers Process Tasks: Workers process tasks concurrently, reading from the channel and simulating task processing.
//! Wait for Completion: The main program waits for all workers to finish processing using sync.WaitGroup.
//! Final Message: After all tasks are processed, the program prints a confirmation and exits.
