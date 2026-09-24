package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPiPerfReports = `{
  "block": {
    "attributes": {
      "analysis_report_id": {
        "computed": true,
        "description": "A unique identifier for the analysis report.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the performance analysis report.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the analysis report was created in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "end_time": {
        "computed": true,
        "description": "The end time defined for the analysis report in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description": "An immutable, AWS Region-unique identifier for a data source. Performance Insights gathers metrics from this data source. To use an Amazon RDS instance as a data source, specify its DbiResourceId value.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_type": {
        "computed": true,
        "description": "The AWS service for which Performance Insights returns metrics. Valid values are RDS and DOCDB.",
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The start time defined for the analysis report in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the analysis report.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::PI::PerfReports",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPiPerfReportsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPiPerfReports), &result)
	return &result
}
