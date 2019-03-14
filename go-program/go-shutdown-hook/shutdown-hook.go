package go_shutdown_hook

import (
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals chan os.Signal
var funcChannel chan func()
var isStarted = false
var funcArray []func()
var done chan struct{}

func Add(f func()) {
	if !isStarted {
		shutdownSignals = make(chan os.Signal, 1)
		funcChannel = make(chan func())
		done = make(chan struct{})
		signal.Notify(shutdownSignals, syscall.SIGINT, syscall.SIGTERM)
		start()
		isStarted = true
	}
	funcChannel <- f
}

func Wait() {
	if isStarted {
		<-done
	}
}

func executeHooks() {
	for _, f := range funcArray {
		f()
	}
}

func start() {
	go func() {
		shutdown := false
		for !shutdown {
			select {
			case <-shutdownSignals:
				executeHooks()
				shutdown = true
			case f := <-funcChannel:
				funcArray = append(funcArray, f)
			}
		}
		done <- struct{}{}
	}()
}

/**
程序退出函数
*/
//func beforeExit(handler func(error)) (error, chan bool) {
//	done := make(chan bool)
//	c := make(chan os.Signal, 2)
//	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
//	go func() {
//		<-c
//		handler(nil)
//		done <- true
//	}()
//	return nil, done
//}
//func doingExitHandle(e error) {
//	fmt.Printf("应用程序退出\n%v", e.Error())
//}
