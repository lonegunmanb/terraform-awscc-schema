package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsDataProvider = `{
  "block": {
    "attributes": {
      "data_provider_arn": {
        "computed": true,
        "description": "The data provider ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_provider_creation_time": {
        "computed": true,
        "description": "The data provider creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_provider_identifier": {
        "computed": true,
        "description": "The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_provider_name": {
        "computed": true,
        "description": "The property describes a name to identify the data provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The optional description of the data provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine": {
        "description": "The property describes a data engine for the data provider.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exact_settings": {
        "computed": true,
        "description": "The property describes the exact settings which can be modified",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "settings": {
        "computed": true,
        "description": "The property identifies the exact type of settings for the data provider.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DMS::DataProvider",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsDataProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsDataProvider), &result)
	return &result
}
