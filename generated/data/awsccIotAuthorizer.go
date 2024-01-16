package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotAuthorizer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_caching_for_http": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "signing_disabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "token_key_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "token_signing_public_keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::IoT::Authorizer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotAuthorizer), &result)
	return &result
}
