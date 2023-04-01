package writeas

import (
	"context"
	"testing"
)

func TestClient_CreateContributor(t *testing.T) {
	c := NewClientWith(Config{URL: "http://localhost:7777/api"})
	ctx := context.Background()
	_, err := c.LogIn(ctx, "test", "test")
	if err != nil {
		t.Fatalf("login: %s", err)
	}

	tests := []struct {
		name  string
		AName string
		ASlug string
		AOrg  string
	}{
		{
			name:  "good",
			AName: "Bob Contrib",
			ASlug: "bob",
			AOrg:  "write-as",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err = c.CreateContributor(ctx, &AuthorParams{
				Name:     test.AName,
				Slug:     test.ASlug,
				OrgAlias: test.AOrg,
			})
			if err != nil {
				t.Fatalf("create %s: %s", test.name, err)
			}
		})
	}
}
