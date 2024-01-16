package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftEndpointAccess = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description": "The DNS address of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_identifier": {
        "computed": true,
        "description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_create_time": {
        "computed": true,
        "description": "The time (UTC) that the endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_status": {
        "computed": true,
        "description": "The status of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the cluster accepts incoming connections.",
        "description_kind": "plain",
        "type": "number"
      },
      "resource_owner": {
        "computed": true,
        "description": "The AWS account ID of the owner of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_group_name": {
        "computed": true,
        "description": "The subnet group name where Amazon Redshift chooses to deploy the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint": {
        "computed": true,
        "description": "The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_interfaces": {
              "computed": true,
              "description": "One or more network interfaces of the endpoint. Also known as an interface endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone": {
                    "computed": true,
                    "description": "The Availability Zone.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description": "The network interface identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "private_ip_address": {
                    "computed": true,
                    "description": "The IPv4 address of the network interface within the subnet.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "The subnet identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "vpc_endpoint_id": {
              "computed": true,
              "description": "The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description": "The VPC identifier that the endpoint is associated.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "A list of vpc security group ids to apply to the created endpoint access.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_security_groups": {
        "computed": true,
        "description": "A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "The status of the VPC security group.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_security_group_id": {
              "computed": true,
              "description": "The identifier of the VPC security group.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Redshift::EndpointAccess",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftEndpointAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftEndpointAccess), &result)
	return &result
}
