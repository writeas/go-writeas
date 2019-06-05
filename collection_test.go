package writeas

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var testConfig = Config{
	URL: devAPIURL,
}

func TestGetCollection(t *testing.T) {
	dwac := NewClientWith(testConfig)

	res, err := dwac.GetCollection("tester")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	}
	if res == nil {
		t.Error("Expected collection to not be nil")
	}
}

func TestGetCollectionPosts(t *testing.T) {
	dwac := NewClientWith(testConfig)
	posts := []Post{}

	t.Run("Get all posts in collection", func(t *testing.T) {
		res, err := dwac.GetCollectionPosts("tester")
		if err != nil {
			t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
		}
		if len(*res) == 0 {
			t.Error("Expected at least on post in collection")
		}
		posts = *res
	})
	t.Run("Get one post from collection", func(t *testing.T) {
		res, err := dwac.GetCollectionPost("tester", posts[0].Slug)
		if err != nil {
			t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
		}

		if res == nil {
			t.Errorf("No post returned!")
		}

		if len(res.Content) == 0 {
			t.Errorf("Post content is empty!")
		}
	})
}

func TestGetUserCollections(t *testing.T) {
	wac := NewClientWith(testConfig)
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
	wac := NewClientWith(testConfig)
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
	wac := NewClientWith(testConfig)

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
