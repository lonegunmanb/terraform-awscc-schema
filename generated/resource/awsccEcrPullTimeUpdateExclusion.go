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
        "description": "The ARN of the IAM principal to remove from the pull time update exclusion list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ARN of the IAM principal to remove from the pull time update exclusion list.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrPullTimeUpdateExclusionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPullTimeUpdateExclusion), &result)
	return &result
}
