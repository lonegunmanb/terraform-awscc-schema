package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccDmsReplicationTaskSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccDmsReplicationTaskSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
