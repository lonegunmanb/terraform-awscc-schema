package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesWorkspaceIpGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IP access control group.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_desc": {
        "computed": true,
        "description": "The description of the group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description": "The identifier of the IP access control group.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "description": "The name of the group.",
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
      "tags": {
        "computed": true,
        "description": "The tags for the IP access control group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_rules": {
        "computed": true,
        "description": "The rules for the IP access control group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_rule": {
              "computed": true,
              "description": "The IP address range, in CIDR notation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_desc": {
              "computed": true,
              "description": "The description of the rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for an IP access control group for Amazon WorkSpaces.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspacesWorkspaceIpGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesWorkspaceIpGroup), &result)
	return &result
}
