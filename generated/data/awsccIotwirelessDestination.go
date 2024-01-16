package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotwirelessDestination = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Destination arn. Returned after successful create.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Destination description",
        "description_kind": "plain",
        "type": "string"
      },
      "expression": {
        "computed": true,
        "description": "Destination expression",
        "description_kind": "plain",
        "type": "string"
      },
      "expression_type": {
        "computed": true,
        "description": "Must be RuleName",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Unique name of destination",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "AWS role ARN that grants access",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the destination.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTWireless::Destination",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotwirelessDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotwirelessDestination), &result)
	return &result
}
