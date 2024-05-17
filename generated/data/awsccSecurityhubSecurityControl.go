package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubSecurityControl = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_update_reason": {
        "computed": true,
        "description": "The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "An object that identifies the name of a control parameter, its current value, and whether it has been customized.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "value": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "boolean": {
                    "computed": true,
                    "description": "A control parameter that is a boolean.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "double": {
                    "computed": true,
                    "description": "A control parameter that is a double.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "enum": {
                    "computed": true,
                    "description": "A control parameter that is a enum.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enum_list": {
                    "computed": true,
                    "description": "A control parameter that is a list of enums.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "integer": {
                    "computed": true,
                    "description": "A control parameter that is a integer.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "integer_list": {
                    "computed": true,
                    "description": "A control parameter that is a list of integers.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "string": {
                    "computed": true,
                    "description": "A control parameter that is a string.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "string_list": {
                    "computed": true,
                    "description": "A control parameter that is a list of strings.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "value_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "security_control_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for a security control across standards, such as ` + "`" + `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1` + "`" + `. This parameter doesn't mention a specific standard.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_control_id": {
        "computed": true,
        "description": "The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::SecurityControl",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubSecurityControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubSecurityControl), &result)
	return &result
}
