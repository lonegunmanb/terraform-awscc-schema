package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRamSourceAssociation = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The date and time when the association was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The date and time when the association was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_share_arn": {
        "computed": true,
        "description": "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_id": {
        "computed": true,
        "description": "Specifies the ID of the source account to associate with the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description": "The type of the source.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RAM::SourceAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRamSourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRamSourceAssociation), &result)
	return &result
}
