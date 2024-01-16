package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComprehendFlywheel = `{
  "block": {
    "attributes": {
      "active_model_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_lake_s3_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_security_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_lake_kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "model_kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volume_kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "flywheel_name": {
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
      "model_type": {
        "computed": true,
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
      },
      "task_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "document_classification_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "labels": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "entity_recognition_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "entity_types": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "language_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Comprehend::Flywheel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccComprehendFlywheelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendFlywheel), &result)
	return &result
}
