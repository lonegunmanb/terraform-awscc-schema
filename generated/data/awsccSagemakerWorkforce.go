package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerWorkforce = `{
  "block": {
    "attributes": {
      "cognito_config": {
        "computed": true,
        "description": "The configuration of an Amazon Cognito workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_id": {
              "computed": true,
              "description": "The client ID for your Amazon Cognito user pool.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_pool": {
              "computed": true,
              "description": "The ID for your Amazon Cognito user pool.",
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
        "description": "The IP address type for the workforce. IPv4 only or dualstack (IPv4 and IPv6).",
        "description_kind": "plain",
        "type": "string"
      },
      "oidc_config": {
        "computed": true,
        "description": "The configuration of an OIDC Identity Provider (IdP) private workforce.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_request_extra_params": {
              "computed": true,
              "description": "A string to string map of identifiers specific to the custom identity provider (IdP) being used.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "authorization_endpoint": {
              "computed": true,
              "description": "The OIDC IdP authorization endpoint used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_id": {
              "computed": true,
              "description": "The OIDC IdP client ID used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_secret": {
              "computed": true,
              "description": "The OIDC IdP client secret used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "issuer": {
              "computed": true,
              "description": "The OIDC IdP issuer used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "jwks_uri": {
              "computed": true,
              "description": "The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "logout_endpoint": {
              "computed": true,
              "description": "The OIDC IdP logout endpoint used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "scope": {
              "computed": true,
              "description": "An array of string identifiers used to refer to the specific pieces of user data or claims that the client application wants to access.",
              "description_kind": "plain",
              "type": "string"
            },
            "token_endpoint": {
              "computed": true,
              "description": "The OIDC IdP token endpoint used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            },
            "user_info_endpoint": {
              "computed": true,
              "description": "The OIDC IdP user info endpoint used to configure your private workforce.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_ip_config": {
        "computed": true,
        "description": "A list of IP address ranges used to access your training data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidrs": {
              "computed": true,
              "description": "A list of one to ten Classless Inter-Domain Routing (CIDR) values.",
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
      "sub_domain": {
        "computed": true,
        "description": "The subdomain for your OIDC Identity Provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "workforce_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the private workforce.",
        "description_kind": "plain",
        "type": "string"
      },
      "workforce_name": {
        "computed": true,
        "description": "The name of the private workforce.",
        "description_kind": "plain",
        "type": "string"
      },
      "workforce_vpc_config": {
        "computed": true,
        "description": "The VPC configuration for the workforce.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The VPC security group IDs.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description": "The VPC subnets.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description": "The ID of the VPC.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Workforce",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerWorkforceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerWorkforce), &result)
	return &result
}
