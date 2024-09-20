package resource

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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "document_classifier_name": {
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
      "input_data_config": {
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
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "split": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "data_format": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "document_read_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "feature_types": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "document_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "test_s3_uri": {
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
            "label_delimiter": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_uri": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "test_s3_uri": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "language_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "s3_uri": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "version_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_kms_key_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Document Classifier enables training document classifier models.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccComprehendDocumentClassifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendDocumentClassifier), &result)
	return &result
}
