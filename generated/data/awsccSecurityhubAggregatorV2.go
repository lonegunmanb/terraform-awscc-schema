package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubAggregatorV2 = `{
  "block": {
    "attributes": {
      "aggregation_region": {
        "computed": true,
        "description": "The aggregation Region of the AggregatorV2",
        "description_kind": "plain",
        "type": "string"
      },
      "aggregator_v2_arn": {
        "computed": true,
        "description": "The ARN of the AggregatorV2 being created and assigned as the unique identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "linked_regions": {
        "computed": true,
        "description": "The list of included Regions",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "region_linking_mode": {
        "computed": true,
        "description": "Indicates to link a list of included Regions",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with the Security Hub V2 resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::AggregatorV2",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubAggregatorV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubAggregatorV2), &result)
	return &result
}
