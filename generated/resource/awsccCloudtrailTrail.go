package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailTrail = `{
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
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "equals": {
                    "computed": true,
                    "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "field": {
                    "computed": true,
                    "description": "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "not_ends_with": {
                    "computed": true,
                    "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "not_equals": {
                    "computed": true,
                    "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "not_starts_with": {
                    "computed": true,
                    "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "starts_with": {
                    "computed": true,
                    "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "name": {
              "computed": true,
              "description": "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "aggregation_configurations": {
        "computed": true,
        "description": "Specifies the aggregation configuration to aggregate CloudTrail Events. A maximum of 1 aggregation configuration is allowed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_category": {
              "computed": true,
              "description": "The category of events to be aggregated.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "templates": {
              "computed": true,
              "description": "Contains all templates in an aggregation configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_logs_log_group_arn": {
        "computed": true,
        "description": "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cloudwatch_logs_role_arn": {
        "computed": true,
        "description": "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_log_file_validation": {
        "computed": true,
        "description": "Specifies whether log file validation is enabled. The default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "event_selectors": {
        "computed": true,
        "description": "Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_resources": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "exclude_management_event_sources": {
              "computed": true,
              "description": "An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing \"kms.amazonaws.com\". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "include_management_events": {
              "computed": true,
              "description": "Specify if you want your event selector to include management events for your trail.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "read_write_type": {
              "computed": true,
              "description": "Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "include_global_service_events": {
        "computed": true,
        "description": "Specifies whether the trail is publishing events from global services such as IAM to the log files.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "insight_selectors": {
        "computed": true,
        "description": "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_categories": {
              "computed": true,
              "description": "The categories of events for which to log insights. By default, insights are logged for management events only.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "insight_type": {
              "computed": true,
              "description": "The type of insight to log on a trail.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "is_logging": {
        "description": "Whether the CloudTrail is currently logging AWS API calls.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "is_multi_region_trail": {
        "computed": true,
        "description": "Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_organization_trail": {
        "computed": true,
        "description": "Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "kms_key_id": {
        "computed": true,
        "description": "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_bucket_name": {
        "description": "Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_key_prefix": {
        "computed": true,
        "description": "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sns_topic_name": {
        "computed": true,
        "description": "Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "trail_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudtrailTrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailTrail), &result)
	return &result
}
