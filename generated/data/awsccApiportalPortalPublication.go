package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalPortalPublication = `{
  "block": {
    "attributes": {
      "change_token": {
        "computed": true,
        "description": "Customer-supplied change-trigger string. Any change forces a republish on UpdateHandler. Customers may set this to a literal token, a CFN parameter, or a Fn::Sub-composed string (for example, a concatenation of dependent resources' UpdatedAt attributes) to chain republish through CFN's dependency graph.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the changes made in this published version. Forwarded to PublishPortal.description.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description": "The ARN of the parent portal that this PortalPublication targets.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::PortalPublication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalPortalPublicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalPortalPublication), &result)
	return &result
}
