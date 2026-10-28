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
            "access_apps_native_data_store": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
            "adobe_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "amazon_s_three_action": {
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
            "approve_flow_share_requests": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "apps": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "asana_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "athena_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aurora_data_source": {
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
            "bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bedrock_managed_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "box_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "build_calculated_field_with_q": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "canva_agent_action": {
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
            "comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "confluence_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_adobe_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_apps": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_bedrock_managed_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_box_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_confluence_knowledge_base": {
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
            "create_and_update_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_github_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_google_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_idc_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_knowledge_bases": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_one_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_q_business_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_s3_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_share_point_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_textract_action": {
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
            "create_and_update_web_crawler_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_and_update_zendesk_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_athena_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_aurora_data_source": {
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
            "create_dashboard_executive_summary_with_q": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_databricks_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_db_2_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_denodo_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_document_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_dremio_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_dynamo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_exasol_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_file_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_git_hub_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_google_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_google_big_query_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_google_sheets_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_impala_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_jira_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_maria_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_mongo_atlas_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_mongo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_my_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_open_search_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_oracle_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_pay_pal_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_postgre_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_presto_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_radiant_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_rds_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_redshift_auto_discovered_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_redshift_manual_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_s3_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_s3_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_s3_tables_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_salesforce_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_sap_hana_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_service_now_data_source": {
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
            "create_snowflake_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_spaces": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_spark_data_source": {
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
            "create_sql_server_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_square_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_starburst_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_teradata_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_timestream_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_trino_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "create_twitter_data_source": {
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
            "databricks_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "db_2_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "denodo_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "document_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dremio_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dynamo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "edit_visual_with_q": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exasol_data_source": {
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
            "extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file_data_source": {
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
            "generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "git_hub_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "github_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "google_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "google_big_query_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "google_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "google_sheets_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "idc_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "impala_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inbound_email_trigger": {
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
            "intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "invoke_apps_ai_inference": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "jira_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "jira_data_source": {
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
            "linear_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "manage_shared_folders": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maria_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monday_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mongo_atlas_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mongo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "my_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "notion_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "one_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "open_search_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "oracle_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pay_pal_data_source": {
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
            "postgre_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "presto_data_source": {
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
            "q_business_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "quick_event_trigger": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "radiant_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rds_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redshift_auto_discovered_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redshift_manual_data_source": {
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
            "s3_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_tables_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "salesforce_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_hana_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_trigger": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_now_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_adobe_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_amazon_s_three_action": {
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
            "share_apps": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_athena_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_aurora_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_bedrock_managed_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_box_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_chat_agents": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_confluence_knowledge_base": {
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
            "share_databricks_data_source": {
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
            "share_db_2_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_denodo_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_document_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_dremio_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_dynamo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_exasol_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_file_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_git_hub_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_github_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_google_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_google_big_query_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_google_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_google_sheets_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_idc_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_impala_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_jira_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_knowledge_bases": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_maria_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_mongo_atlas_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_mongo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_my_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_one_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_open_search_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_oracle_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_pay_pal_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_point_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_postgre_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_presto_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_q_business_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_radiant_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_rds_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_redshift_auto_discovered_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_redshift_manual_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_s3_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_s3_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_s3_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_s3_tables_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_salesforce_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_hana_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_service_now_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_share_point_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_snowflake_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_spaces": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_spark_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_sql_server_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_square_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_starburst_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_teradata_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_textract_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_timestream_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_trino_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_twitter_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_web_crawler_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "share_zendesk_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "slack_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "snowflake_data_source": {
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
            "spark_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_server_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "square_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "starburst_data_source": {
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
            "teradata_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "textract_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timestream_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trigger": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trino_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "twitter_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_adobe_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_athena_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_aurora_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_databricks_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_db_2_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_denodo_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_document_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_dremio_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_dynamo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_exasol_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_file_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_git_hub_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_google_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_google_big_query_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_google_sheets_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_impala_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_jira_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_maria_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_mongo_atlas_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_mongo_db_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_my_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_open_search_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_oracle_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_pay_pal_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_postgre_sql_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_presto_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_radiant_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_rds_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_redshift_auto_discovered_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_redshift_manual_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_s3_analytics_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_s3_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_s3_tables_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_salesforce_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_sap_hana_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_service_now_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_snowflake_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_spark_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_sql_server_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_square_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_starburst_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_teradata_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_timestream_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_trino_data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_twitter_data_source": {
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
            "use_amazon_bedrock_ars_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_amazon_bedrock_fs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_amazon_bedrock_krs_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_amazon_s_three_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_asana_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_bamboo_hr_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_bedrock_managed_knowledge_base": {
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
            "use_box_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_box_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_browser_extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_canva_agent_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_comprehend_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_comprehend_medical_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_confluence_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_confluence_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_excel_add_in_extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_fact_set_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_generic_http_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_github_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_google_calendar_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_google_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_hubspot_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_hugging_face_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_idc_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_intercom_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_jira_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_linear_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_mcp_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_monday_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_ms_exchange_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_ms_teams_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_new_relic_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_notion_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_one_drive_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_one_drive_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_open_api_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_outlook_add_in_extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_pager_duty_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_powerpoint_add_in_extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_q_business_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_s3_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_salesforce_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sand_p_global_energy_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sand_pgmi_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sap_bill_of_material_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sap_business_partner_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sap_material_stock_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sap_physical_inventory_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_sap_product_master_data_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_service_now_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_share_point_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_share_point_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_slack_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_smartsheet_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_textract_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_web_crawler_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_word_add_in_extension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_zendesk_action": {
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
            },
            "web_crawler_knowledge_base": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zendesk_action": {
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
