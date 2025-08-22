package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksAddon = `{
  "block": {
    "attributes": {
      "addon_name": {
        "computed": true,
        "description": "Name of Addon",
        "description_kind": "plain",
        "type": "string"
      },
      "addon_version": {
        "computed": true,
        "description": "Version of Addon",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the add-on",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "Name of Cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_values": {
        "computed": true,
        "description": "The configuration values to use with the add-on",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_config": {
        "computed": true,
        "description": "The custom namespace configuration to use with the add-on",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "namespace": {
              "computed": true,
              "description": "The custom namespace for creating the add-on",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "pod_identity_associations": {
        "computed": true,
        "description": "An array of pod identities to apply to this add-on.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "computed": true,
              "description": "The IAM role ARN that the pod identity association is created for.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_account": {
              "computed": true,
              "description": "The Kubernetes service account that the pod identity association is created for.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "preserve_on_delete": {
        "computed": true,
        "description": "PreserveOnDelete parameter value",
        "description_kind": "plain",
        "type": "bool"
      },
      "resolve_conflicts": {
        "computed": true,
        "description": "Resolve parameter value conflicts",
        "description_kind": "plain",
        "type": "string"
      },
      "service_account_role_arn": {
        "computed": true,
        "description": "IAM role to bind to the add-on's service account",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::EKS::Addon",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEksAddonSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksAddon), &result)
	return &result
}
