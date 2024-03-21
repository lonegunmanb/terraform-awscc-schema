package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmplifyDomain = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_sub_domain_creation_patterns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "auto_sub_domain_iam_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "certificate_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "certificate_verification_dns_record": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "certificate_record": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_auto_sub_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sub_domain_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branch_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "update_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Amplify::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAmplifyDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmplifyDomain), &result)
	return &result
}
