package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2DomainName = `{
  "block": {
    "attributes": {
      "domain_name": {
        "computed": true,
        "description": "The custom domain name for your API in Amazon API Gateway. Uppercase letters are not supported.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_configurations": {
        "computed": true,
        "description": "The domain name configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_arn": {
              "computed": true,
              "description": "An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.",
              "description_kind": "plain",
              "type": "string"
            },
            "certificate_name": {
              "computed": true,
              "description": "The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.",
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint_type": {
              "computed": true,
              "description": "The endpoint type.",
              "description_kind": "plain",
              "type": "string"
            },
            "ownership_verification_certificate_arn": {
              "computed": true,
              "description": "The Amazon resource name (ARN) for the public certificate issued by ACMlong. This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.",
              "description_kind": "plain",
              "type": "string"
            },
            "security_policy": {
              "computed": true,
              "description": "The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are ` + "`" + `` + "`" + `TLS_1_0` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `TLS_1_2` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
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
        "description": "The mutual TLS authentication configuration for a custom domain name.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "truststore_uri": {
              "computed": true,
              "description": "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, ` + "`" + `` + "`" + `s3://bucket-name/key-name` + "`" + `` + "`" + `. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.",
              "description_kind": "plain",
              "type": "string"
            },
            "truststore_version": {
              "computed": true,
              "description": "The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
      "tags": {
        "computed": true,
        "description": "The collection of tags associated with a domain name.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::DomainName",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2DomainNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2DomainName), &result)
	return &result
}
