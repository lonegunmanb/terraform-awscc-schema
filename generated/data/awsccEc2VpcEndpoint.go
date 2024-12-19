package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEndpoint = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_entries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "dns_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_record_ip_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_dns_only_for_inbound_resolver_endpoint": {
              "computed": true,
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
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "policy_document": {
        "computed": true,
        "description": "An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.\n For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. For example, if you have a JSON policy, you can convert it to YAML before including it in the YAML template, and CFNlong converts the policy to JSON format before calling the API actions for privatelink. Alternatively, you can include the JSON directly in the YAML, as shown in the following ` + "`" + `` + "`" + `Properties` + "`" + `` + "`" + ` section:\n ` + "`" + `` + "`" + `Properties: VpcEndpointType: 'Interface' ServiceName: !Sub 'com.amazonaws.${AWS::Region}.logs' PolicyDocument: '{ \"Version\":\"2012-10-17\", \"Statement\": [{ \"Effect\":\"Allow\", \"Principal\":\"*\", \"Action\":[\"logs:Describe*\",\"logs:Get*\",\"logs:List*\",\"logs:FilterLogEvents\"], \"Resource\":\"*\" }] }'` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "private_dns_enabled": {
        "computed": true,
        "description": "Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ` + "`" + `` + "`" + `kinesis.us-east-1.amazonaws.com` + "`" + `` + "`" + `), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.\n To use a private hosted zone, you must set the following VPC attributes to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `: ` + "`" + `` + "`" + `enableDnsHostnames` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `enableDnsSupport` + "`" + `` + "`" + `.\n This property is supported only for interface endpoints.\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_table_ids": {
        "computed": true,
        "description": "The IDs of the route tables. Routing is supported only for gateway endpoints.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "service_name": {
        "computed": true,
        "description": "The name of the endpoint service.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_network_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint_type": {
        "computed": true,
        "description": "The type of endpoint.\n Default: Gateway",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPCEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpoint), &result)
	return &result
}
