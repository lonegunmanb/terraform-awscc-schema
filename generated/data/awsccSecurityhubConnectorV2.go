package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubConnectorV2 = `{
  "block": {
    "attributes": {
      "connector_arn": {
        "computed": true,
        "description": "The ARN of the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_id": {
        "computed": true,
        "description": "The ID of the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_status": {
        "computed": true,
        "description": "The status of the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp formatted in ISO8601",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the connector",
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
        "description": "The ARN of KMS key used for the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "last_checked_at": {
        "computed": true,
        "description": "The timestamp formatted in ISO8601",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp formatted in ISO8601",
        "description_kind": "plain",
        "type": "string"
      },
      "message": {
        "computed": true,
        "description": "The message of the connector status change",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The provider configuration of the connector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "jira_cloud": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_status": {
                    "computed": true,
                    "description": "The auth status of the connector",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auth_url": {
                    "computed": true,
                    "description": "The authUrl of the JiraCloud connector",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cloud_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "project_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_now": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_status": {
                    "computed": true,
                    "description": "The auth status of the connector",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description": "The ARN of secrets manager containing ClientId and ClientSecret",
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
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::ConnectorV2",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubConnectorV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubConnectorV2), &result)
	return &result
}
