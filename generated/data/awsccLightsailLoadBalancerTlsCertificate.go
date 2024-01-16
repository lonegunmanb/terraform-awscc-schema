package data

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
        "type": [
          "set",
          "string"
        ]
      },
      "certificate_domain_name": {
        "computed": true,
        "description": "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_name": {
        "computed": true,
        "description": "The SSL/TLS certificate name.",
        "description_kind": "plain",
        "type": "string"
      },
      "https_redirection_enabled": {
        "computed": true,
        "description": "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_attached": {
        "computed": true,
        "description": "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
        "description_kind": "plain",
        "type": "bool"
      },
      "load_balancer_name": {
        "computed": true,
        "description": "The name of your load balancer.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::Lightsail::LoadBalancerTlsCertificate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailLoadBalancerTlsCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailLoadBalancerTlsCertificate), &result)
	return &result
}
