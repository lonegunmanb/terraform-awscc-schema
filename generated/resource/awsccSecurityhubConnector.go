package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubConnector = `{
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
        "description": "The date and time for createdAt in UTC and ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The principal that created the connector",
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
      "enablement_status": {
        "computed": true,
        "description": "The enablement status of the connector",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "issues": {
        "computed": true,
        "description": "The list of health issues associated with the connector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The code identifying the type of health issue",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The message describing the health issue",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "last_checked_at": {
        "computed": true,
        "description": "The date and time for lastCheckedAt in UTC and ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The date and time for lastUpdatedAt in UTC and ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "message": {
        "computed": true,
        "description": "The message associated with the connector status change",
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
        "description": "The CSPM provider configuration for the connector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure": {
              "description": "The configuration settings for an Azure CSPM provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_config_connector_arn": {
                    "description": "The ARN of the AWS Config connector used for the Azure integration",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "azure_regions": {
                    "description": "The list of Azure regions to include in the connector scope",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "scope_configuration": {
                    "description": "The scope configuration for an Azure connector",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "scope_type": {
                          "description": "The scope type for the Azure connector",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scope_values": {
                          "computed": true,
                          "description": "The list of scope values for the Azure connector",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
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
    "description": "Resource schema for AWS::SecurityHub::Connector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubConnector), &result)
	return &result
}
