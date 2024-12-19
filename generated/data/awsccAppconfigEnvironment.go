package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigEnvironment = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_check": {
        "computed": true,
        "description": "On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The environment ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitors": {
        "computed": true,
        "description": "Amazon CloudWatch alarms to monitor during the deployment process.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarm_arn": {
              "computed": true,
              "description": "Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.",
              "description_kind": "plain",
              "type": "string"
            },
            "alarm_role_arn": {
              "computed": true,
              "description": "ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value can be up to 256 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::AppConfig::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppconfigEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigEnvironment), &result)
	return &result
}
