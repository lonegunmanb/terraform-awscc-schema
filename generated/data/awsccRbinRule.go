package data

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
              "type": "string"
            },
            "resource_tag_value": {
              "computed": true,
              "description": "The tag value of the resource",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "unlock_delay_value": {
              "computed": true,
              "description": "The unlock delay period, measured in the unit specified for UnlockDelayUnit.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "resource_tag_value": {
              "computed": true,
              "description": "The tag value of the resource",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type retained by the retention rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "retention_period": {
        "computed": true,
        "description": "Information about the retention period for which the retention rule is to retain resources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "retention_period_unit": {
              "computed": true,
              "description": "The retention period unit of the rule",
              "description_kind": "plain",
              "type": "string"
            },
            "retention_period_value": {
              "computed": true,
              "description": "The retention period value of the rule.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The state of the retention rule. Only retention rules that are in the available state retain resources.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "String which you can use to describe or define the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Rbin::Rule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRbinRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRbinRule), &result)
	return &result
}
