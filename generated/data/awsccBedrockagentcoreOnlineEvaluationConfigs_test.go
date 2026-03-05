package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAwsccBedrockagentcoreOnlineEvaluationConfigsSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AwsccBedrockagentcoreOnlineEvaluationConfigsSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
