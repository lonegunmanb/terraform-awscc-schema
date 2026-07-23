package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccStoragegatewayTapePoolSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccStoragegatewayTapePoolSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
