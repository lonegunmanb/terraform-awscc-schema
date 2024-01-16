package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2TargetGroup = `{
  "block": {
    "attributes": {
      "health_check_enabled": {
        "computed": true,
        "description": "Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "health_check_interval_seconds": {
        "computed": true,
        "description": "The approximate amount of time, in seconds, between health checks of an individual target.",
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_path": {
        "computed": true,
        "description": "[HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_port": {
        "computed": true,
        "description": "The port the load balancer uses when performing health checks on targets. ",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_protocol": {
        "computed": true,
        "description": "The protocol the load balancer uses when performing health checks on targets. ",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_timeout_seconds": {
        "computed": true,
        "description": "The amount of time, in seconds, during which no response from a target means a failed health check.",
        "description_kind": "plain",
        "type": "number"
      },
      "healthy_threshold_count": {
        "computed": true,
        "description": "The number of consecutive health checks successes required before considering an unhealthy target healthy. ",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description": "The type of IP address used for this target group. The possible values are ipv4 and ipv6. ",
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "matcher": {
        "computed": true,
        "description": "[HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "grpc_code": {
              "computed": true,
              "description": "You can specify values between 0 and 99. You can specify multiple values, or a range of values. The default value is 12.",
              "description_kind": "plain",
              "type": "string"
            },
            "http_code": {
              "computed": true,
              "description": "For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200. You can specify multiple values or a range of values. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the target group.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.",
        "description_kind": "plain",
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol to use for routing traffic to the targets.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol_version": {
        "computed": true,
        "description": "[HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The value for the tag. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key name of the tag. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_group_arn": {
        "computed": true,
        "description": "The ARN of the Target Group",
        "description_kind": "plain",
        "type": "string"
      },
      "target_group_attributes": {
        "computed": true,
        "description": "The attributes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The value of the attribute.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The name of the attribute.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_group_full_name": {
        "computed": true,
        "description": "The full name of the target group.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_group_name": {
        "computed": true,
        "description": "The name of the target group.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description": "The type of target that you must specify when registering targets with this target group. You can't specify targets for a target group using more than one target type.",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "The targets.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "An Availability Zone or all. This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "The ID of the target. If the target type of the target group is instance, specify an instance ID. If the target type is ip, specify an IP address. If the target type is lambda, specify the ARN of the Lambda function. If the target type is alb, specify the ARN of the Application Load Balancer target. ",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port on which the target is listening. If the target group protocol is GENEVE, the supported port is 6081. If the target type is alb, the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "set"
        }
      },
      "unhealthy_threshold_count": {
        "computed": true,
        "description": "The number of consecutive health check failures required before considering a target unhealthy.",
        "description_kind": "plain",
        "type": "number"
      },
      "vpc_id": {
        "computed": true,
        "description": "The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ElasticLoadBalancingV2::TargetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticloadbalancingv2TargetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2TargetGroup), &result)
	return &result
}
