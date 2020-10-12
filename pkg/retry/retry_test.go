package retry

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/agsant01/concurrent-task-executer/pkg/counter"

	"github.com/agsant01/concurrent-task-executer/internal/data"
	"github.com/google/uuid"
)

func testTask() (string, error) {
	id := uuid.New().String()
	rand.Seed(int64(time.Now().Nanosecond()))
	fail := rand.Intn(2)
	sleep := rand.Intn(2) + 1
	fmt.Println("[INFO testTask] Entering Task:", id, "Fail Mode:", fail, "Sleep for: ", sleep)
	time.Sleep(time.Duration(sleep) * time.Second)
	if fail == 1 {
		return id, fmt.Errorf("[ERROR testTask] Task %v Failed", id)
	}
	fmt.Println("[INFO testTask] Ending Task:", id)
	return id, nil
}

func TestRunTasksInThread(t *testing.T) {
	fmt.Println("Starting Test...")
	retry := 2

	channelToTest := make(chan data.Result)

	tasks := []func() (string, error){testTask, testTask, testTask}

	flags := counter.New(1)

	go runTasksInThread(channelToTest, tasks, retry, 0, &flags)

	go verifyThreadsCompletion(channelToTest, &flags)

	for {
		val, ok := <-channelToTest

		if ok == false {
			break
		}

		fmt.Printf("Printing task %v\n", val)
	}
}

func TestConcurrentRetryEven(t *testing.T) {
	tasks := []func() (string, error){testTask, testTask, testTask, testTask, testTask, testTask}
	retry := 4

	channelTest := ConcurrentRetry(tasks, 3, retry)

	var resultIds []string
	for {
		val, ok := <-channelTest

		if ok == false {
			fmt.Println("[DEBUG TestConcurrentRetryEven] No more results awaiting")
			break
		}

		resultIds = append(resultIds, val.ID)

		fmt.Printf("[DEBUG TestConcurrentRetryEven] Printing task %v\n", val)
	}

	fmt.Printf("[DEBUG TestConcurrentRetryEven] Successes: %v IDs: %v\n", len(resultIds), resultIds)
}

func TestConcurrentRetryOdd(t *testing.T) {
	tasks := []func() (string, error){testTask, testTask, testTask, testTask, testTask, testTask, testTask}
	retry := 4

	channelTest := ConcurrentRetry(tasks, 3, retry)

	var resultIds []string
	for {
		val, ok := <-channelTest

		if ok == false {
			fmt.Println("[DEBUG TestConcurrentRetryOdd] No more results awaiting")
			break
		}

		resultIds = append(resultIds, val.ID)

		fmt.Printf("[DEBUG TestConcurrentRetryOdd] Printing task %v\n", val)
	}

	fmt.Printf("[DEBUG TestConcurrentRetryOdd] Successes: %v IDs: %v\n", len(resultIds), resultIds)
}

func TestMultipleConcurrentRetryOdd(t *testing.T) {
	task1 := []func() (string, error){testTask, testTask}
	task2 := []func() (string, error){testTask, testTask, testTask}
	retry := 4

	ch1 := ConcurrentRetry(task1, 2, retry)
	ch2 := ConcurrentRetry(task2, 1, retry)

	var resultIds1 []string
	var resultIds2 []string

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		if ok1 != false {
			resultIds1 = append(resultIds1, val1.ID)
			fmt.Printf("[DEBUG TestMultipleConcurrentRetryOdd] CH1 Printing task %v\n", val1)
		}

		if ok2 != false {
			resultIds2 = append(resultIds2, val2.ID)
			fmt.Printf("[DEBUG TestMultipleConcurrentRetryOdd] CH2 Printing task %v\n", val2)
		}

		if ok1 == false && ok2 == false {
			fmt.Println("[DEBUG TestMultipleConcurrentRetryOdd] Both Channels are closed")
			break
		}

	}

	fmt.Printf("[DEBUG TestMultipleConcurrentRetryOdd] Successes1: %v IDs: %v\n", len(resultIds1), resultIds1)
	fmt.Printf("[DEBUG TestMultipleConcurrentRetryOdd] Successes2: %v IDs: %v\n", len(resultIds2), resultIds2)
}
