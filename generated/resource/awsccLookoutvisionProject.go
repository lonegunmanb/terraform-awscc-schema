package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_name": {
        "description": "The name of the Amazon Lookout for Vision project",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLookoutvisionProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutvisionProject), &result)
	return &result
}
