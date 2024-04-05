package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeconnectionsConnection = `{
  "block": {
    "attributes": {
      "connection_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_name": {
        "computed": true,
        "description": "The name of the connection. Connection names must be unique in an AWS user account.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_status": {
        "computed": true,
        "description": "The current status of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_arn": {
        "computed": true,
        "description": "The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_type": {
        "computed": true,
        "description": "The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Specifies the tags applied to a connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeConnections::Connection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodeconnectionsConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeconnectionsConnection), &result)
	return &result
}
