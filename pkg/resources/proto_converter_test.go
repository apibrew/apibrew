package resources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtoDescriptorToResource(t *testing.T) {
	assert.Equal(t, ResourceResource.Properties, ResourceResourceNew.Properties)
}
