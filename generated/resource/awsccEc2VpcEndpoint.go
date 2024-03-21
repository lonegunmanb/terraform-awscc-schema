package resource

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
      "id": {
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
        "description": "An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.\n For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns_enabled": {
        "computed": true,
        "description": "Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ` + "`" + `` + "`" + `kinesis.us-east-1.amazonaws.com` + "`" + `` + "`" + `), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.\n To use a private hosted zone, you must set the following VPC attributes to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `: ` + "`" + `` + "`" + `enableDnsHostnames` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `enableDnsSupport` + "`" + `` + "`" + `.\n This property is supported only for interface endpoints.\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "route_table_ids": {
        "computed": true,
        "description": "The IDs of the route tables. Routing is supported only for gateway endpoints.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "service_name": {
        "description": "The name of the endpoint service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "vpc_endpoint_type": {
        "computed": true,
        "description": "The type of endpoint.\n Default: Gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).\n An endpoint of type ` + "`" + `` + "`" + `Interface` + "`" + `` + "`" + ` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.\n An endpoint of type ` + "`" + `` + "`" + `gateway` + "`" + `` + "`" + ` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [W",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpoint), &result)
	return &result
}
