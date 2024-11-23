package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRbinRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Rule Arn is unique for each rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the retention rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclude_resource_tags": {
        "computed": true,
        "description": "Information about the exclude resource tags used to identify resources that are excluded by the retention rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_tag_key": {
              "computed": true,
              "description": "The tag key of the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_tag_value": {
              "computed": true,
              "description": "The tag value of the resource",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description": "The unique ID of the retention rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "lock_configuration": {
        "computed": true,
        "description": "Information about the retention rule lock configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "unlock_delay_unit": {
              "computed": true,
              "description": "The unit of time in which to measure the unlock delay. Currently, the unlock delay can be measure only in days.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "unlock_delay_value": {
              "computed": true,
              "description": "The unlock delay period, measured in the unit specified for UnlockDelayUnit.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "lock_state": {
        "computed": true,
        "description": "The lock state for the retention rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "Information about the resource tags used to identify resources that are retained by the retention rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_tag_key": {
              "computed": true,
              "description": "The tag key of the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_tag_value": {
              "computed": true,
              "description": "The tag value of the resource",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "resource_type": {
        "description": "The resource type retained by the retention rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_period": {
        "description": "Information about the retention period for which the retention rule is to retain resources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "retention_period_unit": {
              "description": "The retention period unit of the rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retention_period_value": {
              "description": "The retention period value of the rule.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The state of the retention rule. Only retention rules that are in the available state retain resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Information about the tags assigned to the retention rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A unique identifier for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "String which you can use to describe or define the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Rbin::Rule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRbinRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRbinRule), &result)
	return &result
}
