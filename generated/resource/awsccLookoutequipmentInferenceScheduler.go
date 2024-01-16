package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutequipmentInferenceScheduler = `{
  "block": {
    "attributes": {
      "data_delay_offset_in_minutes": {
        "computed": true,
        "description": "A period of time (in minutes) by which inference on the data is delayed after the data starts.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "data_input_configuration": {
        "description": "Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inference_input_name_configuration": {
              "computed": true,
              "description": "Specifies configuration information for the input data for the inference, including timestamp format and delimiter.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "component_timestamp_delimiter": {
                    "computed": true,
                    "description": "Indicates the delimiter character used between items in the data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timestamp_format": {
                    "computed": true,
                    "description": "The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "input_time_zone_offset": {
              "computed": true,
              "description": "Indicates the difference between your time zone and Greenwich Mean Time (GMT).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_input_configuration": {
              "description": "Specifies configuration information for the input data for the inference, including input data S3 location.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "data_output_configuration": {
        "description": "Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The ID number for the AWS KMS key used to encrypt the inference output.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_output_configuration": {
              "description": "Specifies configuration information for the output results from the inference, including output S3 location.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "data_upload_frequency": {
        "description": "How often data is uploaded to the source S3 bucket for the input data.",
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
      "inference_scheduler_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the inference scheduler being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_scheduler_name": {
        "computed": true,
        "description": "The name of the inference scheduler being created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_name": {
        "description": "The name of the previously trained ML model being used to create the inference scheduler.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_side_kms_key_id": {
        "computed": true,
        "description": "Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags associated with the inference scheduler.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key for the specified tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the specified tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for LookoutEquipment InferenceScheduler.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLookoutequipmentInferenceSchedulerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutequipmentInferenceScheduler), &result)
	return &result
}
