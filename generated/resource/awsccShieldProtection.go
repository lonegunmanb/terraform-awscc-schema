package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccShieldProtection = `{
  "block": {
    "attributes": {
      "application_layer_automatic_response_configuration": {
        "computed": true,
        "description": "The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description": "Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "block": {
                    "computed": true,
                    "description": "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF ` + "`" + `Block` + "`" + ` action.\nYou must specify exactly one action, either ` + "`" + `Block` + "`" + ` or ` + "`" + `Count` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "count": {
                    "computed": true,
                    "description": "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF ` + "`" + `Count` + "`" + ` action.\nYou must specify exactly one action, either ` + "`" + `Block` + "`" + ` or ` + "`" + `Count` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "status": {
              "computed": true,
              "description": "Indicates whether automatic application layer DDoS mitigation is enabled for the protection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "health_check_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the health check to associate with the protection.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Friendly name for the Protection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protection_arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the protection.",
        "description_kind": "plain",
        "type": "string"
      },
      "protection_id": {
        "computed": true,
        "description": "The unique identifier (ID) of the protection.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description": "The ARN (Amazon Resource Name) of the resource to be protected.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tag key-value pairs for the Protection object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Enables AWS Shield Advanced for a specific AWS resource. The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccShieldProtectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccShieldProtection), &result)
	return &result
}
