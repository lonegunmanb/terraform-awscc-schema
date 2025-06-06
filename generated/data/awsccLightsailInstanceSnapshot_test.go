package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsccLightsailInstanceSnapshotSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsccLightsailInstanceSnapshotSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
