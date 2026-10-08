package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeCampaign = `{
  "block": {
    "attributes": {
      "campaign_arn": {
        "computed": true,
        "description": "The ARN of the campaign.",
        "description_kind": "plain",
        "type": "string"
      },
      "campaign_config": {
        "computed": true,
        "description": "The configuration details of a campaign.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_metadata_with_recommendations": {
              "computed": true,
              "description": "Whether metadata with recommendations is enabled for the campaign.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "item_exploration_config": {
              "computed": true,
              "description": "Specifies the exploration configuration hyperparameters.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "ranking_influence": {
              "computed": true,
              "description": "A map of ranking influence values for POPULARITY and FRESHNESS.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "number"
              ]
            },
            "sync_with_latest_solution_version": {
              "computed": true,
              "description": "Whether the campaign automatically updates to use the latest solution version.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "creation_date_time": {
        "computed": true,
        "description": "The time at which the campaign was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_date_time": {
        "computed": true,
        "description": "The time at which the campaign was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "min_provisioned_tps": {
        "computed": true,
        "description": "Specifies the requested minimum provisioned transactions per second.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "The name of the campaign.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "solution_version_arn": {
        "description": "The ARN of the solution version to deploy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the campaign.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the campaign.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A deployment of a solution version that provides real-time recommendations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeCampaignSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeCampaign), &result)
	return &result
}
