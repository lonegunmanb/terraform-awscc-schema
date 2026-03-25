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
            "action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "add_or_run_anomaly_detection_for_analyses": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "analysis": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "approve_flow_share_requests": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "asana_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "automate": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "build_calculated_field_with_q": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "chat_agent": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_confluence_action": {
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
            "create_and_update_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_github_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_knowledge_bases": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_and_update_textract_action": {
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
            "create_and_update_zendesk_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_chat_agents": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "create_dashboard_executive_summary_with_q": {
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
            "dashboard": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "edit_visual_with_q": {
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
            "extension": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "flow": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "github_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "include_content_in_scheduled_reports_email": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "jira_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "linear_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "manage_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "monday_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "notion_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "perform_flow_ui_task": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "print_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "publish_without_approval": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rename_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "research": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_analyses": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_confluence_action": {
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
            "share_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_github_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_knowledge_bases": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_textract_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "share_zendesk_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "slack_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "space": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subscribe_dashboard_email_reports": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "textract_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "topic": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_agent_web_search": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_bedrock_models": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_github_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_textract_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "use_zendesk_action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "view_account_spice_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "zendesk_action": {
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
