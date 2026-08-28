package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyCustomDetectionRuleAssociation = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID the association applies to.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description": "The service-generated unique identifier of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the association was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description": "Whether the rule runs in LIVE mode (generates findings) or DRY_RUN mode (evaluates without generating findings).",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description": "The catalog identifier of the custom detection rule to associate.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags applied to the association.",
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
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The time the association was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::CustomDetectionRuleAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyCustomDetectionRuleAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyCustomDetectionRuleAssociation), &result)
	return &result
}
