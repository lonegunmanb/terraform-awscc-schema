package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectIntegrationAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_arn": {
        "computed": true,
        "description": "ARN of Integration being associated with the instance",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_association_id": {
        "computed": true,
        "description": "Identifier of the association with Connect Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "computed": true,
        "description": "Specifies the integration type to be associated with the instance",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::IntegrationAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectIntegrationAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectIntegrationAssociation), &result)
	return &result
}
