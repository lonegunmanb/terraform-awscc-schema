package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the group.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
        "description": "A name for the group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "description": "The namespace that you want the group to be a part of.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_id": {
        "computed": true,
        "description": "The principal ID of the group.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::Group Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightGroup), &result)
	return &result
}
