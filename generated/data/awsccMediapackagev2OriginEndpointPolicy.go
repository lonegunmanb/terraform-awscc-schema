package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2OriginEndpointPolicy = `{
  "block": {
    "attributes": {
      "cdn_auth_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe settings to enable CDN authorization headers in MediaPackage.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cdn_identifier_secret_arns": {
              "computed": true,
              "description": "\u003cp\u003eThe ARN for the secret in Secrets Manager that your CDN uses for authorization to access the endpoint.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "secrets_role_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe ARN for the IAM role that gives MediaPackage read access to Secrets Manager and KMS for CDN authorization.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "channel_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
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
      "origin_endpoint_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaPackageV2::OriginEndpointPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagev2OriginEndpointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2OriginEndpointPolicy), &result)
	return &result
}
