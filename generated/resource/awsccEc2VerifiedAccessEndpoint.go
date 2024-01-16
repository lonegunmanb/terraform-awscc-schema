package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VerifiedAccessEndpoint = `{
  "block": {
    "attributes": {
      "application_domain": {
        "description": "The DNS name for users to reach your application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "attachment_type": {
        "description": "The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the AWS Verified Access endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "device_validation_domain": {
        "computed": true,
        "description": "Returned if endpoint has a device trust provider attached.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_certificate_arn": {
        "description": "The ARN of a public TLS/SSL certificate imported into or created with ACM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_domain": {
        "computed": true,
        "description": "A DNS name that is generated for the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_domain_prefix": {
        "description": "A custom identifier that gets prepended to a DNS name that is generated for the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_type": {
        "description": "The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.",
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
      "last_updated_time": {
        "computed": true,
        "description": "The last updated time.",
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_options": {
        "computed": true,
        "description": "The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "load_balancer_arn": {
              "computed": true,
              "description": "The ARN of the load balancer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The IP port number.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "computed": true,
              "description": "The IP protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_ids": {
              "computed": true,
              "description": "The IDs of the subnets.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "network_interface_options": {
        "computed": true,
        "description": "The options for network-interface type endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_interface_id": {
              "computed": true,
              "description": "The ID of the network interface.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The IP port number.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "computed": true,
              "description": "The IP protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "policy_document": {
        "computed": true,
        "description": "The AWS Verified Access policy document.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_enabled": {
        "computed": true,
        "description": "The status of the Verified Access policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of the security groups for the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "sse_specification": {
        "computed": true,
        "description": "The configuration options for customer provided KMS encryption.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_managed_key_enabled": {
              "computed": true,
              "description": "Whether to encrypt the policy with the provided key or disable encryption",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "KMS Key Arn used to encrypt the group policy",
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
        "description": "The endpoint status.",
        "description_kind": "plain",
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
      "verified_access_endpoint_id": {
        "computed": true,
        "description": "The ID of the AWS Verified Access endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_access_group_id": {
        "description": "The ID of the AWS Verified Access group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "verified_access_instance_id": {
        "computed": true,
        "description": "The ID of the AWS Verified Access instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The AWS::EC2::VerifiedAccessEndpoint resource creates an AWS EC2 Verified Access Endpoint.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VerifiedAccessEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VerifiedAccessEndpoint), &result)
	return &result
}
