package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerPartnerApp = `{
  "block": {
    "attributes": {
      "app_version": {
        "computed": true,
        "description": "The version of the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_config": {
        "computed": true,
        "description": "A collection of settings that specify the maintenance schedule for the PartnerApp.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "admin_users": {
              "computed": true,
              "description": "A list of users with administrator privileges for the PartnerApp.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "arguments": {
              "computed": true,
              "description": "A list of arguments to pass to the PartnerApp.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the created PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "auth_type": {
        "computed": true,
        "description": "The Auth type of PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "base_url": {
        "computed": true,
        "description": "The AppServerUrl based on app and account-info.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "The client token for the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_version_eol_date": {
        "computed": true,
        "description": "The end-of-life date for the current version of the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_auto_minor_version_upgrade": {
        "computed": true,
        "description": "Enables automatic minor version upgrades for the PartnerApp.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_iam_session_based_identity": {
        "computed": true,
        "description": "Enables IAM Session based Identity for PartnerApp.",
        "description_kind": "plain",
        "type": "bool"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The execution role for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The AWS KMS customer managed key used to encrypt the data associated with the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_config": {
        "computed": true,
        "description": "A collection of settings that specify the maintenance schedule for the PartnerApp.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maintenance_window_start": {
              "computed": true,
              "description": "The maintenance window start day and time for the PartnerApp.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the PartnerApp.",
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
      },
      "tier": {
        "computed": true,
        "description": "The tier of the PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of PartnerApp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::PartnerApp",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerPartnerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerPartnerApp), &result)
	return &result
}
