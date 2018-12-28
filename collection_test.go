package writeas

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetCollection(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetCollection("blog")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("Collection: %+v", res)
		if res.Title != "write.as" {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}
}

func TestGetCollectionPosts(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetCollectionPosts("blog")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		if len(*res) == 0 {
			t.Errorf("No posts returned!")
		}
	}
}

func TestGetUserCollections(t *testing.T) {
	wac := NewDevClient()
	_, err := wac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut()

	res, err := wac.GetUserCollections()
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("User collections: %+v", res)
		if len(*res) == 0 {
			t.Errorf("No collections returned!")
		}
	}
}

func TestCreateAndDeleteCollection(t *testing.T) {
	wac := NewDevClient()
	_, err := wac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut()

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-%v", now)
	c, err := wac.CreateCollection(&CollectionParams{
		Alias: alias,
		Title: fmt.Sprintf("Test Collection %v", now),
	})
	if err != nil {
		t.Fatalf("Unable to create collection %q: %v", alias, err)
	}

	p := &DeleteCollectionParams{Alias: c.Alias}
	if err := wac.DeleteCollection(p); err != nil {
		t.Fatalf("Unable to delete collection %q: %v", alias, err)
	}
}

func TestDeleteCollectionUnauthenticated(t *testing.T) {
	wac := NewDevClient()

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-does-not-exist-%v", now)
	p := &DeleteCollectionParams{Alias: alias}
	err := wac.DeleteCollection(p)
	if err == nil {
		t.Fatalf("Should not be able to delete collection %q unauthenticated.", alias)
	}

	if !strings.Contains(err.Error(), "Not authenticated") {
		t.Fatalf("Error message should be more informative: %v", err)
	}
}

func ExampleClient_GetCollection() {
	c := NewClient()
	coll, err := c.GetCollection("blog")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%s", coll.Title)
	// Output: write.as
}

func TestCollectPostsAnonymous(t *testing.T) {
	// Create a post anonymously.
	wac := NewDevClient()
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

	// Log in.
	if _, err := wac.LogIn("demo", "demo"); err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut()

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-%v", now)

	// Create a collection.
	_, err = wac.CreateCollection(&CollectionParams{
		Alias: alias,
		Title: fmt.Sprintf("Test Collection %v", now),
	})
	if err != nil {
		t.Fatalf("Unable to create collection %q: %v", alias, err)
	}
	defer wac.DeleteCollection(&DeleteCollectionParams{Alias: alias})

	// Move the anonymous post to this collection.
	res, err := wac.CollectPosts(&CollectPostParams{
		Alias: alias,
		Posts: []*CollectPost{
			{
				ID:    p.ID,
				Token: p.Token,
			},
		},
	})
	if err != nil {
		t.Fatalf("Could not collect post %q: %v", p.ID, err)
	}

	for _, cr := range res {
		if cr.Code != http.StatusOK {
			t.Errorf("Failed to move post: %v", cr.ErrorMessage)
		} else {
			t.Logf("Moved post %q", cr.Post.ID)
		}
	}
}
