package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccComprehendDocumentClassifierEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the document classifier endpoint.",
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
      "desired_inference_units": {
        "description": "The desired number of inference units to be used by the model. Each inference unit represents throughput of 100 characters per second.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "endpoint_name": {
        "description": "The name of the endpoint. The name must be unique within the AWS Region and account.",
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
      "last_modified_time": {
        "computed": true,
        "description": "The date and time that the endpoint was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_arn": {
        "description": "The Amazon Resource Name (ARN) of the document classifier model to which the endpoint is attached.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The second part of a key-value pair that forms a tag associated with a given resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "An Amazon Comprehend endpoint that hosts a custom document classification model for real-time inference.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccComprehendDocumentClassifierEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccComprehendDocumentClassifierEndpoint), &result)
	return &result
}
