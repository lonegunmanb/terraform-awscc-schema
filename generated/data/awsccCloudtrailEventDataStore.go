package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailEventDataStore = `{
  "block": {
    "attributes": {
      "advanced_event_selectors": {
        "computed": true,
        "description": "The advanced event selectors that were used to select events for the data store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field_selectors": {
              "computed": true,
              "description": "Contains all selector statements in an advanced event selector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ends_with": {
                    "computed": true,
                    "description": "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "equals": {
                    "computed": true,
                    "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "field": {
                    "computed": true,
                    "description": "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "not_ends_with": {
                    "computed": true,
                    "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "not_equals": {
                    "computed": true,
                    "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "not_starts_with": {
                    "computed": true,
                    "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "starts_with": {
                    "computed": true,
                    "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "set"
              }
            },
            "name": {
              "computed": true,
              "description": "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "billing_mode": {
        "computed": true,
        "description": "The mode that the event data store will use to charge for event storage.",
        "description_kind": "plain",
        "type": "string"
      },
      "context_key_selectors": {
        "computed": true,
        "description": "An array that enriches event records in an existing event data store by including additional information specified in individual ContexKeySelector entries. If you add ContextKeySelectors, you must set MaxEventSize to Large.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "equals": {
              "computed": true,
              "description": "An operator that includes events that match the exact value of the event record field specified in Type.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description": "Specifies the type of the event record field in ContextKeySelector. Valid values include RequestContext, TagContext.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "created_timestamp": {
        "computed": true,
        "description": "The timestamp of the event data store's creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_data_store_arn": {
        "computed": true,
        "description": "The ARN of the event data store.",
        "description_kind": "plain",
        "type": "string"
      },
      "federation_enabled": {
        "computed": true,
        "description": "Indicates whether federation is enabled on an event data store.",
        "description_kind": "plain",
        "type": "bool"
      },
      "federation_role_arn": {
        "computed": true,
        "description": "The ARN of the role used for event data store federation.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ingestion_enabled": {
        "computed": true,
        "description": "Indicates whether the event data store is ingesting events.",
        "description_kind": "plain",
        "type": "bool"
      },
      "insight_selectors": {
        "computed": true,
        "description": "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing event data store. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "insight_type": {
              "computed": true,
              "description": "The type of Insights to log on an event data store.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "insights_destination": {
        "computed": true,
        "description": "Specifies the ARN of the event data store that will collect Insights events. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_event_size": {
        "computed": true,
        "description": "Specifies the maximum size allowed for the event. Valid values are Standard and Large. If you add ContextKeySelectors, this value must be set to Large.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_region_enabled": {
        "computed": true,
        "description": "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The name of the event data store.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization_enabled": {
        "computed": true,
        "description": "Indicates that an event data store is collecting logged events for an organization.",
        "description_kind": "plain",
        "type": "bool"
      },
      "retention_period": {
        "computed": true,
        "description": "The retention period, in days.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "termination_protection_enabled": {
        "computed": true,
        "description": "Indicates whether the event data store is protected from termination.",
        "description_kind": "plain",
        "type": "bool"
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudTrail::EventDataStore",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudtrailEventDataStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailEventDataStore), &result)
	return &result
}
