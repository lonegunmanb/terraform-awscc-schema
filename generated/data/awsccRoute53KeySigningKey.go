package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53KeySigningKey = `{
  "block": {
    "attributes": {
      "hosted_zone_id": {
        "computed": true,
        "description": "The unique string (ID) used to identify a hosted zone.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_management_service_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53::KeySigningKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53KeySigningKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53KeySigningKey), &result)
	return &result
}
