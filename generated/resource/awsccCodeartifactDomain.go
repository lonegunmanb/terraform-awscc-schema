package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeartifactDomain = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description": "The name of the domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "The ARN of an AWS Key Management Service (AWS KMS) key associated with a domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the domain. This field is used for GetAtt",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The 12-digit account ID of the AWS account that owns the domain. This field is used for GetAtt",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions_policy_document": {
        "computed": true,
        "description": "The access control resource policy on the provided domain.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The resource schema to create a CodeArtifact domain.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodeartifactDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeartifactDomain), &result)
	return &result
}
