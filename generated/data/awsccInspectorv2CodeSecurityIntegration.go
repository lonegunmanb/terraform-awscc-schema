package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInspectorv2CodeSecurityIntegration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Code Security Integration ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "authorization_url": {
        "computed": true,
        "description": "Authorization URL for OAuth flow",
        "description_kind": "plain",
        "type": "string"
      },
      "create_integration_details": {
        "computed": true,
        "description": "Create Integration Details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gitlab_self_managed": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_token": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "Creation timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "Last update timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Code Security Integration name",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Integration Status",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "Reason for the current status",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "Integration Type",
        "description_kind": "plain",
        "type": "string"
      },
      "update_integration_details": {
        "computed": true,
        "description": "Update Integration Details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "github": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "installation_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "gitlab_self_managed": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::InspectorV2::CodeSecurityIntegration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInspectorv2CodeSecurityIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInspectorv2CodeSecurityIntegration), &result)
	return &result
}
