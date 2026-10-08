package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotPolicyPrincipalAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the AWS IoT policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal": {
        "description": "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IoT::PolicyPrincipalAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotPolicyPrincipalAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotPolicyPrincipalAttachment), &result)
	return &result
}
