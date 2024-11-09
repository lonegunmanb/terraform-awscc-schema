package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsCollaboration = `{
  "block": {
    "attributes": {
      "analytics_engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creator_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creator_member_abilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "creator_payment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "query_compute": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "is_responsible": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_encryption_metadata": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_cleartext": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "allow_duplicates": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "allow_joins_on_columns_with_different_names": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "preserve_nulls": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
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
      "members": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "display_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "member_abilities": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "payment_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_compute": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_responsible": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_log_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::CleanRooms::Collaboration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCleanroomsCollaborationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsCollaboration), &result)
	return &result
}
