package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2TrustStoreRevocation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "revocation_contents": {
        "computed": true,
        "description": "The attributes required to create a trust store revocation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "revocation_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "revocation_id": {
        "computed": true,
        "description": "The ID associated with the revocation.",
        "description_kind": "plain",
        "type": "number"
      },
      "trust_store_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the trust store.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trust_store_revocations": {
        "computed": true,
        "description": "The data associated with a trust store revocation",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "number_of_revoked_entries": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "revocation_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "revocation_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "trust_store_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStoreRevocation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticloadbalancingv2TrustStoreRevocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2TrustStoreRevocation), &result)
	return &result
}
