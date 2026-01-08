package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of KMS key used for the connector",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the connector",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_name": {
        "description": "The third-party provider configuration for the connector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "jira_cloud": {
              "computed": true,
              "description": "The initial configuration settings required to establish an integration between Security Hub and Jira Cloud",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "project_key": {
                    "computed": true,
                    "description": "The project key for a Jira Cloud instance",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_now": {
              "computed": true,
              "description": "The initial configuration settings required to establish an integration between Security Hub and ServiceNow ITSM",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_name": {
                    "computed": true,
                    "description": "The instance name of ServiceNow ITSM",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the ServiceNow credentials",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource schema for AWS::SecurityHub::ConnectorV2",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubConnectorV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubConnectorV2), &result)
	return &result
}
