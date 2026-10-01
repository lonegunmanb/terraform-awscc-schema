package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkflowmonitorScope = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_id": {
        "computed": true,
        "description": "The identifier for the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the scope.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "targets": {
        "description": "The targets for the scope.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region": {
              "description": "The AWS Region for the target resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_identifier": {
              "description": "A target identifier is a pair of identifying information for a scope target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_id": {
                    "description": "A target ID is an internally-generated identifier for a target.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_id": {
                          "description": "The account ID for the target.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "target_type": {
                    "description": "The type of the target. Currently always ACCOUNT.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::NetworkFlowMonitor::Scope. Creates a scope to define the resources that Network Flow Monitor monitors for network performance.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkflowmonitorScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorScope), &result)
	return &result
}
