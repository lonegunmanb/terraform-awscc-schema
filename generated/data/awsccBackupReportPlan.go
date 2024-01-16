package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupReportPlan = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "report_delivery_channel": {
        "computed": true,
        "description": "A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "formats": {
              "computed": true,
              "description": "A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "s3_bucket_name": {
              "computed": true,
              "description": "The unique name of the S3 bucket that receives your reports.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key_prefix": {
              "computed": true,
              "description": "The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "report_plan_arn": {
        "computed": true,
        "description": "An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.",
        "description_kind": "plain",
        "type": "string"
      },
      "report_plan_description": {
        "computed": true,
        "description": "An optional description of the report plan with a maximum of 1,024 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "report_plan_name": {
        "computed": true,
        "description": "The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).",
        "description_kind": "plain",
        "type": "string"
      },
      "report_plan_tags": {
        "computed": true,
        "description": "Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "report_setting": {
        "computed": true,
        "description": "Identifies the report template for the report. Reports are built using a report template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "accounts": {
              "computed": true,
              "description": "The list of AWS accounts that a report covers.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "framework_arns": {
              "computed": true,
              "description": "The Amazon Resource Names (ARNs) of the frameworks a report covers.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "organization_units": {
              "computed": true,
              "description": "The list of AWS organization units that a report covers.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "regions": {
              "computed": true,
              "description": "The list of AWS regions that a report covers.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "report_template": {
              "computed": true,
              "description": "Identifies the report template for the report. Reports are built using a report template. The report templates are: ` + "`" + `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT` + "`" + `",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Backup::ReportPlan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupReportPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupReportPlan), &result)
	return &result
}
