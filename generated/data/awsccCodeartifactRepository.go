package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeartifactRepository = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A text description of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The name of the domain that contains the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_owner": {
        "computed": true,
        "description": "The 12-digit account ID of the AWS account that owns the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_connections": {
        "computed": true,
        "description": "A list of external connections associated with the repository.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the repository. This is used for GetAtt",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_policy_document": {
        "computed": true,
        "description": "The access control resource policy on the provided repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_name": {
        "computed": true,
        "description": "The name of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "upstreams": {
        "computed": true,
        "description": "A list of upstream repositories associated with the repository.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::CodeArtifact::Repository",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodeartifactRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeartifactRepository), &result)
	return &result
}
