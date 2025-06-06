package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccS3ExpressAccessPointSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccS3ExpressAccessPointSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
