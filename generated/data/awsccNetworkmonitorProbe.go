package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmonitorProbe = `{
  "block": {
    "attributes": {
      "address_family": {
        "computed": true,
        "description": "The IP address family the probe uses, computed by the service from the source and destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time and date when the probe was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination": {
        "computed": true,
        "description": "The literal IPv4 or IPv6 address the probe sends traffic to. The service rejects a domain name and a CIDR block.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_port": {
        "computed": true,
        "description": "The port the probe sends traffic to. Applies to the TCP protocol only.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "The time and date when the probe was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_name": {
        "computed": true,
        "description": "The name of the monitor the probe belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "packet_size": {
        "computed": true,
        "description": "The size, in bytes, of the packets the probe sends.",
        "description_kind": "plain",
        "type": "number"
      },
      "probe_arn": {
        "computed": true,
        "description": "The ARN of the probe.",
        "description_kind": "plain",
        "type": "string"
      },
      "probe_id": {
        "computed": true,
        "description": "The ID of the probe.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol the probe uses.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_arn": {
        "computed": true,
        "description": "The ARN of the subnet the probe sends traffic from.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the probe.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the probe.",
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
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC containing the probe's source subnet.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkMonitor::Probe",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmonitorProbeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmonitorProbe), &result)
	return &result
}
