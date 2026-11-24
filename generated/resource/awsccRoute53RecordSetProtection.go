package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecordSetProtection = `{
  "block": {
    "attributes": {
      "comment": {
        "computed": true,
        "description": "A comment that describes the protection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_zone_id": {
        "description": "The ID of the hosted zone that contains the record set to protect.",
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
      "record_name": {
        "description": "The name of the record set to protect.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "record_type": {
        "description": "The DNS record type of the record set to protect.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Route53::RecordSetProtection. Protects a record set (all record sets with the given name and type) in a hosted zone against changes until the protection is deleted. The protection must DependsOn (or reference with Ref) the AWS::Route53::RecordSet it protects. Adding or removing a protection and changing that record set must be two separate stack updates: a removal is applied in the cleanup phase after updates, and an addition in the same update blocks the record set's rollback. To recover a rollback blocked this way (UPDATE_ROLLBACK_FAILED), delete the protection, then run ContinueUpdateRollback. While a protection exists, creating any record set with its name and type, including one with a new SetIdentifier, is refused. Changing Comment replaces the protection by deleting it first, so the record set is unprotected for the duration of the replacement. While a protection exists, including one kept by DeletionPolicy Retain or created outside the stack, deleting its AWS::Route53::RecordSet fails with ResourceConflict; delete the protection first. When a create is retried after its outcome was lost, an existing protection with the same Comment is taken as the one it created; without a Comment, an existing protection that has no comment is adopted.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecordSetProtectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecordSetProtection), &result)
	return &result
}
