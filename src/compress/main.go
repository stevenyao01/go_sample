package main

import (
	"sync"
)

func main(){

	wg := sync.WaitGroup{}
	mode := []CompressMode{GzipMode,FlateMode,SnappyMode,LzoMode,ZlibMode}
	for i := 0;i< len(mode);i++{
		wg.Add(1)
		go func(i int) {
			c,err := New("testdata/fiction.txt")
			if err != nil {
				panic(err)
			}
			c,err = c.WithMode(mode[i])
			if err != nil {
				panic(err)
			}
			c.Compress(true)
			wg.Done()
		}(i)
	}
	wg.Wait()
}