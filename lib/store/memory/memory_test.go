package memory

import (
	"testing"

	"github.com/bentemple/anubis/lib/store/storetest"
)

func TestImpl(t *testing.T) {
	storetest.Common(t, factory{}, nil)
}
