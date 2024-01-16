package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInsightsAccessScopeAnalysis = `{
  "block": {
    "attributes": {
      "analyzed_eni_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "findings_found": {
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
      "network_insights_access_scope_analysis_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_access_scope_analysis_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_access_scope_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      }
    },
    "description": "Data Source schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NetworkInsightsAccessScopeAnalysisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInsightsAccessScopeAnalysis), &result)
	return &result
}
