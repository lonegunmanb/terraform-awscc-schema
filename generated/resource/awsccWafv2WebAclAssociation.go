package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2WebAclAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "web_acl_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Associates WebACL to Application Load Balancer, CloudFront or API Gateway.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWafv2WebAclAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2WebAclAssociation), &result)
	return &result
}
