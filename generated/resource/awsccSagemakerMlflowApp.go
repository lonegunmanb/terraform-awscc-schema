package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerMlflowApp = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the MLflow App.",
        "description_kind": "plain",
        "type": "string"
      },
      "artifact_store_uri": {
        "description": "The S3 URI for a general purpose bucket to use as the MLflow App artifact store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time that the MLflow App was created.",
        "description_kind": "plain",
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
        "description": "The date and time that the MLflow App was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "mlflow_app_id": {
        "computed": true,
        "description": "The server-generated identifier of the MLflow App.",
        "description_kind": "plain",
        "type": "string"
      },
      "mlflow_version": {
        "computed": true,
        "description": "The MLflow version used by the MLflow App.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_registration_mode": {
        "computed": true,
        "description": "Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the MLflow App.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Name (ARN) for an IAM role in your account that the MLflow App uses to access the artifact store in Amazon S3.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the MLflow App.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the MLflow App.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "weekly_maintenance_window_start": {
        "computed": true,
        "description": "The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled. For example: Tue:03:30.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::SageMaker::MlflowApp",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerMlflowAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerMlflowApp), &result)
	return &result
}
