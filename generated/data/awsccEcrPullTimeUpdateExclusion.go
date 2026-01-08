package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrPullTimeUpdateExclusion = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_arn": {
        "computed": true,
        "description": "The ARN of the IAM principal to remove from the pull time update exclusion list.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECR::PullTimeUpdateExclusion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrPullTimeUpdateExclusionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPullTimeUpdateExclusion), &result)
	return &result
}
