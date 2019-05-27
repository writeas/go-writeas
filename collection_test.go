package writeas

import (
	"fmt"
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

func TestGetCollectionPost(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetCollectionPost("blog", "extending-write-as")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	}

	if res == nil {
		t.Errorf("No post returned!")
	}

	if len(res.Content) == 0 {
		t.Errorf("Post content is empty!")
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

	if err := wac.DeleteCollection(c.Alias); err != nil {
		t.Fatalf("Unable to delete collection %q: %v", alias, err)
	}
}

func TestDeleteCollectionUnauthenticated(t *testing.T) {
	wac := NewDevClient()

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-does-not-exist-%v", now)
	err := wac.DeleteCollection(alias)
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
