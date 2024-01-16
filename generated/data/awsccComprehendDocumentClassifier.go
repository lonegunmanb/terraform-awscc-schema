package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComprehendDocumentClassifier = `{
  "block": {
    "attributes": {
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
      "document_classifier_name": {
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
      "input_data_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "augmented_manifests": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "split": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "data_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "document_reader_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "document_read_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "document_read_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "feature_types": {
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
            },
            "document_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "documents": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "test_s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "label_delimiter": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "test_s3_uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "language_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "model_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "model_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "output_data_config": {
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
      "version_name": {
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
    "description": "Data Source schema for AWS::Comprehend::DocumentClassifier",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccComprehendDocumentClassifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendDocumentClassifier), &result)
	return &result
}
