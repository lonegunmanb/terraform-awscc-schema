package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDomainName = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_configuration": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `EndpointConfiguration` + "`" + `` + "`" + ` property type specifies the endpoint types of an Amazon API Gateway domain name.\n ` + "`" + `` + "`" + `EndpointConfiguration` + "`" + `` + "`" + ` is a property of the [AWS::ApiGateway::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html) resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
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
      "mutual_tls_authentication": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "truststore_uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "truststore_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "ownership_verification_certificate_arn": {
        "computed": true,
        "description": "The ARN of the public certificate issued by ACM to validate ownership of your custom domain. Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "regional_certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "regional_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "regional_hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "routing_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_policy": {
        "computed": true,
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
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::DomainName",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayDomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDomainName), &result)
	return &result
}
