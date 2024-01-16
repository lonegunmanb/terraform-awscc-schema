package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsccIamServiceLinkedRoleSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsccIamServiceLinkedRoleSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
