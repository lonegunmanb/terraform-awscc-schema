package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConformancePack = `{
  "block": {
    "attributes": {
      "conformance_pack_input_parameters": {
        "computed": true,
        "description": "A list of ConformancePackInputParameter objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_name": {
              "computed": true,
              "description": "Key part of key-value pair with value being parameter value",
              "description_kind": "plain",
              "type": "string"
            },
            "parameter_value": {
              "computed": true,
              "description": "Value part of key-value pair with key being parameter Name",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "conformance_pack_name": {
        "computed": true,
        "description": "Name of the conformance pack which will be assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_s3_bucket": {
        "computed": true,
        "description": "AWS Config stores intermediate files while processing conformance pack template.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_s3_key_prefix": {
        "computed": true,
        "description": "The prefix for delivery S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_body": {
        "computed": true,
        "description": "A string containing full conformance pack template body. You can only specify one of the template body or template S3Uri fields.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_s3_uri": {
        "computed": true,
        "description": "Location of file containing the template body which points to the conformance pack template that is located in an Amazon S3 bucket. You can only specify one of the template body or template S3Uri fields.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_ssm_document_details": {
        "computed": true,
        "description": "The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "document_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "document_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Config::ConformancePack",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigConformancePackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConformancePack), &result)
	return &result
}
