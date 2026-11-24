package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalPortalPublication = `{
  "block": {
    "attributes": {
      "change_token": {
        "description": "Customer-supplied change-trigger string. Any change forces a republish on UpdateHandler. Customers may set this to a literal token, a CFN parameter, or a Fn::Sub-composed string (for example, a concatenation of dependent resources' UpdatedAt attributes) to chain republish through CFN's dependency graph.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the changes made in this published version. Forwarded to PublishPortal.description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_arn": {
        "description": "The ARN of the parent portal that this PortalPublication targets.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::PortalPublication. Drives the publish/unpublish lifecycle of an AWS::APIPortal::Portal. A separate resource keyed on PortalArn with a customer-controlled ChangeToken as the change-trigger. Customers may set ChangeToken to a literal value, a CFN parameter (manual republish), or a Fn::Sub-composed string referencing dependent resources' UpdatedAt attributes (auto-republish through CFN's dependency graph).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalPortalPublicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalPortalPublication), &result)
	return &result
}
