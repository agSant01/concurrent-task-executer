package retry

import (
	"fmt"

	"github.com/agsant01/concurrent-task-executer/internal/data"
)

// ObserveChannel keeps an receiving and printing messages until channel is Closed
func ObserveChannel(channel <-chan data.Result) {
	var resultIds []string
	for {
		val, ok := <-channel

		if ok != false {
			resultIds = append(resultIds, val.ID)
			fmt.Printf("[DEBUG ObserveChannel] Task Result %v\n", val)
		} else {
			fmt.Printf("[DEBUG ObserveChannel] Channel %v closed...\n", channel)
			break
		}

	}
	fmt.Printf("[DEBUG ObserveChannel] Successes: %v IDs: %v\n", len(resultIds), resultIds)
}
