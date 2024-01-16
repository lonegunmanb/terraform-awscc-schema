package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmplifyDomain = `{
  "block": {
    "attributes": {
      "app_id": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "auto_sub_domain_iam_role": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_record": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sub_domain_settings": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branch_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAmplifyDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmplifyDomain), &result)
	return &result
}
