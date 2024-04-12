package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderWorkflow = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the workflow.",
        "description_kind": "plain",
        "type": "string"
      },
      "change_description": {
        "computed": true,
        "description": "The change description of the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data": {
        "computed": true,
        "description": "The data of the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The KMS key identifier used to encrypt the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the workflow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "description": "The type of the workflow denotes whether the workflow is used to build, test, or distribute.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "The uri of the workflow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description": "The version of the workflow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::Workflow",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderWorkflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderWorkflow), &result)
	return &result
}
