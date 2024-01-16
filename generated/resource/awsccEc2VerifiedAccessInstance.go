package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VerifiedAccessInstance = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Time this Verified Access Instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the AWS Verified Access instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fips_enabled": {
        "computed": true,
        "description": "Indicates whether FIPS is enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "Time this Verified Access Instance was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configurations": {
        "computed": true,
        "description": "The configuration options for AWS Verified Access instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs": {
              "computed": true,
              "description": "Sends Verified Access logs to CloudWatch Logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Indicates whether logging is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "log_group": {
                    "computed": true,
                    "description": "The ID of the CloudWatch Logs log group.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "include_trust_context": {
              "computed": true,
              "description": "Include claims from trust providers in Verified Access logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kinesis_data_firehose": {
              "computed": true,
              "description": "Sends Verified Access logs to Kinesis.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream": {
                    "computed": true,
                    "description": "The ID of the delivery stream.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Indicates whether logging is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "log_version": {
              "computed": true,
              "description": "Select log version for Verified Access logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3": {
              "computed": true,
              "description": "Sends Verified Access logs to Amazon S3.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description": "The bucket name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description": "The ID of the AWS account that owns the Amazon S3 bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Indicates whether logging is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The bucket prefix.",
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "verified_access_instance_id": {
        "computed": true,
        "description": "The ID of the AWS Verified Access instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_access_trust_provider_ids": {
        "computed": true,
        "description": "The IDs of the AWS Verified Access trust providers.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "verified_access_trust_providers": {
        "computed": true,
        "description": "AWS Verified Access trust providers.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of trust provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_trust_provider_type": {
              "computed": true,
              "description": "The type of device-based trust provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "trust_provider_type": {
              "computed": true,
              "description": "The type of trust provider (user- or device-based).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_trust_provider_type": {
              "computed": true,
              "description": "The type of user-based trust provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "verified_access_trust_provider_id": {
              "computed": true,
              "description": "The ID of the trust provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VerifiedAccessInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VerifiedAccessInstance), &result)
	return &result
}
