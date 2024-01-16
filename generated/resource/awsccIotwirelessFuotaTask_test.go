package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccIotwirelessFuotaTaskSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccIotwirelessFuotaTaskSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
