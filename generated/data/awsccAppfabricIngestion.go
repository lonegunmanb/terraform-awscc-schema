package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppfabricIngestion = `{
  "block": {
    "attributes": {
      "app": {
        "computed": true,
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_bundle_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the app bundle for the ingestion.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ingestion.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the ingestion was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ingestion_type": {
        "computed": true,
        "description": "The type of the ingestion.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The status of the ingestion.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of the key-value pairs of the tag or tags to assign to the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tenant_id": {
        "computed": true,
        "description": "The ID of the application tenant.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when the ingestion was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppFabric::Ingestion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppfabricIngestionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppfabricIngestion), &result)
	return &result
}
