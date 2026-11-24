package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53HealthCheckSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_source_config": {
        "computed": true,
        "description": "\u003cp\u003eThe VPC, subnets, and security groups for the health check source.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "subnet_configuration_list": {
              "computed": true,
              "description": "\u003cp\u003eA complex type that contains one \u003ccode\u003eSubnetConfiguration\u003c/code\u003e element for each Availability Zone that Route 53 uses for private endpoint health checks in the Amazon Web Services Region. Each element must be in a different Availability Zone. Most Regions require three elements. Regions that have only two such Availability Zones, such as US West (N. California), require two elements. Route 53 creates network interfaces in each subnet and sends health check requests from all of them.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_id": {
                    "computed": true,
                    "description": "\u003cp\u003eThe ID of the Availability Zone that contains the subnet, for example \u003ccode\u003euse1-az1\u003c/code\u003e. Specify the Availability Zone ID, not the Availability Zone name. Each subnet in a health check source must be in a different Availability Zone.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_group_id_list": {
                    "computed": true,
                    "description": "\u003cp\u003eThe IDs of one to ten security groups in the same VPC to associate with the network interfaces that Route 53 creates in this subnet. The security groups must allow outbound traffic to the endpoints that the health checks monitor.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "\u003cp\u003eThe ID of the subnet. The subnet must have free IP addresses for the network interfaces that Route 53 creates.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "vpc_id": {
              "computed": true,
              "description": "\u003cp\u003eThe ID of the VPC that contains the subnets for the health check source. All subnets in \u003ccode\u003eSubnetConfigurationList\u003c/code\u003e must belong to this VPC.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "health_check_source_id": {
        "computed": true,
        "description": "\u003cp\u003eThe ID of the health check source. Route 53 assigns this ID when it creates the source.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_source_status": {
        "computed": true,
        "description": "\u003cp\u003eThe current status of the health check source. The health check source status can have one of the following values:\u003c/p\u003e\n         \u003cdl\u003e\n            \u003cdt\u003eCREATE_PENDING, CREATE_IN_PROGRESS\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eRoute 53 is creating the source. Health checks that reference the source don't send requests yet.\u003c/p\u003e\n            \u003c/dd\u003e\n            \u003cdt\u003eCREATE_COMPLETED\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eThe source is ready to use.\u003c/p\u003e\n            \u003c/dd\u003e\n            \u003cdt\u003eCREATE_FAILED\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eRoute 53 couldn't create the source. \u003ccode\u003eStatusReasonCode\u003c/code\u003e and \u003ccode\u003eStatusReason\u003c/code\u003e explain why.\u003c/p\u003e\n            \u003c/dd\u003e\n            \u003cdt\u003eDELETE_PENDING, DELETE_IN_PROGRESS\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eRoute 53 is deleting the source.\u003c/p\u003e\n            \u003c/dd\u003e\n            \u003cdt\u003eDELETE_COMPLETED\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eRoute 53 has deleted the source.\u003c/p\u003e\n            \u003c/dd\u003e\n            \u003cdt\u003eDELETE_FAILED\u003c/dt\u003e\n            \u003cdd\u003e\n               \u003cp\u003eRoute 53 couldn't delete the source. \u003ccode\u003eStatusReasonCode\u003c/code\u003e and \u003ccode\u003eStatusReason\u003c/code\u003e explain why.\u003c/p\u003e\n            \u003c/dd\u003e\n         \u003c/dl\u003e\n         \u003cp\u003eIf the status is \u003ccode\u003eCREATE_FAILED\u003c/code\u003e, delete the source and create a new one. If the status is \u003ccode\u003eDELETE_FAILED\u003c/code\u003e, retry the \u003ccode\u003eDeleteHealthCheckSource\u003c/code\u003e request.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Route53::HealthCheckSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53HealthCheckSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53HealthCheckSource), &result)
	return &result
}
