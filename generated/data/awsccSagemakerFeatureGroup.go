package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerFeatureGroup = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "A timestamp of FeatureGroup creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description about the FeatureGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_time_feature_name": {
        "computed": true,
        "description": "The Event Time Feature Name.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_definitions": {
        "computed": true,
        "description": "An Array of Feature Definition",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "feature_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "feature_group_name": {
        "computed": true,
        "description": "The Name of the FeatureGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_group_status": {
        "computed": true,
        "description": "The status of the feature group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "offline_store_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_catalog_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "disable_glue_table_creation": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "s3_storage_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table_format": {
              "computed": true,
              "description": "Format for the offline store feature group. Iceberg is the optimal format for feature groups shared between offline and online stores.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "online_store_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_online_store": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "security_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "storage_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ttl_duration": {
              "computed": true,
              "description": "TTL configuration of the feature group",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "unit": {
                    "computed": true,
                    "description": "Unit of ttl configuration",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "Value of ttl configuration",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "record_identifier_feature_name": {
        "computed": true,
        "description": "The Record Identifier Feature Name.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pair to apply to this resource.",
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
          "nesting_mode": "list"
        }
      },
      "throughput_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provisioned_read_capacity_units": {
              "computed": true,
              "description": "For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.",
              "description_kind": "plain",
              "type": "number"
            },
            "provisioned_write_capacity_units": {
              "computed": true,
              "description": "For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.",
              "description_kind": "plain",
              "type": "number"
            },
            "throughput_mode": {
              "computed": true,
              "description": "Throughput mode configuration of the feature group",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::FeatureGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerFeatureGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerFeatureGroup), &result)
	return &result
}
