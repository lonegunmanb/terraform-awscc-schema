package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TrafficMirrorFilter = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of a traffic mirror filter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_services": {
        "computed": true,
        "description": "The network service that is associated with the traffic mirror filter.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags for a traffic mirror filter.",
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
      "traffic_mirror_filter_id": {
        "computed": true,
        "description": "The ID of a traffic mirror filter.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::EC2::TrafficMirrorFilter",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TrafficMirrorFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TrafficMirrorFilter), &result)
	return &result
}
