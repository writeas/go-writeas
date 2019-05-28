package writeas

import (
	"fmt"
	"testing"
)

func TestPostRoundTrip(t *testing.T) {
	var id, token string
	dwac := NewClient()
	t.Run("Create post", func(t *testing.T) {
		p, err := dwac.CreatePost(&PostParams{
			Title:   "Title!",
			Content: "This is a post.",
			Font:    "sans",
		})
		if err != nil {
			t.Errorf("Post create failed: %v", err)
			return
		}
		t.Logf("Post created: %+v", p)
		id, token = p.ID, p.Token
	})
	t.Run("Get post", func(t *testing.T) {
		res, err := dwac.GetPost(id)
		if err != nil {
			t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
		} else {
			t.Logf("Post: %+v", res)
			if res.Content != "This is a post." {
				t.Errorf("Unexpected fetch results: %+v\n", res)
			}
		}
	})
	t.Run("Update post", func(t *testing.T) {
		p, err := dwac.UpdatePost(id, token, &PostParams{
			Content: "Now it's been updated!",
		})
		if err != nil {
			t.Errorf("Post update failed: %v", err)
			return
		}
		t.Logf("Post updated: %+v", p)
	})
	t.Run("Delete post", func(t *testing.T) {
		err := dwac.DeletePost(id, token)
		if err != nil {
			t.Errorf("Post delete failed: %v", err)
			return
		}
		t.Logf("Post deleted!")
	})
}

func TestPinUnPin(t *testing.T) {
	dwac := NewDevClient()
	_, err := dwac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer dwac.LogOut()

	t.Run("Pin post", func(t *testing.T) {
		err := dwac.PinPost("tester", &PinnedPostParams{ID: "olx6uk7064heqltf"})
		if err != nil {
			t.Fatalf("Pin failed: %v", err)
		}
	})
	t.Run("Unpin post", func(t *testing.T) {
		err := dwac.UnpinPost("tester", &PinnedPostParams{ID: "olx6uk7064heqltf"})
		if err != nil {
			t.Fatalf("Unpin failed: %v", err)
		}
	})
}

func ExampleClient_CreatePost() {
	dwac := NewDevClient()

	// Publish a post
	p, err := dwac.CreatePost(&PostParams{
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
