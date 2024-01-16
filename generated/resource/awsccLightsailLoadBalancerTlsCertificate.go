package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailLoadBalancerTlsCertificate = `{
  "block": {
    "attributes": {
      "certificate_alternative_names": {
        "computed": true,
        "description": "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "certificate_domain_name": {
        "description": "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_name": {
        "description": "The SSL/TLS certificate name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "https_redirection_enabled": {
        "computed": true,
        "description": "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_attached": {
        "computed": true,
        "description": "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "load_balancer_name": {
        "description": "The name of your load balancer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "load_balancer_tls_certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The validation status of the SSL/TLS certificate.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailLoadBalancerTlsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailLoadBalancerTlsCertificate), &result)
	return &result
}
