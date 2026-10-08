package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccScnInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The AWS account ID that owns the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The instance creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_description": {
        "computed": true,
        "description": "The instance description.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The instance identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_name": {
        "computed": true,
        "description": "The instance name.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The instance last modified timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "version_number": {
        "computed": true,
        "description": "The version number of the instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "web_app_dns_domain": {
        "computed": true,
        "description": "The WebApp DNS domain name of the instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SCN::Instance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccScnInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccScnInstance), &result)
	return &result
}
