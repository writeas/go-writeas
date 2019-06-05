package writeas

import "testing"

func TestAuthentication(t *testing.T) {
	dwac := NewClient(devAPIURL)

	// Log in
	_, err := dwac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}

	// Log out
	err = dwac.LogOut()
	if err != nil {
		t.Fatalf("Unable to log out: %v", err)
	}
}
