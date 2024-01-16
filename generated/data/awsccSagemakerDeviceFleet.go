package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerDeviceFleet = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description for the edge device fleet",
        "description_kind": "plain",
        "type": "string"
      },
      "device_fleet_name": {
        "computed": true,
        "description": "The name of the edge device fleet",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_config": {
        "computed": true,
        "description": "S3 bucket and an ecryption key id (if available) to store outputs for the fleet",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The KMS key id used for encryption on the S3 bucket",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_output_location": {
              "computed": true,
              "description": "The Amazon Simple Storage (S3) bucket URI",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description": "Role associated with the device fleet",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Associate tags with the resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::DeviceFleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerDeviceFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerDeviceFleet), &result)
	return &result
}
