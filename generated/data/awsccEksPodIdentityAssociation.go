package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksPodIdentityAssociation = `{
  "block": {
    "attributes": {
      "association_arn": {
        "computed": true,
        "description": "The ARN of the pod identity association.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description": "The ID of the pod identity association.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The cluster that the pod identity association is created for.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The Kubernetes namespace that the pod identity association is created for.",
        "description_kind": "plain",
        "type": "string"
      },
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
    "description": "Data Source schema for AWS::EKS::PodIdentityAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEksPodIdentityAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksPodIdentityAssociation), &result)
	return &result
}
