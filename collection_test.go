package writeas

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetCollection(t *testing.T) {
	dwac := NewDevClient()

	res, err := dwac.GetCollection(context.Background(), "tester")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	}
	if res == nil {
		t.Error("Expected collection to not be nil")
	}
}

func TestGetCollectionPosts(t *testing.T) {
	dwac := NewDevClient()
	posts := []Post{}
	ctx := context.Background()

	t.Run("Get all posts in collection", func(t *testing.T) {
		res, err := dwac.GetCollectionPosts(ctx, "tester")
		if err != nil {
			t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
		}
		if len(*res) == 0 {
			t.Error("Expected at least on post in collection")
		}
		posts = *res
	})
	t.Run("Get one post from collection", func(t *testing.T) {
		res, err := dwac.GetCollectionPost(ctx, "tester", posts[0].Slug)
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
	wac := NewDevClient()
	ctx := context.Background()
	_, err := wac.LogIn(ctx, "demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut(ctx)

	res, err := wac.GetUserCollections(ctx)
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
	ctx := context.Background()
	_, err := wac.LogIn(ctx, "demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut(ctx)

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-%v", now)
	c, err := wac.CreateCollection(ctx, &CollectionParams{
		Alias: alias,
		Title: fmt.Sprintf("Test Collection %v", now),
	})
	if err != nil {
		t.Fatalf("Unable to create collection %q: %v", alias, err)
	}

	if err := wac.DeleteCollection(ctx, c.Alias); err != nil {
		t.Fatalf("Unable to delete collection %q: %v", alias, err)
	}
}

func TestDeleteCollectionUnauthenticated(t *testing.T) {
	wac := NewDevClient()

	now := time.Now().Unix()
	alias := fmt.Sprintf("test-collection-does-not-exist-%v", now)
	err := wac.DeleteCollection(context.Background(), alias)
	if err == nil {
		t.Fatalf("Should not be able to delete collection %q unauthenticated.", alias)
	}

	if !strings.Contains(err.Error(), "Not authenticated") {
		t.Fatalf("Error message should be more informative: %v", err)
	}
}
