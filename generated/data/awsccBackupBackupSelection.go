package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupBackupSelection = `{
  "block": {
    "attributes": {
      "backup_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_selection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "conditions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "string_equals": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "string_like": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "string_not_equals": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "string_not_like": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_value": {
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
            "iam_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "list_of_tags": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "condition_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "condition_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "not_resources": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "selection_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "selection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Backup::BackupSelection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupBackupSelectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupBackupSelection), &result)
	return &result
}
