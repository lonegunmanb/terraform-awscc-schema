package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsccBatchServiceEnvironmentsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsccBatchServiceEnvironmentsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
