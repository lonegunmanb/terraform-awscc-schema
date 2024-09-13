package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveEventBridgeRuleTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "An eventbridge rule template's ARN (Amazon Resource Name)",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Placeholder documentation for __timestampIso8601",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A resource's optional description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_bridge_rule_template_id": {
        "computed": true,
        "description": "An eventbridge rule template's id. AWS provided templates have ids that start with ` + "`" + `aws-` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "event_targets": {
        "computed": true,
        "description": "Placeholder documentation for __listOfEventBridgeRuleTemplateTarget",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "description": "Target ARNs must be either an SNS topic or CloudWatch log group.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "event_type": {
        "description": "The type of event to match with the rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description": "An eventbridge rule template group's id. AWS provided template groups have ids that start with ` + "`" + `aws-` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "group_identifier": {
        "description": "An eventbridge rule template group's identifier. Can be either be its id or current name.",
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
      "identifier": {
        "computed": true,
        "description": "Placeholder documentation for __string",
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "Placeholder documentation for __timestampIso8601",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Represents the tags associated with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Definition of AWS::MediaLive::EventBridgeRuleTemplate Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveEventBridgeRuleTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveEventBridgeRuleTemplate), &result)
	return &result
}
