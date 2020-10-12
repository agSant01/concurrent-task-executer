package retry

import (
	"fmt"

	"github.com/agsant01/concurrent-task-executer/internal/data"
	"github.com/agsant01/concurrent-task-executer/pkg/counter"
)

// ConcurrentRetry is...
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan data.Result {
	channelToReturn := make(chan data.Result)

	go func(channel chan<- data.Result, tasks []func() (string, error), concurrent int, retry int) {
		threadsTasks := make([][]func() (string, error), concurrent)

		threadFinishedFlag := counter.New(len(threadsTasks))

		for i := 0; i < len(tasks); i++ {
			threadsTasks[i%concurrent] = append(threadsTasks[i%concurrent], tasks[i])
		}

		for threadID, tasksForThread := range threadsTasks {
			go runTasksInThread(channel, tasksForThread, retry, threadID, &threadFinishedFlag)
		}

		go verifyThreadsCompletion(channel, &threadFinishedFlag)

	}(channelToReturn, tasks, concurrent, retry)

	return channelToReturn
}

func verifyThreadsCompletion(channel chan<- data.Result, flags *counter.ThreadCounter) {
	for {
		if flags.AreCompleted() {
			close(channel)
			break
		}
	}
}

func runTasksInThread(channel chan<- data.Result, concurrentTasks []func() (string, error), retry int, threadID int, threadFlag *counter.ThreadCounter) {
	fmt.Println("[INFO runTasksInThread] Entering Thread:", threadID, "Tasks:", len(concurrentTasks))

	for idx, taskFunc := range concurrentTasks {
		fmt.Printf("[INFO runTasksInThread] Thread: %v Task:%v\n", threadID, idx)
		retries := 0
		var resErr error
		for retries < retry && resErr == nil {
			result, resErr := taskFunc()
			if resErr == nil {
				fmt.Printf("[DEBUG runTasksInThread] Thread: %v Task: %v Sending back throught channel %v...\n", threadID, idx, channel)
				channel <- data.Result{Index: idx, Result: result, ID: result}
				break
			} else {
				retries++
				fmt.Printf("[TASK-ERROR runTasksInThread] Thread: %v Task: %v Task Error Retry %v of %v\n", threadID, idx, retries, retry)
			}
		}
	}

	threadFlag.Completed()
}
