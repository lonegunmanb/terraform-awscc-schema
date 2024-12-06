package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsIntegration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_name": {
        "description": "User provided identifier for integration, unique to the user account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_status": {
        "computed": true,
        "description": "Status of creation for the Integration and its resources",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "description": "The type of the Integration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_config": {
        "description": "OpenSearchResourceConfig for the given Integration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "open_search_resource_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dashboard_viewer_principals": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_source_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "retention_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Schema for Logs Integration Resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsIntegration), &result)
	return &result
}
