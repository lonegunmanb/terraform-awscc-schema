package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "region_linking_mode": {
        "description": "Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.\n The selected option also determines how to use the Regions provided in the Regions list.\n The options are as follows:\n  +   ` + "`" + `` + "`" + `ALL_REGIONS` + "`" + `` + "`" + ` - Aggregates findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. \n  +   ` + "`" + `` + "`" + `ALL_REGIONS_EXCEPT_SPECIFIED` + "`" + `` + "`" + ` - Aggregates findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ` + "`" + `` + "`" + `Regions` + "`" + `` + "`" + ` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. \n  +   ` + "`" + `` + "`" + `SPECIFIED_REGIONS` + "`" + `` + "`" + ` - Aggregates findings only from the Regions listed in the ` + "`" + `` + "`" + `Regions` + "`" + `` + "`" + ` parameter. Security Hub does not automatically aggregate findings from new Regions. \n  +   ` + "`" + `` + "`" + `NO_REGIONS` + "`" + `` + "`" + ` - Aggregates no data because no Regions are selected as linked Regions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "regions": {
        "computed": true,
        "description": "If ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `ALL_REGIONS_EXCEPT_SPECIFIED` + "`" + `` + "`" + `, then this is a space-separated list of Regions that don't replicate and send findings to the home Region.\n If ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `SPECIFIED_REGIONS` + "`" + `` + "`" + `, then this is a space-separated list of Regions that do replicate and send findings to the home Region. \n An ` + "`" + `` + "`" + `InvalidInputException` + "`" + `` + "`" + ` error results if you populate this field while ` + "`" + `` + "`" + `RegionLinkingMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `NO_REGIONS` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::SecurityHub::FindingAggregator` + "`" + `` + "`" + ` resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide* \n This resource must be created in the Region that you want to designate as your aggregation Region.\n Cross-Region aggregation is also a prerequisite for using [central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in ASH.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubFindingAggregatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubFindingAggregator), &result)
	return &result
}
