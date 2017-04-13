package directoryIterator

import (
	"io/ioutil"
	"runtime"
	"sync"
)

func Start (path string, fp func(string)) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ch := make(chan string)
	wg.Add(1)
	go ProcessDir(ch, fp)
	
	ch <- path 

	wg.Wait()
}

var wg sync.WaitGroup

func ProcessDir(ch chan string, fp func(string)) {
	node := <-ch

	files, _ := ioutil.ReadDir(node)
	
	for _, file := range files {
		if file.IsDir() {
			directory := node + "/" + file.Name()

			wg.Add(1)
			go ProcessDir(ch, fp)
			ch <- directory
		} else if !file.IsDir() {
			wg.Add(1)
			go ProcessFile(node + "/" + file.Name(), fp)
		}
	}

	wg.Done()
}

func ProcessFile(f string, fp func(string)) {
	fp (f)
	
	wg.Done()
}
