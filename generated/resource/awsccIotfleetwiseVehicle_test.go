package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccIotfleetwiseVehicleSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccIotfleetwiseVehicleSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
