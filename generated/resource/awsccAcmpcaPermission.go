package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAcmpcaPermission = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "certificate_authority_arn": {
        "description": "The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "description": "The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_account": {
        "computed": true,
        "description": "The ID of the calling account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Permission set on private certificate authority",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAcmpcaPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAcmpcaPermission), &result)
	return &result
}
