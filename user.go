package writeas

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type (
	// AuthUser represents a just-authenticated user. It contains information
	// that'll only be returned once (now) per user session.
	AuthUser struct {
		AccessToken string `json:"access_token,omitempty"`
		Password    string `json:"password,omitempty"`
		User        *User  `json:"user"`
	}

	// User represents a registered Write.as user.
	User struct {
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Created  time.Time `json:"created"`

		// Optional properties
		Subscription *UserSubscription `json:"subscription"`
	}

	// UserSubscription contains information about a user's Write.as
	// subscription.
	UserSubscription struct {
		Name       string    `json:"name"`
		Begin      time.Time `json:"begin"`
		End        time.Time `json:"end"`
		AutoRenew  bool      `json:"auto_renew"`
		Active     bool      `json:"is_active"`
		Delinquent bool      `json:"is_delinquent"`
	}
)

// GetMe retrieves the authenticated User's information.
// See: https://developers.write.as/docs/api/#retrieve-authenticated-user
func (c *Client) GetMe(ctx context.Context, verbose bool) (*User, error) {
	if c.Token() == "" {
		return nil, fmt.Errorf("Unable to get user; no access token given.")
	}

	params := ""
	if verbose {
		params = "?verbose=true"
	}
	env, err := c.get(ctx, "/me"+params, nil)
	if err != nil {
		return nil, err
	}

	status := env.Code
	if status == http.StatusUnauthorized {
		return nil, fmt.Errorf("invalid or expired token")
	}

	var u *User
	var ok bool
	if u, ok = env.Data.(*User); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}

	return u, nil
}
