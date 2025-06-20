package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "linked_regions": {
        "description": "The list of included Regions",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "region_linking_mode": {
        "description": "Indicates to link a list of included Regions",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with the Security Hub V2 resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "The AWS::SecurityHub::AggregatorV2 resource represents the AWS Security Hub AggregatorV2 in your account. One aggregatorv2 resource is created for each account in non opt-in region in which you configure region linking mode.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubAggregatorV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubAggregatorV2), &result)
	return &result
}
