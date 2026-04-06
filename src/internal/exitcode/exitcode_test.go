package exitcode

import (
	"errors"
	"testing"
)

func TestFromError(t *testing.T) {
	if got := FromError(nil); got != Success {
		t.Fatalf("expected success exit code, got %d", got)
	}

	if got := FromError(Wrap(Network, errors.New("provider request failed"))); got != Network {
		t.Fatalf("expected network exit code, got %d", got)
	}

	if got := FromError(errors.New("boom")); got != Internal {
		t.Fatalf("expected internal exit code, got %d", got)
	}
}
