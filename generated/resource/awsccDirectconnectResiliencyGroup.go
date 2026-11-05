package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectResiliencyGroup = `{
  "block": {
    "attributes": {
      "connection_arns": {
        "computed": true,
        "description": "The ARNs of the connections that are members of the resiliency group. When specified, this is the complete, authoritative membership list: on update, connections not present are disassociated, and an explicit empty list removes all members. Omit the property to leave the existing membership unmanaged (unchanged on update). On create, an omitted or empty list means no connections are associated.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "intended_resiliency_model": {
        "computed": true,
        "description": "The resiliency model that the resiliency group is intended to meet. Required when creating a resiliency group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resiliency_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resiliency group.",
        "description_kind": "plain",
        "type": "string"
      },
      "resiliency_group_id": {
        "computed": true,
        "description": "The ID of the resiliency group.",
        "description_kind": "plain",
        "type": "string"
      },
      "resiliency_group_name": {
        "description": "The name of the resiliency group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the resiliency group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::DirectConnect::ResiliencyGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDirectconnectResiliencyGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectResiliencyGroup), &result)
	return &result
}
