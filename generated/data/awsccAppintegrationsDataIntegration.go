package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppintegrationsDataIntegration = `{
  "block": {
    "attributes": {
      "data_integration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the data integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_integration_id": {
        "computed": true,
        "description": "The unique identifer of the data integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The data integration description.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_configuration": {
        "computed": true,
        "description": "The configuration for what files should be pulled from the source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "Restrictions for what files should be pulled from the source.",
              "description_kind": "plain",
              "type": [
                "map",
                [
                  "list",
                  "string"
                ]
              ]
            },
            "folders": {
              "computed": true,
              "description": "Identifiers for the source folders to pull all files from recursively.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
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
      "kms_key": {
        "computed": true,
        "description": "The KMS key of the data integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_configuration": {
        "computed": true,
        "description": "The configuration for what data should be pulled from the source.",
        "description_kind": "plain",
        "type": [
          "map",
          [
            "map",
            [
              "list",
              "string"
            ]
          ]
        ]
      },
      "schedule_config": {
        "computed": true,
        "description": "The name of the data and how often it should be pulled from the source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "first_execution_from": {
              "computed": true,
              "description": "The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.",
              "description_kind": "plain",
              "type": "string"
            },
            "object": {
              "computed": true,
              "description": "The name of the object to pull from the data source.",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule_expression": {
              "computed": true,
              "description": "How often the data should be pulled from data source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_uri": {
        "computed": true,
        "description": "The URI of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the data integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key to identify the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Corresponding tag value for the key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::AppIntegrations::DataIntegration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppintegrationsDataIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppintegrationsDataIntegration), &result)
	return &result
}
