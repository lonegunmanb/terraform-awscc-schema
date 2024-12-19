package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccNetworkmanagerDirectConnectGatewayAttachmentSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccNetworkmanagerDirectConnectGatewayAttachmentSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
