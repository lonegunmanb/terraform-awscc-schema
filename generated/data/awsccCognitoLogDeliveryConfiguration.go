package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoLogDeliveryConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "event_source": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "log_level": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::LogDeliveryConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoLogDeliveryConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoLogDeliveryConfiguration), &result)
	return &result
}
