package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoGetDevices(t *testing.T) {
	ether, err := AutoGetDevices([]string{"1.1.1.1"})
	assert.NoError(t, err)
	PrintDeviceInfo(ether)
}
