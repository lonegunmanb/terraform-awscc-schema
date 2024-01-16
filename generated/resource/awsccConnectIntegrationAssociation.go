package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectIntegrationAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_arn": {
        "description": "ARN of Integration being associated with the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_association_id": {
        "computed": true,
        "description": "Identifier of the association with Connect Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "description": "Specifies the integration type to be associated with the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::IntegrationAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectIntegrationAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectIntegrationAssociation), &result)
	return &result
}
