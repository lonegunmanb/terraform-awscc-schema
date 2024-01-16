package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2WebAclAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "web_acl_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::WAFv2::WebACLAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWafv2WebAclAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2WebAclAssociation), &result)
	return &result
}
