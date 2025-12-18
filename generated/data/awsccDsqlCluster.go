package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDsqlCluster = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time of when the cluster was created in ISO-8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "Whether deletion protection is enabled in this cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption_details": {
        "computed": true,
        "description": "The encryption configuration details for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_status": {
              "computed": true,
              "description": "The status of encryption for the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption_type": {
              "computed": true,
              "description": "The type of encryption that protects data in the cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "endpoint": {
        "computed": true,
        "description": "The DSQL cluster endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description": "The ID of the created cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_encryption_key": {
        "computed": true,
        "description": "The KMS key that encrypts data on the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_region_properties": {
        "computed": true,
        "description": "The Multi-region properties associated to this cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "clusters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "witness_region": {
              "computed": true,
              "description": "The witness region in a multi-region cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "policy_document": {
        "computed": true,
        "description": "The IAM policy applied to the cluster resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_version": {
        "computed": true,
        "description": "The version number of the cluster's resource based policy",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_endpoint": {
        "computed": true,
        "description": "The DSQL cluster VPC endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint_service_name": {
        "computed": true,
        "description": "The VPC endpoint service name.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DSQL::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDsqlClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDsqlCluster), &result)
	return &result
}
