package task

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// GetSimulatedTasks is...
func GetSimulatedTasks(quantitySimulatedTTasks int) []func() (string, error) {
	var arrayOfTasks []func() (string, error)

	if quantitySimulatedTTasks < 1 {
		quantitySimulatedTTasks = rand.Intn(10) + 1
	}

	for i := 0; i < quantitySimulatedTTasks; i++ {
		arrayOfTasks = append(arrayOfTasks, Task)
	}

	return arrayOfTasks
}

// Task is a...
func Task() (string, error) {
	rand.Seed(int64(time.Now().Nanosecond()))
	secondsToSleep := rand.Intn(2) + 2

	var taskID = uuid.New().String()

	fmt.Printf("[INFO Task] Task Id: %v Will Sleep for %v seconds...\n", taskID, secondsToSleep)

	time.Sleep(time.Duration(secondsToSleep) * time.Second)

	return taskID, nil
}
