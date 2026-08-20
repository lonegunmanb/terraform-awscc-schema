package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin_account_id": {
        "description": "The account ID containing the VPC to protect.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_region": {
        "description": "The region containing the VPC to protect.",
        "description_kind": "plain",
        "required": true,
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
      "vpc_id": {
        "description": "The VPC ID to protect.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A Source Network resource represents a VPC that is protected by AWS Elastic Disaster Recovery.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDrsSourceNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDrsSourceNetwork), &result)
	return &result
}
