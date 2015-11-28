package main

import (
	"encoding/json"
	"fmt"
	"os"
  "io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// decode JSON from file to struct
func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &post)  
  return
}

// Iterative Fibonacci
func fibonacciIterative(n int) int {
    current, prev := 0, 1
    for i := 0; i < n; i++ {
        current, prev = current + prev, current
    }
    return current
}

// Recursive Fibonacci
func fibonacciRecursive(n int) int {
    if n < 2 {
        return n
    } 
    return fibonacciRecursive(n - 1) + fibonacciRecursive(n - 2)
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Post is", post)
}
