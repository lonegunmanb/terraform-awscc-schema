package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutvisionProject = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_name": {
        "computed": true,
        "description": "The name of the Amazon Lookout for Vision project",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LookoutVision::Project",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLookoutvisionProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutvisionProject), &result)
	return &result
}
