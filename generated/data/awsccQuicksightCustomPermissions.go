package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightCustomPermissions = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capabilities": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_or_run_anomaly_detection_for_analyses": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_dashboard_email_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_data_sources": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_datasets": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_themes": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_threshold_alerts": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_spice_dataset": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_csv": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_csv_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_excel": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_excel_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_pdf": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_to_pdf_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "include_content_in_scheduled_reports_email": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "print_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rename_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_analyses": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_dashboards": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_data_sources": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_datasets": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subscribe_dashboard_email_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "view_account_spice_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "custom_permissions_name": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::QuickSight::CustomPermissions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightCustomPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightCustomPermissions), &result)
	return &result
}
