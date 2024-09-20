package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoveryreadinessRecoveryGroup = `{
  "block": {
    "attributes": {
      "cells": {
        "computed": true,
        "description": "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "recovery_group_arn": {
        "computed": true,
        "description": "A collection of tags associated with a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "recovery_group_name": {
        "computed": true,
        "description": "The name of the recovery group to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource.",
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
      }
    },
    "description": "AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoveryreadinessRecoveryGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoveryreadinessRecoveryGroup), &result)
	return &result
}
