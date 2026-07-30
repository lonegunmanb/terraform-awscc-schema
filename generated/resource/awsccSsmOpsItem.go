package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmOpsItem = `{
  "block": {
    "attributes": {
      "category": {
        "computed": true,
        "description": "The category of the OpsItem.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The user who created the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time the OpsItem was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description of the OpsItem.",
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
      "last_modified_by": {
        "computed": true,
        "description": "The user who last modified the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time the OpsItem was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "ops_item_arn": {
        "computed": true,
        "description": "The ARN of the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "ops_item_id": {
        "computed": true,
        "description": "The ID of the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "ops_item_type": {
        "computed": true,
        "description": "The type of OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "The priority of the OpsItem.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "severity": {
        "computed": true,
        "description": "The severity of the OpsItem.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "description": "The origin of the OpsItem.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the OpsItem.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "title": {
        "description": "The title of the OpsItem.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the OpsItem.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::SSM::OpsItem.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmOpsItemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmOpsItem), &result)
	return &result
}
