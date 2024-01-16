package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamDirectoryConfig = `{
  "block": {
    "attributes": {
      "certificate_based_auth_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_authority_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "directory_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organizational_unit_distinguished_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "service_account_credentials": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "account_password": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppStream::DirectoryConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamDirectoryConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamDirectoryConfig), &result)
	return &result
}
