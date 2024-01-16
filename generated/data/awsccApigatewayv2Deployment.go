package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Deployment = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for the deployment resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_name": {
        "computed": true,
        "description": "The name of an existing stage to associate with the deployment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2DeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Deployment), &result)
	return &result
}
