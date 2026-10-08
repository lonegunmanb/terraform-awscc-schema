package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerReportGenerator = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the report generator.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the report was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the report generator.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "report_context": {
        "computed": true,
        "description": "Details of the license configurations and asset groups that this generator reports on.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "license_asset_group_arns": {
              "computed": true,
              "description": "Amazon Resource Names (ARNs) of the license asset groups to include in the report.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "license_configuration_arns": {
              "computed": true,
              "description": "Amazon Resource Names (ARNs) of the license configurations that this generator reports on.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "report_end_date": {
              "computed": true,
              "description": "End date for the report data collection period.",
              "description_kind": "plain",
              "type": "string"
            },
            "report_start_date": {
              "computed": true,
              "description": "Start date for the report data collection period.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "report_creator_account": {
        "computed": true,
        "description": "The AWS account ID used to create the report generator.",
        "description_kind": "plain",
        "type": "string"
      },
      "report_frequency": {
        "computed": true,
        "description": "Details about how frequently reports are generated.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "period": {
              "computed": true,
              "description": "Time period between each report. The period can be daily, weekly, or monthly.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Number of times within the frequency period that a report is generated. The only supported value is 1.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "report_generator_name": {
        "computed": true,
        "description": "Name of the report generator.",
        "description_kind": "plain",
        "type": "string"
      },
      "report_type": {
        "computed": true,
        "description": "Type of reports to generate. The report type determines the data reported on.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "s3_location": {
        "computed": true,
        "description": "Details of the S3 bucket that report generator reports are published to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": "Name of the S3 bucket reports are published to.",
              "description_kind": "plain",
              "type": "string"
            },
            "key_prefix": {
              "computed": true,
              "description": "Prefix of the S3 bucket reports are published to.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::LicenseManager::ReportGenerator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerReportGeneratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerReportGenerator), &result)
	return &result
}
