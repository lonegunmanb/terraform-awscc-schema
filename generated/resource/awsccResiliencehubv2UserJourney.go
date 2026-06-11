package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubv2UserJourney = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the user journey was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the user journey.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the user journey.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description": "The ARN of the resilience policy to associate with this user journey.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "system_identifier": {
        "description": "The system ARN or system ID that owns this user journey.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the user journey was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_journey_id": {
        "computed": true,
        "description": "The server-generated user journey ID.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a user journey within a Resilience Hub system.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubv2UserJourneySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2UserJourney), &result)
	return &result
}
