package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupBackupPlan = `{
  "block": {
    "attributes": {
      "backup_plan": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "advanced_backup_settings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "backup_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "backup_plan_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "backup_plan_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "completion_window_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "copy_actions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination_backup_vault_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "lifecycle": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "delete_after_days": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "move_to_cold_storage_after_days": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "opt_in_to_archive_for_supported_resources": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "enable_continuous_backup": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "index_actions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "lifecycle": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_after_days": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "move_to_cold_storage_after_days": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "opt_in_to_archive_for_supported_resources": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "recovery_point_tags": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "rule_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "scan_actions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "malware_scanner": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "scan_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "schedule_expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "schedule_expression_timezone": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_window_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "target_backup_vault": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_logically_air_gapped_backup_vault_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "scan_settings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "malware_scanner": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_types": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "scanner_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "backup_plan_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_plan_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Backup::BackupPlan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupBackupPlan), &result)
	return &result
}
