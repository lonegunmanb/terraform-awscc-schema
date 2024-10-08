package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerMlflowTrackingServer = `{
  "block": {
    "attributes": {
      "artifact_store_uri": {
        "description": "The Amazon S3 URI for MLFlow Tracking Server artifacts.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "automatic_model_registration": {
        "computed": true,
        "description": "A flag to enable Automatic SageMaker Model Registration.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mlflow_version": {
        "computed": true,
        "description": "The MLFlow Version used on the MLFlow Tracking Server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tracking_server_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the MLFlow Tracking Server.",
        "description_kind": "plain",
        "type": "string"
      },
      "tracking_server_name": {
        "description": "The name of the MLFlow Tracking Server.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tracking_server_size": {
        "computed": true,
        "description": "The size of the MLFlow Tracking Server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "weekly_maintenance_window_start": {
        "computed": true,
        "description": "The start of the time window for maintenance of the MLFlow Tracking Server in UTC time.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::MlflowTrackingServer",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerMlflowTrackingServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerMlflowTrackingServer), &result)
	return &result
}
