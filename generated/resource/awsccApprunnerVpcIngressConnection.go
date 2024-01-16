package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApprunnerVpcIngressConnection = `{
  "block": {
    "attributes": {
      "domain_name": {
        "computed": true,
        "description": "The Domain name associated with the VPC Ingress Connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ingress_vpc_configuration": {
        "description": "The configuration of customer?s VPC and related VPC endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_endpoint_id": {
              "description": "The ID of the VPC endpoint that your App Runner service connects to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_id": {
              "description": "The ID of the VPC that the VPC endpoint is used in.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "service_arn": {
        "description": "The Amazon Resource Name (ARN) of the service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the VpcIngressConnection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      "vpc_ingress_connection_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the VpcIngressConnection.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_ingress_connection_name": {
        "computed": true,
        "description": "The customer-provided Vpc Ingress Connection name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApprunnerVpcIngressConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApprunnerVpcIngressConnection), &result)
	return &result
}
