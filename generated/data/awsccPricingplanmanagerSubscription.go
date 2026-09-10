package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPricingplanmanagerSubscription = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the subscription was created, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_plan_tier": {
        "computed": true,
        "description": "The plan tier currently active on the subscription as reported by the API. Populated by the Read handler. Diverges from PlanTier after a non-reversible CloudFormation rollback; surface via drift detection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan_family": {
        "computed": true,
        "description": "The name of the pricing plan family.",
        "description_kind": "plain",
        "type": "string"
      },
      "plan_tier": {
        "computed": true,
        "description": "The tier of the pricing plan. CloudFormation does not change the tier of an existing subscription; a stack update that changes the tier, upgrading or downgrading it, is rejected.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arns": {
        "computed": true,
        "description": "The ARNs of resources associated with the subscription.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The status of the subscription. PENDING_APPROVAL means a paid-tier subscription has been created but is not yet active and incurs no charges until it is approved out of band via a separate ApprovePaidSubscription call; CloudFormation never approves it. Free-tier subscriptions are activated immediately and do not use this status. ACTIVE means the subscription is in effect and, for paid tiers, billing has started. SYNC_IN_PROGRESS means a change is being applied. FAILED means provisioning did not complete; see StatusReason.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "A human-readable explanation of why the subscription is in its current status. Populated only when Status is FAILED, where it carries the reason the subscription could not be provisioned. Empty for all other statuses.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The time the subscription was last modified, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "usage_level": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::PricingPlanManager::Subscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPricingplanmanagerSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPricingplanmanagerSubscription), &result)
	return &result
}
