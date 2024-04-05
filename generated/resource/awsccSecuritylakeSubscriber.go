package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecuritylakeSubscriber = `{
  "block": {
    "attributes": {
      "access_types": {
        "description": "The Amazon S3 or AWS Lake Formation access type.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "data_lake_arn": {
        "description": "The ARN for the data lake.",
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
      "resource_share_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_share_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_bucket_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sources": {
        "description": "The supported AWS services from which logs and events are collected.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_log_source": {
              "computed": true,
              "description": "Amazon Security Lake supports log and event collection for natively supported AWS services.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_name": {
                    "computed": true,
                    "description": "The name for a AWS source. This must be a Regionally unique value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_version": {
                    "computed": true,
                    "description": "The version for a AWS source. This must be a Regionally unique value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "custom_log_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_name": {
                    "computed": true,
                    "description": "The name for a third-party custom source. This must be a Regionally unique value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_version": {
                    "computed": true,
                    "description": "The version for a third-party custom source. This must be a Regionally unique value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "subscriber_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_description": {
        "computed": true,
        "description": "The description for your subscriber account in Security Lake.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscriber_identity": {
        "description": "The AWS identity used to access your data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_id": {
              "description": "The external ID used to establish trust relationship with the AWS identity.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "principal": {
              "description": "The AWS identity principal.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "subscriber_name": {
        "description": "The name of your Security Lake subscriber account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subscriber_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The name of the tag. This is a general label that acts as a category for a more specific tag value (value).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value that is associated with the specified tag key (key). This value acts as a descriptor for the tag key. A tag value cannot be null, but it can be an empty string.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SecurityLake::Subscriber",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecuritylakeSubscriberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecuritylakeSubscriber), &result)
	return &result
}
