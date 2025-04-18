package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchConsumableResource = `{
  "block": {
    "attributes": {
      "available_quantity": {
        "computed": true,
        "description": "Available Quantity of ConsumableResource.",
        "description_kind": "plain",
        "type": "number"
      },
      "consumable_resource_arn": {
        "computed": true,
        "description": "ARN of the Consumable Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "consumable_resource_name": {
        "computed": true,
        "description": "Name of ConsumableResource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "in_use_quantity": {
        "computed": true,
        "description": "In-use Quantity of ConsumableResource.",
        "description_kind": "plain",
        "type": "number"
      },
      "resource_type": {
        "description": "Type of Consumable Resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "total_quantity": {
        "description": "Total Quantity of ConsumableResource.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::Batch::ConsumableResource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBatchConsumableResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchConsumableResource), &result)
	return &result
}
