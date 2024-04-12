package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverOutpostResolver = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The OutpostResolver ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The OutpostResolver creation time",
        "description_kind": "plain",
        "type": "string"
      },
      "creator_request_id": {
        "computed": true,
        "description": "The id of the creator request.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_count": {
        "computed": true,
        "description": "The number of OutpostResolvers.",
        "description_kind": "plain",
        "type": "number"
      },
      "modification_time": {
        "computed": true,
        "description": "The OutpostResolver last modified time",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The OutpostResolver name.",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description": "The Outpost ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_resolver_id": {
        "computed": true,
        "description": "Id",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_instance_type": {
        "computed": true,
        "description": "The OutpostResolver instance type.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The OutpostResolver status, possible values are CREATING, OPERATIONAL, UPDATING, DELETING, ACTION_NEEDED, FAILED_CREATION and FAILED_DELETION.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "The OutpostResolver status message.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::OutpostResolver",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverOutpostResolverSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverOutpostResolver), &result)
	return &result
}
