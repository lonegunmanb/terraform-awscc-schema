package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderInfrastructureConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the infrastructure configuration.",
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
      "instance_metadata_options": {
        "computed": true,
        "description": "The instance metadata option settings for the infrastructure configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "http_put_response_hop_limit": {
              "computed": true,
              "description": "Limit the number of hops that an instance metadata request can traverse to reach its destination.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "http_tokens": {
              "computed": true,
              "description": "Indicates whether a signed token header is required for instance metadata retrieval requests. The values affect the response as follows: ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "instance_profile_name": {
        "description": "The instance profile of the infrastructure configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_types": {
        "computed": true,
        "description": "The instance types of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "key_pair": {
        "computed": true,
        "description": "The EC2 key pair of the infrastructure configuration..",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging": {
        "computed": true,
        "description": "The logging configuration of the infrastructure configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_logs": {
              "computed": true,
              "description": "The S3 path in which to store the logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket_name": {
                    "computed": true,
                    "description": "S3BucketName",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_key_prefix": {
                    "computed": true,
                    "description": "S3KeyPrefix",
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
      "name": {
        "description": "The name of the infrastructure configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "The tags attached to the resource created by Image Builder.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "security_group_ids": {
        "computed": true,
        "description": "The security group IDs of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "sns_topic_arn": {
        "computed": true,
        "description": "The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The subnet ID of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the component.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "terminate_instance_on_failure": {
        "computed": true,
        "description": "The terminate instance on failure configuration of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::InfrastructureConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderInfrastructureConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderInfrastructureConfiguration), &result)
	return &result
}
