package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2TrustStore = `{
  "block": {
    "attributes": {
      "ca_certificates_bundle_s3_bucket": {
        "computed": true,
        "description": "The name of the S3 bucket to fetch the CA certificate bundle from.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_certificates_bundle_s3_key": {
        "computed": true,
        "description": "The name of the S3 object to fetch the CA certificate bundle from.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ca_certificates_bundle_s3_object_version": {
        "computed": true,
        "description": "The version of the S3 bucket that contains the CA certificate bundle.",
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
        "computed": true,
        "description": "The name of the trust store.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "number_of_ca_certificates": {
        "computed": true,
        "description": "The number of certificates associated with the trust store.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The status of the trust store, could be either of ACTIVE or CREATING.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the trust store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "trust_store_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the trust store.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStore",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticloadbalancingv2TrustStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2TrustStore), &result)
	return &result
}
