package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsLogGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The CloudWatch log group ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_protection_policy": {
        "computed": true,
        "description": "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per topic.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_group_class": {
        "computed": true,
        "description": "The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS, with STANDARD being the default class",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_group_name": {
        "computed": true,
        "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "retention_in_days": {
        "computed": true,
        "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, and 3653.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::Logs::LogGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsLogGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsLogGroup), &result)
	return &result
}
