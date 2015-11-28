package main

import (
	"testing"
  "time"
)

// Test the decode function
func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Post ID is not the same as post.json", post.Id)
	}
  if post.Content != "Hello World!" {
    t.Error("Post content is not the same as post.json", post.Id)
  }
}

func TestUnmarshal(t *testing.T) {
	post, err := unmarshal("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Post ID is not the same as post.json", post.Id)
	}
  if post.Content != "Hello World!" {
    t.Error("Post content is not the same as post.json", post.Id)
  }
}


// Test the encode function
func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

// Long running test case
func TestLongRunningTest(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping long running test in short mode")
  }
  time.Sleep(10 * time.Second)
}

