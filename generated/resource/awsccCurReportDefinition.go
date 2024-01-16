package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCurReportDefinition = `{
  "block": {
    "attributes": {
      "additional_artifacts": {
        "computed": true,
        "description": "A list of manifests that you want Amazon Web Services to create for this report.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "additional_schema_elements": {
        "computed": true,
        "description": "A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "billing_view_arn": {
        "computed": true,
        "description": "The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compression": {
        "description": "The compression format that AWS uses for the report.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "format": {
        "description": "The format that AWS saves the report in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "refresh_closed_reports": {
        "description": "Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "report_name": {
        "description": "The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "report_versioning": {
        "description": "Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket": {
        "description": "The S3 bucket where AWS delivers the report.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_prefix": {
        "description": "The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_region": {
        "description": "The region of the S3 bucket that AWS delivers the report into.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_unit": {
        "description": "The granularity of the line items in the report.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::CUR::ReportDefinition resource creates a Cost \u0026 Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCurReportDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCurReportDefinition), &result)
	return &result
}
