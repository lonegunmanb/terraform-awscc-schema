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
        "description_kind": "plain",
        "type": "string"
      },
      "finding_aggregator_arn": {
        "computed": true,
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
        "description": "Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.\n The selected option also determines how to use the Regions provided in the Regions list.\n The options are as follows:\n  +   ` + "`" + `` + "`" + `ALL_REGIONS` + "`" + `` + "`" + ` - Aggregates findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. \n  +   ` + "`" + `` + "`" + `ALL_REGIONS_EXCEPT_SPECIFIED` + "`" + `` + "`" + ` - Aggregates findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ` + "`" + `` + "`" + `Regions` + "`" + `` + "`" + ` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. \n  +   ` + "`" + `` + "`" + `SPECIFIED_REGIONS` + "`" + `` + "`" + ` - Aggregates findings only from the Regions listed in the ` + "`" + `` + "`" + `Regions` + "`" + `` + "`" + ` parameter. Security Hub does not automatically aggregate findings from new Regions. \n  +   ` + "`" + `` + "`" + `NO_REGIONS` + "`" + `` + "`" + ` - Aggregates no data because no Regions are selected as linked Regions.",
        "description_kind": "plain",
        "type": "string"
      },
      "regions": {
        "computed": true,
        "description": "If ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `ALL_REGIONS_EXCEPT_SPECIFIED` + "`" + `` + "`" + `, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.\n If ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `SPECIFIED_REGIONS` + "`" + `` + "`" + `, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region. \n An ` + "`" + `` + "`" + `InvalidInputException` + "`" + `` + "`" + ` error results if you populate this field while ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `NO_REGIONS` + "`" + `` + "`" + `.",
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
