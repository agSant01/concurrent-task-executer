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
	id := uuid.New().String()

	rand.Seed(int64(time.Now().Nanosecond()))

	fail := rand.Intn(2)

	sleep := rand.Intn(2) + 2

	fmt.Println("[INFO Task] Entering Task:", id, "| Fail Mode:", fail, "| Will Sleep for: ", sleep)

	time.Sleep(time.Duration(sleep) * time.Second)

	if fail == 1 {
		return id, fmt.Errorf("[ERROR Task] Task %v Failed", id)
	}

	fmt.Println("[INFO Task] Ending Task:", id)

	return id, nil
}
