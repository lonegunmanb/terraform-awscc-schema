package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsccRtbfabricResponderGatewaySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsccRtbfabricResponderGatewaySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
