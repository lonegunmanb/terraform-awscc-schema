package resource

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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "capabilities": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "add_or_run_anomaly_detection_for_analyses": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "analysis": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "automate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "chat_agent": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_dashboard_email_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_data_sources": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_datasets": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_themes": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_threshold_alerts": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_chat_agents": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_spice_dataset": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dashboard": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_csv": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_csv_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_excel": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_excel_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_pdf": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "export_to_pdf_in_scheduled_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "flow": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "include_content_in_scheduled_reports_email": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "perform_flow_ui_task": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "print_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_without_approval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rename_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "research": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_analyses": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_dashboards": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_data_sources": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_datasets": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "space": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subscribe_dashboard_email_reports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_agent_web_search": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_bedrock_models": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "view_account_spice_capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "custom_permissions_name": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of the AWS::QuickSight::CustomPermissions Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightCustomPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightCustomPermissions), &result)
	return &result
}
