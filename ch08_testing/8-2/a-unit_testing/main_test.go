package main

import (
	"testing"
	"time"
)

// ===================================Test the decode function
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
//=====================================Test  unmarshal  function
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
//=========================================================== Test the encode function
func TestEncode(t *testing.T) {
	//暂时跳过对编码函数的测试
	t.Skip("Skipping encoding for now")
}
// ====================================Long running test case   伪造一个长时间执行的测试用例
func TestLongRunningTest(t *testing.T) {
	//当传递-short参数的时候，则会执行下面的t.Skip(),这样就会跳过了
  if testing.Short() {
    t.Skip("Skipping long running test in short mode")
  }
  time.Sleep(10 * time.Second)
}

