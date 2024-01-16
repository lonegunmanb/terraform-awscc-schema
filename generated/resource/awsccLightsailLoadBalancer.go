package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailLoadBalancer = `{
  "block": {
    "attributes": {
      "attached_instances": {
        "computed": true,
        "description": "The names of the instances attached to the load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "health_check_path": {
        "computed": true,
        "description": "The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., \"/\").",
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
      "instance_port": {
        "description": "The instance port where you're creating your load balancer.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ip_address_type": {
        "computed": true,
        "description": "The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_name": {
        "description": "The name of your load balancer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_stickiness_enabled": {
        "computed": true,
        "description": "Configuration option to enable session stickiness.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "session_stickiness_lb_cookie_duration_seconds": {
        "computed": true,
        "description": "Configuration option to adjust session stickiness cookie duration parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "tls_policy_name": {
        "computed": true,
        "description": "The name of the TLS policy to apply to the load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::LoadBalancer",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailLoadBalancerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailLoadBalancer), &result)
	return &result
}
