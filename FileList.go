package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

// FileList scan directory for nsiBudget xml files
func FileList(where, pattern string, ch chan string) {
	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {

			log.Fatalln(err)
		}

		matched, err := filepath.Match(pattern, info.Name())

		if err != nil {

			log.Fatalln(err)
		}
		if !info.IsDir() && matched {
			ch <- path
			//log.Println("From1:", info.Name(), !info.IsDir() && matched, len(ch))
		}

		return nil
	}
	log.Println("Scaning for ", pattern, " into ", where)
	filepath.Walk(where, walk)
	//close(ch)
}

func FileListWorker(pool chan int, where []string, pattern string, ch chan string) {
	sc := new(sync.WaitGroup)
	log.Println("FileListWorker :: Start")
	pool <- 1
	for p := range where {
		path := where[p]
		_, err := os.Stat(path)
		if err == nil {
			pool <- 1
			sc.Add(1)
			go func(sc *sync.WaitGroup, where, pattern string, ch chan string) {
				FileList(where, pattern, ch)
				<-pool
				sc.Done()
			}(sc, path, pattern, ch)
		}
	}
	sc.Wait()
	close(ch)
	log.Println("FileListWorker :: End")

}
