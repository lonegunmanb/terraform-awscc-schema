package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrPullTimeUpdateExclusion = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_arn": {
        "description": "Principal arn that should not update image pull times.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ECR::PullTimeUpdateExclusion controls the exclusion configuration for ecr image pull time update.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrPullTimeUpdateExclusionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPullTimeUpdateExclusion), &result)
	return &result
}
