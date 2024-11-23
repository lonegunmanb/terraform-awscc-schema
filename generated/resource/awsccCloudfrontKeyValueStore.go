package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The source type of the import source for the key value store.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "key_value_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the key value store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The key value store. Use this to separate data from function code, allowing you to update data without having to publish a new version of a function. The key value store holds keys and their corresponding values.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontKeyValueStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontKeyValueStore), &result)
	return &result
}
