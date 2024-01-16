package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccBillingconductorPricingPlanSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccBillingconductorPricingPlanSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
