package writeas

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreatePost(t *testing.T) {
	wac := NewClient()
	p, err := wac.CreatePost(&PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		t.Errorf("Post create failed: %v", err)
		return
	}
	t.Logf("Post created: %+v", p)

	token := p.Token

	// Update post
	p, err = wac.UpdatePost(p.ID, token, &PostParams{
		Content: "Now it's been updated!",
	})
	if err != nil {
		t.Errorf("Post update failed: %v", err)
		return
	}
	t.Logf("Post updated: %+v", p)

	// Delete post
	err = wac.DeletePost(p.ID, token)
	if err != nil {
		t.Errorf("Post delete failed: %v", err)
		return
	}
	t.Logf("Post deleted!")
}

func TestGetPost(t *testing.T) {
	dwac := NewDevClient()
	res, err := dwac.GetPost("zekk5r9apum6p")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("Post: %+v", res)
		if res.Content != "This is a post." {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}

	wac := NewClient()
	res, err = wac.GetPost("3psnxyhqxy3hq")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		if !strings.HasPrefix(res.Content, "                               Write.as Blog") {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}
}

func TestPinPost(t *testing.T) {
	dwac := NewDevClient()
	_, err := dwac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer dwac.LogOut()

	err = dwac.PinPost("tester", &PinnedPostParams{ID: "olx6uk7064heqltf"})
	if err != nil {
		t.Fatalf("Pin failed: %v", err)
	}
}

func TestUnpinPost(t *testing.T) {
	dwac := NewDevClient()
	_, err := dwac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer dwac.LogOut()

	err = dwac.UnpinPost("tester", &PinnedPostParams{ID: "olx6uk7064heqltf"})
	if err != nil {
		t.Fatalf("Unpin failed: %v", err)
	}
}

func ExampleClient_CreatePost() {
	c := NewClient()

	// Publish a post
	p, err := c.CreatePost(&PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		fmt.Printf("Unable to create: %v", err)
		return
	}

	fmt.Printf("%s", p.Content)
	// Output: This is a post.
}
