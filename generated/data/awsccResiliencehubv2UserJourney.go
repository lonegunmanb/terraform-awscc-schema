package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the user journey.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description": "The ARN of the resilience policy to associate with this user journey.",
        "description_kind": "plain",
        "type": "string"
      },
      "system_identifier": {
        "computed": true,
        "description": "The system ARN or system ID that owns this user journey.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::ResilienceHubV2::UserJourney",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResiliencehubv2UserJourneySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2UserJourney), &result)
	return &result
}
