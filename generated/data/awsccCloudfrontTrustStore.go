package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontTrustStore = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the trust store",
        "description_kind": "plain",
        "type": "string"
      },
      "ca_certificates_bundle_source": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ca_certificates_bundle_s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description": "The S3 bucket containing the CA certificates bundle PEM file",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description": "The S3 object key of the CA certificates bundle PEM file",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "The S3 bucket region",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "The S3 object version of the CA certificates bundle PEM file",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "e_tag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The last modification timestamp of the trust store PEM file",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A unique name to identify the trust store",
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_ca_certificates": {
        "computed": true,
        "description": "The number of CA certificates in the trust store PEM file",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "Current status of the trust store",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs for resource tagging",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "trust_store_id": {
        "computed": true,
        "description": "The unique identifier for the trust store",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::TrustStore",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontTrustStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontTrustStore), &result)
	return &result
}
