package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsv2ResourcePolicy = `{
  "block": {
    "attributes": {
      "event_bus_arn": {
        "description": "The Amazon Resource Name (ARN) of the event bus whose resource policy this is. The bus must already exist. This resource does not create it.",
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
      "policy_document": {
        "description": "The resource policy document, as a JSON object. The document can be up to 20 KB. This quota is adjustable. An empty object is not a valid policy. To remove the policy, delete this resource. The principals in the document must exist and be visible to the service when the policy is written. When you create a new IAM role or user, that principal might not be immediately visible to the service. You might need to enforce a delay before you include it in the document. For more information, see \"Changes that I make are not always immediately visible\" in the IAM User Guide. Declare Version. Write AWS account and role principals as ARNs rather than as account IDs. Write a single Action, Resource, or principal value as a scalar rather than as a one-element list. The service returns these forms as you wrote them. It normalizes other forms, and a normalized value can appear as drift. A stack update replaces the whole policy with this document, including any change made outside CloudFormation. For more information about event bus resource policies, see the Amazon EventBridge User Guide.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The revision identifier the service assigned to the stored policy. The identifier changes on every successful write.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::EventsV2::ResourcePolicy, the resource policy of an Amazon EventBridge event bus. This manages only the resource policy named \"default\", which the bus owner writes. It does not manage the policy named \"AWS_RAM\", which AWS Resource Access Manager owns on behalf of the bus owner. If the bus already has a default policy, creating this resource fails. Deleting this resource removes all permissions granted by it. An explicit Deny in this policy takes precedence over an Allow in the \"AWS_RAM\" policy, so deleting this resource can widen access. Set DeletionPolicy: Retain if the policy carries a Deny that you rely on. Required permissions: events:PutResourcePolicy, events:GetResourcePolicy, and events:DeleteResourcePolicy. Listing resources of this type also requires events:ListEventBuses and events:ListResourcePolicies.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsv2ResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsv2ResourcePolicy), &result)
	return &result
}
