package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamExternalResourceVerificationToken = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_arn": {
        "computed": true,
        "description": "The ARN of the IPAM that created the token.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_external_resource_verification_token_arn": {
        "computed": true,
        "description": "The ARN of the IPAM external resource verification token.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_external_resource_verification_token_id": {
        "computed": true,
        "description": "The ID of the token.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_id": {
        "description": "The ID of the IPAM that will create the token.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_region": {
        "computed": true,
        "description": "The Region of the IPAM that created the token.",
        "description_kind": "plain",
        "type": "string"
      },
      "not_after": {
        "computed": true,
        "description": "The token expiration.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The token state.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The token status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the token.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "token_name": {
        "computed": true,
        "description": "The token name.",
        "description_kind": "plain",
        "type": "string"
      },
      "token_value": {
        "computed": true,
        "description": "The token value.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A verification token is an AWS-generated random value that you can use to prove ownership of an external resource. For example, you can use a verification token to validate that you control a public IP address range when you bring an IP address range to AWS (BYOIP).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamExternalResourceVerificationTokenSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamExternalResourceVerificationToken), &result)
	return &result
}
