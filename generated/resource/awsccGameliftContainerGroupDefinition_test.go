package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccGameliftContainerGroupDefinitionSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccGameliftContainerGroupDefinitionSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
