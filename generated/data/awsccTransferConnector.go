package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferConnector = `{
  "block": {
    "attributes": {
      "access_role": {
        "computed": true,
        "description": "Specifies the access role for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "as_2_config": {
        "computed": true,
        "description": "Configuration for an AS2 connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "basic_auth_secret_id": {
              "computed": true,
              "description": "ARN or name of the secret in AWS Secrets Manager which contains the credentials for Basic authentication. If empty, Basic authentication is disabled for the AS2 connector",
              "description_kind": "plain",
              "type": "string"
            },
            "compression": {
              "computed": true,
              "description": "Compression setting for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption_algorithm": {
              "computed": true,
              "description": "Encryption algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "local_profile_id": {
              "computed": true,
              "description": "A unique identifier for the local profile.",
              "description_kind": "plain",
              "type": "string"
            },
            "mdn_response": {
              "computed": true,
              "description": "MDN Response setting for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "mdn_signing_algorithm": {
              "computed": true,
              "description": "MDN Signing algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "message_subject": {
              "computed": true,
              "description": "The message subject for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "partner_profile_id": {
              "computed": true,
              "description": "A unique identifier for the partner profile.",
              "description_kind": "plain",
              "type": "string"
            },
            "preserve_content_type": {
              "computed": true,
              "description": "Specifies whether to use the AWS S3 object content-type as the content-type for the AS2 message.",
              "description_kind": "plain",
              "type": "string"
            },
            "signing_algorithm": {
              "computed": true,
              "description": "Signing algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "connector_id": {
        "computed": true,
        "description": "A unique identifier for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_config": {
        "computed": true,
        "description": "Egress configuration for the connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_lattice": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "port_number": {
                    "computed": true,
                    "description": "Port to connect to on the target VPC Lattice resource",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "resource_configuration_arn": {
                    "computed": true,
                    "description": "ARN of the VPC Lattice resource configuration",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "egress_type": {
        "computed": true,
        "description": "Specifies the egress type for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description": "Specifies the logging role for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_policy_name": {
        "computed": true,
        "description": "Security policy for SFTP Connector",
        "description_kind": "plain",
        "type": "string"
      },
      "service_managed_egress_ip_addresses": {
        "computed": true,
        "description": "The list of egress IP addresses of this connector. These IP addresses are assigned automatically when you create the connector.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sftp_config": {
        "computed": true,
        "description": "Configuration for an SFTP connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_concurrent_connections": {
              "computed": true,
              "description": "Specifies the number of active connections that your connector can establish with the remote server at the same time.",
              "description_kind": "plain",
              "type": "number"
            },
            "trusted_host_keys": {
              "computed": true,
              "description": "List of public host keys, for the external server to which you are connecting.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "user_secret_id": {
              "computed": true,
              "description": "ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for connectors. Tags are metadata attached to connectors for any purpose.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "url": {
        "computed": true,
        "description": "URL for Connector",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Transfer::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferConnector), &result)
	return &result
}
