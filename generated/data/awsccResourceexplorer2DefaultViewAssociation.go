package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResourceexplorer2DefaultViewAssociation = `{
  "block": {
    "attributes": {
      "associated_aws_principal": {
        "computed": true,
        "description": "The AWS principal that the default view is associated with, used as the unique identifier for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "view_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResourceExplorer2::DefaultViewAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResourceexplorer2DefaultViewAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResourceexplorer2DefaultViewAssociation), &result)
	return &result
}
