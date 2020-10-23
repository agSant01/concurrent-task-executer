# concurrent-task-executer
This is a concurrent task executer written in GoLang. For Programming Languages Course ICOM 4036.

## Overview

This assignment implements a module that can execute concurrent multiple tasks concurrently with a specified amount of retries for each of the tasks. The number of concurrencies can be specified in the module as well the the required retries.

## Usage

For simplicity this code uses a **Makefile** for running predefined tasks.

### Test
To test the proyect the navigate to de root directory of this module and run `make test`.

### Binary
To generate a binary for the module run `make build`.

### Run
For running the proyect you can execute the created target binary directly or use the convenience `make run` command that will execute the ouput if it lies in the `target/` directory.

## Docs

### ConcurrentRetry

#### Parameters

<br>

| Name | Type | Description |
|-|-|-|
| `tasks` | []func() (string, error) | An array of functions that receive no parameters and return a tuple of (`string`, `error`) |
| `concurrent` | int | Specifies the number of maximum threads allowed. If there are more threads than `tasks` then each task will be executed in its own thread. |
| `retry` | int | Specifies the maximum number of retries per tasks |

<br>

#### Description

A function that allocates tasks to a specified number of threads.

### ObserveChannel

#### Parameters

<br>ObserveChannel( <-chan models.Result)

| Name | Type | Description |
|-|-|-|
| `channel` | <-chan models.Result | The channel to monitor and receive messages from. |

<br>

#### Description

Meant for debugging purposes. This function will print to console all messages received from the channel. 




## Project Struture

### internal/

This folder contains internal logic and structs required by the package internally that do not require to be accessed by external modules. This are the internal counter required to keep track of which tasks are finished and a set of premade tasks made for testing purposes.

#### counter/

Keeps track of which concurrent tasks have been finished using a `mutex Int` structure.

#### task/

Simulated tasks for testing purposes. This tasks are generated with random sleep times and random fail rates.

### pkg/

Contains the concurrent retry related logic. Contains the `Result` struct returned by the channels created by `ConcurrentRetry` function as well as a helper function to monitor the messages sent throught the channel created by it. This helper function is `ObserveChannel`, which taks is to receive the channel messages and print the messages to console.

#### models/

`Result` structure used by the channel to send back messages.

#### retry/

Contains the `ObserveChannel` helper function and the `ConcurrentRetry` logic. This also hides the task of monitoring the tasks for completion to close the channel and the thread creation complexity away from the user.
