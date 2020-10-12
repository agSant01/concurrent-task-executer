package main

import (
	"github.com/agsant01/concurrent-task-executer/internal/task"
	"github.com/agsant01/concurrent-task-executer/pkg/retry"
)

func main() {

	simulatedTasks1 := task.GetSimulatedTasks(4)
	simulatedTasks2 := task.GetSimulatedTasks(7)

	channel1 := retry.ConcurrentRetry(simulatedTasks1, 2, 2)
	channel2 := retry.ConcurrentRetry(simulatedTasks2, 4, 5)

	retry.ObserveChannel(channel1)
	retry.ObserveChannel(channel2)

}
