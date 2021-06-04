package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	//need a io.reader
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 200)

	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
}

func hashBucket(word string) int {
	return int(word[0])
}
