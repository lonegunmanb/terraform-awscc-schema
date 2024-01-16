package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailBucket = `{
  "block": {
    "attributes": {
      "able_to_update_bundle": {
        "computed": true,
        "description": "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.",
        "description_kind": "plain",
        "type": "bool"
      },
      "access_rules": {
        "computed": true,
        "description": "An object that sets the public accessibility of objects in the specified bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_public_overrides": {
              "computed": true,
              "description": "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.",
              "description_kind": "plain",
              "type": "bool"
            },
            "get_object": {
              "computed": true,
              "description": "Specifies the anonymous access to all objects in a bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "bucket_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_name": {
        "computed": true,
        "description": "The name for the bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "bundle_id": {
        "computed": true,
        "description": "The ID of the bundle to use for the bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "object_versioning": {
        "computed": true,
        "description": "Specifies whether to enable or disable versioning of objects in the bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "read_only_access_accounts": {
        "computed": true,
        "description": "An array of strings to specify the AWS account IDs that can access the bucket.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "resources_receiving_access": {
        "computed": true,
        "description": "The names of the Lightsail resources for which to set bucket access.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
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
      },
      "url": {
        "computed": true,
        "description": "The URL of the bucket.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lightsail::Bucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailBucket), &result)
	return &result
}
