package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodestarconnectionsRepositoryLink = `{
  "block": {
    "attributes": {
      "connection_arn": {
        "description": "The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used.",
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
      "owner_id": {
        "description": "the ID of the entity that owns the repository.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_type": {
        "computed": true,
        "description": "The name of the external provider where your third-party code repository is configured.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_link_arn": {
        "computed": true,
        "description": "A unique Amazon Resource Name (ARN) to designate the repository link.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_link_id": {
        "computed": true,
        "description": "A UUID that uniquely identifies the RepositoryLink.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_name": {
        "description": "The repository for which the link is being created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Specifies the tags applied to a RepositoryLink.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodestarconnectionsRepositoryLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodestarconnectionsRepositoryLink), &result)
	return &result
}
