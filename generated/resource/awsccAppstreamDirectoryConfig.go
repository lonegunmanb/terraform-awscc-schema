package resource

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
              "optional": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "directory_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "organizational_unit_distinguished_names": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "service_account_credentials": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "account_password": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::AppStream::DirectoryConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamDirectoryConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamDirectoryConfig), &result)
	return &result
}
