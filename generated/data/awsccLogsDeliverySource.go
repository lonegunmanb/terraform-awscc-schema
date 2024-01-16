package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsDeliverySource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery source.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_type": {
        "computed": true,
        "description": "The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique name of the Log source.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN of the resource that will be sending the logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arns": {
        "computed": true,
        "description": "This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "service": {
        "computed": true,
        "description": "The AWS service that is sending logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that have been assigned to this delivery source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Logs::DeliverySource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsDeliverySourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDeliverySource), &result)
	return &result
}
