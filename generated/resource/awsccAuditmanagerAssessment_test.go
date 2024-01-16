package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-awscc-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAwsccAuditmanagerAssessmentSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AwsccAuditmanagerAssessmentSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
