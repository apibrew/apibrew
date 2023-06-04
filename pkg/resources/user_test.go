package resources

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestResourceMap(t *testing.T) {
	got := UserResource
	want := UserResourceOld

	// check if got and want are deep equal
	assert.Equal(t, got.SourceConfig, want.SourceConfig, "SourceConfig not equal")

	for i := 0; i < len(got.Properties); i++ {
		fmt.Println(got.Properties[i])
		fmt.Println(want.Properties[i])
		assert.Equal(t, got.Properties[i], want.Properties[i], "Property not equal")
	}

	assert.Equal(t, got.Properties, want.Properties, "Properties not equal")
	assert.Equal(t, got.Indexes, want.Indexes, "Indexes not equal")
	assert.Equal(t, got, want, "Resource not equal")
}
