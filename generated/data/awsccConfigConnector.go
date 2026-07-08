package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConnector = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_configuration": {
        "computed": true,
        "description": "The configuration for the connector that specifies the third-party cloud provider connection details.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure": {
              "computed": true,
              "description": "The configuration for connecting to Microsoft Azure.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_identifier": {
                    "computed": true,
                    "description": "The Azure client (application) identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_identifier": {
                    "computed": true,
                    "description": "The Azure tenant identifier.",
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
      "created_time": {
        "computed": true,
        "description": "The time at which the connector was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the connector. AWS Config automatically assigns the name when creating the Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Config::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConnector), &result)
	return &result
}
