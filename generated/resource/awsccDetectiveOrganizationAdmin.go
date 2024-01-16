package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDetectiveOrganizationAdmin = `{
  "block": {
    "attributes": {
      "account_id": {
        "description": "The account ID of the account that should be registered as your Organization's delegated administrator for Detective",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "graph_arn": {
        "computed": true,
        "description": "The Detective graph ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Detective::OrganizationAdmin",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDetectiveOrganizationAdminSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDetectiveOrganizationAdmin), &result)
	return &result
}
