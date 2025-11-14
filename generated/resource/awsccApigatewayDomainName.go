package resource

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
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "domain_name_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_configuration": {
        "computed": true,
        "description": "The endpoint configuration of this DomainName showing the endpoint types and IP address types of the domain name.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "types": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "truststore_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "ownership_verification_certificate_arn": {
        "computed": true,
        "description": "The ARN of the public certificate issued by ACM to validate ownership of your custom domain. Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "regional_certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "security_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::DomainName` + "`" + `` + "`" + ` resource specifies a public custom domain name for your API in API Gateway.\n To create a custom domain name for private APIs, use [AWS::ApiGateway::DomainNameV2](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html).\n You can use a custom domain name to provide a URL that's more intuitive and easier to recall. For more information about using custom domain names, see [Set up Custom Domain Name for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDomainName), &result)
	return &result
}
