package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3) // รอ 3 go routine ทำงานเสร็จ

	start := time.Now()
	go works("go", wg)
	go works("Rust", wg)
	go works("React", wg)

	wg.Wait()
	fmt.Println(time.Since(start))
}

func works(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	b := sha256.Sum256([]byte(s))
	fmt.Println(base64.StdEncoding.EncodeToString(b[:]))
}
