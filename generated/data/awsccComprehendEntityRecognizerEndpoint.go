package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComprehendEntityRecognizerEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the entity recognizer endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The creation date and time of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_inference_units": {
        "computed": true,
        "description": "The number of inference units currently used by the model using this endpoint.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_access_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to trained custom models encrypted with a customer managed key (ModelKmsKeyId).",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_inference_units": {
        "computed": true,
        "description": "The desired number of inference units to be used by the model. Each inference unit represents throughput of 100 characters per second.",
        "description_kind": "plain",
        "type": "number"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint. The name must be unique within the AWS Region and account.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_status": {
        "computed": true,
        "description": "The current status of the endpoint. Because the endpoint updates and creation are asynchronous, wait for the endpoint to be IN_SERVICE before making inference requests.",
        "description_kind": "plain",
        "type": "string"
      },
      "flywheel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the flywheel to which the endpoint is attached.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date and time that the endpoint was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the entity recognizer model to which the endpoint is attached.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the endpoint being created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The initial part of a key-value pair that forms a tag associated with a given resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The second part of a key-value pair that forms a tag associated with a given resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Comprehend::EntityRecognizerEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccComprehendEntityRecognizerEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendEntityRecognizerEndpoint), &result)
	return &result
}
