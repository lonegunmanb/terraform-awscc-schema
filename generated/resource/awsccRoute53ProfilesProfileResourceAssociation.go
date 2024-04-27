package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ProfilesProfileResourceAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of an association between the  Profile and resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_id": {
        "description": "The ID of the  profile that you associated the resource to that is specified by ResourceArn.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_resource_association_id": {
        "computed": true,
        "description": "Primary Identifier for  Profile Resource Association",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description": "The arn of the resource that you associated to the  Profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_properties": {
        "computed": true,
        "description": "A JSON-formatted string with key-value pairs specifying the properties of the associated resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The type of the resource associated to the  Profile.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Route53Profiles::ProfileResourceAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ProfilesProfileResourceAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ProfilesProfileResourceAssociation), &result)
	return &result
}
