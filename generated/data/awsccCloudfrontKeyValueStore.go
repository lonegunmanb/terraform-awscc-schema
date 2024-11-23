package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontKeyValueStore = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "computed": true,
        "description": "A comment for the key value store.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_source": {
        "computed": true,
        "description": "The import source for the key value store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the import source for the key value store.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The source type of the import source for the key value store.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "key_value_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the key value store.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::KeyValueStore",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontKeyValueStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontKeyValueStore), &result)
	return &result
}
