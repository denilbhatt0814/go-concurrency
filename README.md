# Go Concurrency Exercises

## Overview
This repository serves as a practical collection of Go concurrency examples and exercises associated with the GoLang Dojo YouTube playlist. It is designed to provide hands-on experience with Go's powerful concurrency features.

## Contents
Each directory contains specific examples and exercises on various concurrency topics:

1. `01goroutines` - Introduction to basic goroutines.
2. `02channels` - Using channels for goroutine communication.
3. `03bufferedChannels` - Implementing buffered channels.
4. `04channelIterationsNClosing` - Iterating and closing channels.
5. `05channelSelect` - Select statement with channels.
6. `06waitGroup` - Synchronization with WaitGroups.
7. `07mutexRWMutex` - Mutual exclusion with Mutex and RWMutex.
8. `08once` - Using sync.Once for one-time initializations.
9. `09pool` - Pooling resources with sync.Pool.
10. `10signalBroadcasting` - Broadcasting signals to goroutines.
11. `11syncMap` - Working with concurrent maps using sync.Map.
12. `12atomic` - Atomic operations for lock-free concurrency.
13. `13speedrun` - Concurrency speedrun challenges.

## Getting Started
To run the examples, navigate to the desired directory and execute the Go file:

```sh
cd 01goroutines
go run main.go
```

## Prerequisites
Make sure you have Go installed on your machine. The code is tested with Go 1.19.
