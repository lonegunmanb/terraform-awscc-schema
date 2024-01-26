package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccRefactorspacesRouteSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccRefactorspacesRouteSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}