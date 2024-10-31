package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmazonmqConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon MQ configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_strategy": {
        "computed": true,
        "description": "The authentication strategy associated with the configuration. The default is SIMPLE.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_id": {
        "computed": true,
        "description": "The ID of the Amazon MQ configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data": {
        "computed": true,
        "description": "The base64-encoded XML configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_type": {
        "computed": true,
        "description": "The type of broker engine. Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version of the broker engine.",
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
        "description": "The name of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision": {
        "computed": true,
        "description": "The revision number of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Create tags when creating the configuration.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::AmazonMQ::Configuration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAmazonmqConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmazonmqConfiguration), &result)
	return &result
}