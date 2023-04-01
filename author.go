package writeas

import (
	"context"
	"fmt"
	"net/http"
)

type (
	// Author represents a Write.as author.
	Author struct {
		User *User
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	// AuthorParams are used to create or update a Write.as author.
	AuthorParams struct {
		// Name is the public display name of the Author.
		Name string `json:"name"`

		// Slug is the optional slug for the Author.
		Slug string `json:"slug"`

		// OrgAlias is the alias of the organization the Author belongs to.
		OrgAlias string `json:"-"`
	}
)

// CreateContributor creates a new contributor on the given organization.
func (c *Client) CreateContributor(ctx context.Context, sp *AuthorParams) (*Author, error) {
	if sp.OrgAlias == "" {
		return nil, fmt.Errorf("AuthorParams.OrgAlias is required.")
	}

	a := &Author{}
	env, err := c.post(ctx, "/organizations/"+sp.OrgAlias+"/contributors", sp, a)
	if err != nil {
		return nil, err
	}

	var ok bool
	if a, ok = env.Data.(*Author); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}

	status := env.Code
	if status != http.StatusCreated {
		if status == http.StatusBadRequest {
			return nil, fmt.Errorf("Bad request: %s", env.ErrorMessage)
		}
		return nil, fmt.Errorf("Problem creating author: %d. %s\n", status, env.ErrorMessage)
	}
	return a, nil
}
