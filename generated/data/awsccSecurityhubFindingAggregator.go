package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubFindingAggregator = `{
  "block": {
    "attributes": {
      "finding_aggregation_region": {
        "computed": true,
        "description": "The aggregation Region of the FindingAggregator",
        "description_kind": "plain",
        "type": "string"
      },
      "finding_aggregator_arn": {
        "computed": true,
        "description": "The ARN of the FindingAggregator being created and assigned as the unique identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region_linking_mode": {
        "computed": true,
        "description": "Indicates whether to link all Regions, all Regions except for a list of excluded Regions, or a list of included Regions",
        "description_kind": "plain",
        "type": "string"
      },
      "regions": {
        "computed": true,
        "description": "The list of excluded Regions or included Regions",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::FindingAggregator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubFindingAggregatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubFindingAggregator), &result)
	return &result
}
