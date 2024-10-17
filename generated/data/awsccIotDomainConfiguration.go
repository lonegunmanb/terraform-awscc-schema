package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotDomainConfiguration = `{
  "block": {
    "attributes": {
      "application_protocol": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_authorizer_override": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "default_authorizer_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "client_certificate_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_certificate_callback_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_configuration_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_certificate_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "server_certificate_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_ocsp_check": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "server_certificates": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "server_certificate_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "server_certificate_status_detail": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "service_type": {
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
          "nesting_mode": "set"
        }
      },
      "tls_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "validation_certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::DomainConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotDomainConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotDomainConfiguration), &result)
	return &result
}
