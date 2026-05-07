package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccInterconnectConnection = `{
  "block": {
    "attributes": {
      "activation_key": {
        "computed": true,
        "description": "The activation key for accepting a connection proposal from a partner CSP. Mutually exclusive with EnvironmentId.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "attach_point": {
        "computed": true,
        "description": "The logical attachment point in your AWS network where the managed connection will be connected.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The ARN of the resource to attach to.",
              "description_kind": "plain",
              "type": "string"
            },
            "direct_connect_gateway": {
              "computed": true,
              "description": "The ID of the Direct Connect Gateway to attach to.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "bandwidth": {
        "computed": true,
        "description": "The bandwidth of the connection (e.g., 50Mbps, 1Gbps). Required when creating a connection through AWS.",
        "description_kind": "plain",
        "type": "string"
      },
      "billing_tier": {
        "computed": true,
        "description": "The billing tier for the connection.",
        "description_kind": "plain",
        "type": "number"
      },
      "connection_id": {
        "computed": true,
        "description": "The unique identifier for the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the environment for the connection. Required when creating a connection through AWS. Mutually exclusive with ActivationKey.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account": {
        "computed": true,
        "description": "The AWS account ID of the connection owner.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The partner cloud service provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloud_service_provider": {
              "computed": true,
              "description": "The name of the cloud service provider.",
              "description_kind": "plain",
              "type": "string"
            },
            "last_mile_provider": {
              "computed": true,
              "description": "The name of the last mile provider.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "remote_account": {
        "computed": true,
        "description": "The remote account identifier for the connection. Required when creating a connection through AWS. Replaces RemoteOwnerAccount.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "The identifier of the remote account.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "remote_owner_account": {
        "computed": true,
        "description": "Deprecated. Use RemoteAccount instead. The account ID of the remote owner. Required when creating a connection through AWS.",
        "description_kind": "plain",
        "type": "string"
      },
      "shared_id": {
        "computed": true,
        "description": "The shared identifier for the connection pairing.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of managed connection.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Interconnect::Connection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccInterconnectConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccInterconnectConnection), &result)
	return &result
}
