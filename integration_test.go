package gortt

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestClientBasicallyWorking(t *testing.T) {
	user := os.Getenv("RTTAPIUSER")
	pass := os.Getenv("RTTAPIPASS")

	if user == "" || pass == "" {
		t.Skipf("set RTTAPIUSER/RTTAPIPASS for integration tests")
	}

	cli, err := NewClient("", "")
	if err != nil {
		t.Errorf("error creating client: %v", err)
	}

	// Do we get the user and pass from the environment?
	if cli.Username != user {
		t.Errorf("didn't get username out of environment")
	}
	if cli.Password != pass {
		t.Errorf("didn't get username out of environment")
	}

	// Get some trains
	r, err := cli.GetServicesFromStationForDate(context.Background(), "SAL", time.Now())
	if err != nil {
		t.Errorf("error reading trains: %v", err)
	}

	if r.Location.Name != "Salisbury" {
		t.Errorf("got wrong name: expected 'Salisbury', got %q", r.Location.Name)
	}

	if len(r.Services) == 0 {
		t.Errorf("got no services: check any trains are actually running before commencing debugging")
	}

}
