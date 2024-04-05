package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesthinclientEnvironment = `{
  "block": {
    "attributes": {
      "activation_code": {
        "computed": true,
        "description": "Activation code for devices associated with environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The environment ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp in unix epoch format when environment was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_software_set_id": {
        "computed": true,
        "description": "The ID of the software set to apply.",
        "description_kind": "plain",
        "type": "string"
      },
      "desktop_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0.",
        "description_kind": "plain",
        "type": "string"
      },
      "desktop_endpoint": {
        "computed": true,
        "description": "The URL for the identity provider login (only for environments that use AppStream 2.0).",
        "description_kind": "plain",
        "type": "string"
      },
      "desktop_type": {
        "computed": true,
        "description": "The type of VDI.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "Unique identifier of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": "A specification for a time window to apply software updates.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apply_time_of": {
              "computed": true,
              "description": "The desired time zone maintenance window.",
              "description_kind": "plain",
              "type": "string"
            },
            "days_of_the_week": {
              "computed": true,
              "description": "The date of maintenance window.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "end_time_hour": {
              "computed": true,
              "description": "The hour end time of maintenance window.",
              "description_kind": "plain",
              "type": "number"
            },
            "end_time_minute": {
              "computed": true,
              "description": "The minute end time of maintenance window.",
              "description_kind": "plain",
              "type": "number"
            },
            "start_time_hour": {
              "computed": true,
              "description": "The hour start time of maintenance window.",
              "description_kind": "plain",
              "type": "number"
            },
            "start_time_minute": {
              "computed": true,
              "description": "The minute start time of maintenance window.",
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description": "The type of maintenance window.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "pending_software_set_id": {
        "computed": true,
        "description": "The ID of the software set that is pending to be installed.",
        "description_kind": "plain",
        "type": "string"
      },
      "pending_software_set_version": {
        "computed": true,
        "description": "The version of the software set that is pending to be installed.",
        "description_kind": "plain",
        "type": "string"
      },
      "registered_devices_count": {
        "computed": true,
        "description": "Number of devices registered to the environment.",
        "description_kind": "plain",
        "type": "number"
      },
      "software_set_compliance_status": {
        "computed": true,
        "description": "Describes if the software currently installed on all devices in the environment is a supported version.",
        "description_kind": "plain",
        "type": "string"
      },
      "software_set_update_mode": {
        "computed": true,
        "description": "An option to define which software updates to apply.",
        "description_kind": "plain",
        "type": "string"
      },
      "software_set_update_schedule": {
        "computed": true,
        "description": "An option to define if software updates should be applied within a maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp in unix epoch format when environment was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WorkSpacesThinClient::Environment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspacesthinclientEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesthinclientEnvironment), &result)
	return &result
}
