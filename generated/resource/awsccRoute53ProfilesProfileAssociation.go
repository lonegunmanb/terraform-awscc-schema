package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ProfilesProfileAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the  profile association.",
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
      "name": {
        "description": "The name of an association between a  Profile and a VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_association_id": {
        "computed": true,
        "description": "Primary Identifier for  Profile Association",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "description": "The ID of the  profile that you associated with the resource that is specified by ResourceId.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "description": "The resource that you associated the  profile with.",
        "description_kind": "plain",
        "required": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Route53Profiles::ProfileAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ProfilesProfileAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ProfilesProfileAssociation), &result)
	return &result
}
