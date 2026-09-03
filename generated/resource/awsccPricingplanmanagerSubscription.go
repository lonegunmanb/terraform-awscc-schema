package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "plan_family": {
        "description": "The name of the pricing plan family.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan_tier": {
        "description": "The tier of the pricing plan. Upgrades take effect immediately. However, rolling back an upgrade does not revert billing instantly; it schedules a downgrade to the end of the current billing period, and the higher-tier charge applies for the remainder of that month. While a downgrade is scheduled, the CurrentPlanTier property reports the tier currently being billed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arns": {
        "description": "The ARNs of resources associated with the subscription.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::PricingPlanManager::Subscription. Deleting an activated subscription does not terminate it immediately; it schedules a cancellation that takes effect at the end of the current billing period. Until that date the subscription remains active and billing continues. Deleting a subscription that has not yet been activated (PENDING_APPROVAL status) removes it immediately with no further charges.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPricingplanmanagerSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPricingplanmanagerSubscription), &result)
	return &result
}
