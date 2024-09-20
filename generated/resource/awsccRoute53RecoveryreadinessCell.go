package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoveryreadinessCell = `{
  "block": {
    "attributes": {
      "cell_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cell.",
        "description_kind": "plain",
        "type": "string"
      },
      "cell_name": {
        "computed": true,
        "description": "The name of the cell to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cells": {
        "computed": true,
        "description": "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.",
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
      "parent_readiness_scopes": {
        "computed": true,
        "description": "The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
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
    "description": "The API Schema for AWS Route53 Recovery Readiness Cells.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoveryreadinessCellSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoveryreadinessCell), &result)
	return &result
}
