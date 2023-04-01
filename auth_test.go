package writeas

import (
	"context"
	"testing"
)

func TestAuthentication(t *testing.T) {
	dwac := NewDevClient()
	ctx := context.Background()

	// Log in
	_, err := dwac.LogIn(ctx, "demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}

	// Log out
	err = dwac.LogOut(ctx)
	if err != nil {
		t.Fatalf("Unable to log out: %v", err)
	}
}
