package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2InstanceConnectEndpoint = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone of the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The ID of the Availability Zone of the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "The client token of the instance connect endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the EC2 Instance Connect Endpoint was created",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_connect_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_connect_endpoint_id": {
        "computed": true,
        "description": "The ID of the EC2 Instance Connect Endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description": "The ID of the elastic network interface that Amazon EC2 automatically created when creating the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the AWS account that created the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "preserve_client_ip": {
        "computed": true,
        "description": "Indicates whether your client's IP address is preserved as the source when you connect to a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "public_dns_names": {
        "computed": true,
        "description": "The public DNS names of the endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dualstack": {
              "computed": true,
              "description": "The dualstack DNS name of the EC2 Instance Connect Endpoint. A dualstack DNS name supports connections from both IPv4 and IPv6 clients.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "computed": true,
                    "description": "The DNS name of the EC2 Instance Connect Endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fips_dns_name": {
                    "computed": true,
                    "description": "The Federal Information Processing Standards (FIPS) compliant DNS name of the EC2 Instance Connect Endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ipv_4": {
              "computed": true,
              "description": "The IPv4-only DNS name of the EC2 Instance Connect Endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "computed": true,
                    "description": "The DNS name of the EC2 Instance Connect Endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "fips_dns_name": {
                    "computed": true,
                    "description": "The Federal Information Processing Standards (FIPS) compliant DNS name of the EC2 Instance Connect Endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "security_group_ids": {
        "computed": true,
        "description": "The security groups associated with the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "The current state of the EC2 Instance Connect Endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "The message for the current state of the EC2 Instance Connect Endpoint. Can include a failure message",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "description": "The ID of the subnet in which the EC2 Instance Connect Endpoint was created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the EC2 Instance Connect Endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC in which the EC2 Instance Connect Endpoint was created",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::InstanceConnectEndpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2InstanceConnectEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2InstanceConnectEndpoint), &result)
	return &result
}
