package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserviceApplication = `{
  "block": {
    "attributes": {
      "app_configs": {
        "computed": true,
        "description": "List of application configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The configuration key",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The configuration value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "application_id": {
        "computed": true,
        "description": "The identifier of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) format.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_sources": {
        "computed": true,
        "description": "List of data sources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_source_arn": {
              "computed": true,
              "description": "The ARN of the data source.",
              "description_kind": "plain",
              "type": "string"
            },
            "data_source_description": {
              "computed": true,
              "description": "Description of the data source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "endpoint": {
        "computed": true,
        "description": "The endpoint for the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "iam_identity_center_options": {
        "computed": true,
        "description": "Options for configuring IAM Identity Center",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Whether IAM Identity Center is enabled.",
              "description_kind": "plain",
              "type": "bool"
            },
            "iam_identity_center_instance_arn": {
              "computed": true,
              "description": "The ARN of the IAM Identity Center instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "iam_role_for_identity_center_application_arn": {
              "computed": true,
              "description": "The ARN of the IAM role for Identity Center application.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::OpenSearchService::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserviceApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserviceApplication), &result)
	return &result
}
