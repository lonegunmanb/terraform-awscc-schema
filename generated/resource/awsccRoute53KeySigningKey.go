package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53KeySigningKey = `{
  "block": {
    "attributes": {
      "hosted_zone_id": {
        "description": "The unique string (ID) used to identify a hosted zone.",
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
      "key_management_service_arn": {
        "description": "The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "description": "A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53KeySigningKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53KeySigningKey), &result)
	return &result
}
