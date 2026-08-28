package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mode": {
        "description": "Whether the rule runs in LIVE mode (generates findings) or DRY_RUN mode (evaluates without generating findings).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_id": {
        "description": "The catalog identifier of the custom detection rule to associate.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
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
      "updated_at": {
        "computed": true,
        "description": "The time the association was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::GuardDuty::CustomDetectionRuleAssociation. Associates a GuardDuty custom detection rule with the caller's account, enabling the rule in either LIVE or DRY_RUN mode.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGuarddutyCustomDetectionRuleAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyCustomDetectionRuleAssociation), &result)
	return &result
}
