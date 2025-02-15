package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolDomain = `{
  "block": {
    "attributes": {
      "cloudfront_distribution": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "domain": {
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
      "managed_login_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolDomain",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolDomain), &result)
	return &result
}
