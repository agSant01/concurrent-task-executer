package task

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	start := time.Now()

	Task()

	end := time.Now()

	fmt.Println(end.Sub(start).Seconds())

	if end.Sub(start).Seconds() < 2 {
		t.Errorf("Not waited enough time.")
	}
}

func TestGetSimulatedTasks(t *testing.T) {
	cases := []int{2, 3, 1}

	for simulation, amount := range cases {
		var taskList []func() (string, error) = GetSimulatedTasks(amount)

		totalStart := time.Now()
		for idx, task := range taskList {

			simulationStart := time.Now()

			task()

			fmt.Printf("Simulation: %v | Task %v | Took %v\n", simulation, idx, time.Now().Sub(simulationStart))
		}
		simulationTime := time.Now().Sub(totalStart)
		fmt.Printf("Simulation: %v | Took %v\n", simulation, simulationTime)

		if simulationTime.Seconds() < float64(amount*2) {
			t.Errorf("Error")
		}
	}
}
