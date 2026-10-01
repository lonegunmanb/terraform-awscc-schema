package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComprehendEntityRecognizer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that identifies the entity recognizer.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "description": "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.",
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
        "description": "Specifies the format and location of the input data. The S3 bucket containing the input data must be located in the same Region as the entity recognizer being created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "annotations": {
              "computed": true,
              "description": "The S3 location of the CSV file that annotates your training documents.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_uri": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 location where the annotations are located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "test_s3_uri": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 location where the test annotations are located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "augmented_manifests": {
              "computed": true,
              "description": "A list of augmented manifest files that provide training data for a custom model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "annotation_data_s3_uri": {
                    "computed": true,
                    "description": "The S3 prefix to the annotation files that are referred in the augmented manifest file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "attribute_names": {
                    "computed": true,
                    "description": "The JSON attribute that contains the annotations for your training documents.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "document_type": {
                    "computed": true,
                    "description": "The type of augmented manifest.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
                    "computed": true,
                    "description": "The Amazon S3 location of the augmented manifest file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_documents_s3_uri": {
                    "computed": true,
                    "description": "The S3 prefix to the source files (PDFs) that are referred to in the augmented manifest file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "split": {
                    "computed": true,
                    "description": "The purpose of the data you've provided in the augmented manifest.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "data_format": {
              "computed": true,
              "description": "The format of your training data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "documents": {
              "computed": true,
              "description": "The S3 location of the folder that contains the training documents.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input_format": {
                    "computed": true,
                    "description": "Specifies how the text in an input file should be processed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 location where the training documents are located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "test_s3_uri": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 location where the test documents are located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "entity_list": {
              "computed": true,
              "description": "The S3 location of the CSV file that has the entity list.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_uri": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 location where the entity list is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "entity_types": {
              "description": "The entity types in the labeled training data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "description": "An entity type within a labeled training dataset.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "language_code": {
        "description": "The language of the input documents. All documents must be in the same language.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_kms_key_id": {
        "computed": true,
        "description": "ID for the AWS KMS key that Amazon Comprehend uses to encrypt trained custom models.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_policy": {
        "computed": true,
        "description": "The JSON resource-based policy to attach to your custom entity recognizer model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "recognizer_name": {
        "description": "The name given to the entity recognizer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the entity recognizer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the key-value pair that forms a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the key-value pair that forms a tag.",
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
        "description": "The version name given to the entity recognizer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_kms_key_id": {
        "computed": true,
        "description": "ID for the AWS KMS key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_config": {
        "computed": true,
        "description": "Configuration parameters for an optional private VPC containing the resources you are using for your custom entity recognizer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The ID number for a security group on an instance of your private VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description": "The ID for each subnet being used in your private VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "An Amazon Comprehend custom entity recognizer: a trained model that identifies custom entity types in text, created via an asynchronous training job.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccComprehendEntityRecognizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendEntityRecognizer), &result)
	return &result
}
