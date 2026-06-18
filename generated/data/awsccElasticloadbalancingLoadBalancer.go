package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingLoadBalancer = `{
  "block": {
    "attributes": {
      "access_logging_policy": {
        "computed": true,
        "description": "Information about where and how access logs are stored for the load balancer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "emit_interval": {
              "computed": true,
              "description": "The interval for publishing the access logs. You can specify an interval of either 5 minutes or 60 minutes.\n\nDefault: 60 minutes",
              "description_kind": "plain",
              "type": "number"
            },
            "enabled": {
              "computed": true,
              "description": "Specifies whether access logs are enabled for the load balancer.",
              "description_kind": "plain",
              "type": "bool"
            },
            "s3_bucket_name": {
              "computed": true,
              "description": "The name of the Amazon S3 bucket where the access logs are stored.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket_prefix": {
              "computed": true,
              "description": "The logical hierarchy you created for your Amazon S3 bucket, for example ` + "`" + `my-bucket-prefix/prod` + "`" + `. If the prefix is not provided, the log is placed at the root level of the bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "app_cookie_stickiness_policy": {
        "computed": true,
        "description": "Information about a policy for application-controlled session stickiness.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cookie_name": {
              "computed": true,
              "description": "The name of the application cookie used for stickiness.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_name": {
              "computed": true,
              "description": "The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "availability_zones": {
        "computed": true,
        "description": "The Availability Zones for a load balancer in a default VPC. For a load balancer in a nondefault VPC, specify Subnets instead.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "canonical_hosted_zone_name": {
        "computed": true,
        "description": "The name of the Route 53 hosted zone that is associated with the load balancer. Internal-facing load balancers.",
        "description_kind": "plain",
        "type": "string"
      },
      "canonical_hosted_zone_name_id": {
        "computed": true,
        "description": "The ID of the Route 53 hosted zone name that is associated with the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_draining_policy": {
        "computed": true,
        "description": "If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Specifies whether connection draining is enabled for the load balancer.",
              "description_kind": "plain",
              "type": "bool"
            },
            "timeout": {
              "computed": true,
              "description": "The maximum time, in seconds, to keep the existing connections open before deregistering the instances.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "connection_settings": {
        "computed": true,
        "description": "If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "idle_timeout": {
              "computed": true,
              "description": "The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cross_zone": {
        "computed": true,
        "description": "If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.",
        "description_kind": "plain",
        "type": "bool"
      },
      "dns_name": {
        "computed": true,
        "description": "The DNS name for the load balancer",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check": {
        "computed": true,
        "description": "The health check settings to use when evaluating the health of your EC2 instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "healthy_threshold": {
              "computed": true,
              "description": "The number of consecutive health checks successes required before moving the instance to the ` + "`" + `Healthy` + "`" + ` state.",
              "description_kind": "plain",
              "type": "string"
            },
            "interval": {
              "computed": true,
              "description": "The approximate interval, in seconds, between health checks of an individual instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "target": {
              "computed": true,
              "description": "The instance being checked.",
              "description_kind": "plain",
              "type": "string"
            },
            "timeout": {
              "computed": true,
              "description": "The amount of time, in seconds, during which no response means a failed health check.\n\nThis value must be less than the ` + "`" + `Interval` + "`" + ` value.",
              "description_kind": "plain",
              "type": "string"
            },
            "unhealthy_threshold": {
              "computed": true,
              "description": "The number of consecutive health check failures required before moving the instance to the ` + "`" + `Unhealthy` + "`" + ` state.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description": "The IDs of the instances for the load balancer.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "lb_cookie_stickiness_policy": {
        "computed": true,
        "description": "Information about a policy for duration-based session stickiness.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cookie_expiration_period": {
              "computed": true,
              "description": "The time period, in seconds, after which the cookie should be considered stale. If this parameter is not specified, the stickiness session lasts for the duration of the browser session.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_name": {
              "computed": true,
              "description": "The name of the policy. This name must be unique within the set of policies for this load balancer.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "listeners": {
        "computed": true,
        "description": "The Listeners for the load balancer. You can specify at most one listener per port.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instance_port": {
              "computed": true,
              "description": "The port on which the instance is listening.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_protocol": {
              "computed": true,
              "description": "The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.\n\nIf the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.\n\nIf there is another listener with the same ` + "`" + `InstancePort` + "`" + ` whose ` + "`" + `InstanceProtocol` + "`" + ` is secure, (HTTPS or SSL), the listener's ` + "`" + `InstanceProtocol` + "`" + ` must also be secure.\n\nIf there is another listener with the same ` + "`" + `InstancePort` + "`" + ` whose ` + "`" + `InstanceProtocol` + "`" + ` is HTTP or TCP, the listener's ` + "`" + `InstanceProtocol` + "`" + ` must be HTTP or TCP.",
              "description_kind": "plain",
              "type": "string"
            },
            "load_balancer_port": {
              "computed": true,
              "description": "The port on which the load balancer is listening. On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_names": {
              "computed": true,
              "description": "The names of the policies to associate with the listener.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "protocol": {
              "computed": true,
              "description": "The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_certificate_id": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the server certificate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "load_balancer_name": {
        "computed": true,
        "description": "The name of the load balancer. This name must be unique within your set of load balancers for the region.",
        "description_kind": "plain",
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "The policies defined for your Classic Load Balancer. Specify only back-end server policies.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "The policy attributes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "instance_ports": {
              "computed": true,
              "description": "The instance ports for the policy. Required only for some policy types.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "load_balancer_ports": {
              "computed": true,
              "description": "The load balancer ports for the policy. Required only for some policy types.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "policy_name": {
              "computed": true,
              "description": "The name of the policy.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_type": {
              "computed": true,
              "description": "The name of the policy type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "scheme": {
        "computed": true,
        "description": "The type of load balancer. Valid only for load balancers in a VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "The security groups for the load balancer. Valid only for load balancers in a VPC.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source_security_group": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_name": {
              "computed": true,
              "description": "The name of the security group that you can use as part of your inbound rules for your load balancer's back-end instances.",
              "description_kind": "plain",
              "type": "string"
            },
            "owner_alias": {
              "computed": true,
              "description": "The owner of the source security group.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "subnets": {
        "computed": true,
        "description": "The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with a load balancer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `aws:` + "`" + `. You can use any of the following characters: the set of Unicode letters, digits, whitespace, ` + "`" + `_` + "`" + `, ` + "`" + `.` + "`" + `, ` + "`" + `/` + "`" + `, ` + "`" + `=` + "`" + `, ` + "`" + `+` + "`" + `, and ` + "`" + `-` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ElasticLoadBalancing::LoadBalancer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticloadbalancingLoadBalancerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingLoadBalancer), &result)
	return &result
}
