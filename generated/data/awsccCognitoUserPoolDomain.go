package data

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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain": {
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
      "managed_login_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "routing": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary_route_53_health_check_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secondary_region": {
                    "computed": true,
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
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolDomain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolDomain), &result)
	return &result
}
