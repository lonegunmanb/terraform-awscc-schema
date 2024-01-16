package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRekognitionProject = `{
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
        "description": "The name of the project",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::Rekognition::Project type creates an Amazon Rekognition CustomLabels Project. A project is a grouping of the resources needed to create and manage Dataset and ProjectVersions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRekognitionProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRekognitionProject), &result)
	return &result
}
