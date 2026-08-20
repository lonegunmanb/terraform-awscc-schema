package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDrsSourceNetwork = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Source Network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_account_id": {
        "computed": true,
        "description": "The account ID containing the VPC to protect.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin_region": {
        "computed": true,
        "description": "The region containing the VPC to protect.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_network_id": {
        "computed": true,
        "description": "The ID of the Source Network.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags associated with the Source Network.",
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
      "vpc_id": {
        "computed": true,
        "description": "The VPC ID to protect.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DRS::SourceNetwork",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDrsSourceNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDrsSourceNetwork), &result)
	return &result
}
