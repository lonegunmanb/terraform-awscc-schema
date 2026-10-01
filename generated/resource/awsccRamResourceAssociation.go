package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRamResourceAssociation = `{
  "block": {
    "attributes": {
      "association_type": {
        "computed": true,
        "description": "The type of entity included in this association.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time when the association was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "external": {
        "computed": true,
        "description": "Indicates whether the principal belongs to the same organization in AWS Organizations as the AWS account that owns the resource share.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The date and time when the association was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description": "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to associate with the resource share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_share_arn": {
        "description": "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates a specified resource with a resource share.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRamResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRamResourceAssociation), &result)
	return &result
}
