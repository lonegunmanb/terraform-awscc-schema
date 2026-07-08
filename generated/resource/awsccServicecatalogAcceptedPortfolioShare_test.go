package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccServicecatalogAcceptedPortfolioShareSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccServicecatalogAcceptedPortfolioShareSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
