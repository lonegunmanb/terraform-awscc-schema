package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53HostedZone = `{
  "block": {
    "attributes": {
      "hosted_zone_config": {
        "computed": true,
        "description": "A complex type that contains an optional comment.\n\nIf you don't want to specify a comment, omit the HostedZoneConfig and Comment elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "Any comments that you want to include about the hosted zone.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "hosted_zone_tags": {
        "computed": true,
        "description": "Adds, edits, or deletes tags for a health check or a hosted zone.\n\nFor information about using tags for cost allocation, see Using Cost Allocation Tags in the AWS Billing and Cost Management User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the domain. Specify a fully qualified domain name, for example, www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the domain name is fully qualified. This means that Route 53 treats www.example.com (without a trailing dot) and www.example.com. (with a trailing dot) as identical.\n\nIf you're creating a public hosted zone, this is the name you have registered with your DNS registrar. If your domain name is registered with a registrar other than Route 53, change the name servers for your domain to the set of NameServers that are returned by the Fn::GetAtt intrinsic function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name_servers": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "query_logging_config": {
        "computed": true,
        "description": "A complex type that contains information about a configuration for DNS query logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_log_group_arn": {
              "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "vp_cs": {
        "computed": true,
        "description": "A complex type that contains information about the VPCs that are associated with the specified hosted zone.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_id": {
              "description": "The ID of an Amazon VPC.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_region": {
              "description": "The region that an Amazon VPC was created in. See https://docs.aws.amazon.com/general/latest/gr/rande.html for a list of up to date regions.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::Route53::HostedZone.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53HostedZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53HostedZone), &result)
	return &result
}
