package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkflowmonitorScope = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "targets": {
        "computed": true,
        "description": "The targets for the scope.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region": {
              "computed": true,
              "description": "The AWS Region for the target resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_identifier": {
              "computed": true,
              "description": "A target identifier is a pair of identifying information for a scope target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_id": {
                    "computed": true,
                    "description": "A target ID is an internally-generated identifier for a target.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_id": {
                          "computed": true,
                          "description": "The account ID for the target.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "target_type": {
                    "computed": true,
                    "description": "The type of the target. Currently always ACCOUNT.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::NetworkFlowMonitor::Scope",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkflowmonitorScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkflowmonitorScope), &result)
	return &result
}
