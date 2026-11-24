package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerEndpointPolicy = `{
  "block": {
    "attributes": {
      "endpoint_arn": {
        "description": "The Amazon Resource Name (ARN) of the endpoint the policy is attached to. The endpoint must be in the same account and Region as the stack.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint the policy is attached to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The resource policy document, as a JSON object. The service returns the document in its canonical form, so specify principals as full ARNs (for example arn:aws:iam::111122223333:root) and keep the form the service returns, to avoid drift.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::EndpointPolicy. Attaches a resource policy to a SageMaker endpoint. The policy grants other AWS accounts access to the endpoint, for example through a SageMaker Inference Router. An endpoint has at most one resource policy, which is removed when the endpoint is deleted.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerEndpointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerEndpointPolicy), &result)
	return &result
}
