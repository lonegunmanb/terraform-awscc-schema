package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccM2Deployment = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_version": {
        "computed": true,
        "description": "The version number of the application to deploy",
        "description_kind": "plain",
        "type": "number"
      },
      "deployment_id": {
        "computed": true,
        "description": "The deployment ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The environment ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the deployment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::M2::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccM2DeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccM2Deployment), &result)
	return &result
}
