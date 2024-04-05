package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccBedrockAgentAliasSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccBedrockAgentAliasSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
