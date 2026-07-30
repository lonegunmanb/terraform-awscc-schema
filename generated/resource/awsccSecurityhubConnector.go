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
        "description_kind": "plain",
        "type": "string"
      },
      "connector_id": {
        "computed": true,
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
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the connector.",
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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The error code that identifies the type of health issue.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "A human-readable message that describes the health issue.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the connector.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_name": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure": {
              "description": "The configuration for connecting to an Azure environment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_config_connector_arn": {
                    "description": "The ARN of the multi-cloud configuration connector used to establish the connection to Azure.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "azure_regions": {
                    "description": "The list of Azure regions to monitor.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "scope_configuration": {
                    "description": "The scope configuration that defines which Azure resources are monitored.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "scope_type": {
                          "description": "The type of scope. Valid values are ` + "`" + `` + "`" + `tenant` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `subscription` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "scope_values": {
                          "computed": true,
                          "description": "The list of scope values, such as subscription IDs, when the scope type is ` + "`" + `` + "`" + `subscription` + "`" + `` + "`" + `.",
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
    "description": "Creates a connector to a third-party cloud provider in Security Hub CSPM. A connector establishes a connection between Security Hub CSPM and a third-party cloud provider, enabling Security Hub CSPM to ingest security findings and resource data from the connected environment.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubConnector), &result)
	return &result
}
