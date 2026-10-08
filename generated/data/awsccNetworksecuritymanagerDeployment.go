package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworksecuritymanagerDeployment = `{
  "block": {
    "attributes": {
      "associated_policy_list": {
        "computed": true,
        "description": "List of policies associated with this deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_arn": {
              "computed": true,
              "description": "ARN of the associated policy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "associated_scope_list": {
        "computed": true,
        "description": "List of scopes associated with this deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "scope_arn": {
              "computed": true,
              "description": "ARN of the associated scope.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "deployment_arn": {
        "computed": true,
        "description": "The ARN of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description": "Configuration settings for the deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_cross_account_visibility": {
              "computed": true,
              "description": "Whether cross-account visibility is enabled for the deployment.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deployment_description": {
        "computed": true,
        "description": "A description of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description": "The unique identifier of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_name": {
        "computed": true,
        "description": "The name of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the deployment.",
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
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "An ISO 8601 timestamp indicating when the deployment was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkSecurityManager::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworksecuritymanagerDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworksecuritymanagerDeployment), &result)
	return &result
}
