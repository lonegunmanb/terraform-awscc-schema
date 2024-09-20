package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferConnector = `{
  "block": {
    "attributes": {
      "access_role": {
        "description": "Specifies the access role for the connector.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "compression": {
              "computed": true,
              "description": "Compression setting for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_algorithm": {
              "computed": true,
              "description": "Encryption algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_profile_id": {
              "computed": true,
              "description": "A unique identifier for the local profile.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mdn_response": {
              "computed": true,
              "description": "MDN Response setting for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mdn_signing_algorithm": {
              "computed": true,
              "description": "MDN Signing algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message_subject": {
              "computed": true,
              "description": "The message subject for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "partner_profile_id": {
              "computed": true,
              "description": "A unique identifier for the partner profile.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "signing_algorithm": {
              "computed": true,
              "description": "Signing algorithm for this AS2 connector configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "connector_id": {
        "computed": true,
        "description": "A unique identifier for the connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description": "Specifies the logging role for the connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_policy_name": {
        "computed": true,
        "description": "Security policy for SFTP Connector",
        "description_kind": "plain",
        "optional": true,
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
            "trusted_host_keys": {
              "computed": true,
              "description": "List of public host keys, for the external server to which you are connecting.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "user_secret_id": {
              "computed": true,
              "description": "ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "url": {
        "description": "URL for Connector",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Transfer::Connector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferConnector), &result)
	return &result
}
