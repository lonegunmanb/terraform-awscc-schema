package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ProfilesProfileResourceAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of an association between the  Profile and resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "The ID of the  profile that you associated the resource to that is specified by ResourceArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_resource_association_id": {
        "computed": true,
        "description": "Primary Identifier for  Profile Resource Association",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The arn of the resource that you associated to the  Profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_properties": {
        "computed": true,
        "description": "A JSON-formatted string with key-value pairs specifying the properties of the associated resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The type of the resource associated to the  Profile.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53Profiles::ProfileResourceAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ProfilesProfileResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ProfilesProfileResourceAssociation), &result)
	return &result
}
