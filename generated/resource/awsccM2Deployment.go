package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccM2Deployment = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "The application ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_version": {
        "description": "The version number of the application to deploy",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "deployment_id": {
        "computed": true,
        "description": "The deployment ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "description": "The environment ID.",
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
      "status": {
        "computed": true,
        "description": "The status of the deployment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a deployment resource of an AWS Mainframe Modernization (M2) application to a specified environment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccM2DeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccM2Deployment), &result)
	return &result
}
