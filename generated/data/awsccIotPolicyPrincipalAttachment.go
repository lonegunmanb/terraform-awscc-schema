package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotPolicyPrincipalAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "The name of the AWS IoT policy",
        "description_kind": "plain",
        "type": "string"
      },
      "principal": {
        "computed": true,
        "description": "The principal, which can be a certificate ARN (as returned from the CreateCertificate operation) or an Amazon Cognito ID",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::PolicyPrincipalAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotPolicyPrincipalAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotPolicyPrincipalAttachment), &result)
	return &result
}
