package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccVpclatticeListenerSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccVpclatticeListenerSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
