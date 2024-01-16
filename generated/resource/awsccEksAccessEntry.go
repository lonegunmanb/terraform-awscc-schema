package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksAccessEntry = `{
  "block": {
    "attributes": {
      "access_entry_arn": {
        "computed": true,
        "description": "The ARN of the access entry.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_policies": {
        "computed": true,
        "description": "An array of access policies that are associated with the access entry.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_scope": {
              "description": "The access scope of the access policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "namespaces": {
                    "computed": true,
                    "description": "The namespaces to associate with the access scope. Only specify if Type is set to 'namespace'.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "type": {
                    "description": "The type of the access scope.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "policy_arn": {
              "description": "The ARN of the access policy to add to the access entry.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "cluster_name": {
        "description": "The cluster that the access entry is created for.",
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
      "kubernetes_groups": {
        "computed": true,
        "description": "The Kubernetes groups that the access entry is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "principal_arn": {
        "description": "The principal ARN that the access entry is created for.",
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "The node type to associate with the access entry.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "username": {
        "computed": true,
        "description": "The Kubernetes user that the access entry is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "An object representing an Amazon EKS AccessEntry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEksAccessEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksAccessEntry), &result)
	return &result
}
