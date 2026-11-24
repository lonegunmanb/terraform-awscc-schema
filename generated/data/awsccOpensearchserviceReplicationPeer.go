package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserviceReplicationPeer = `{
  "block": {
    "attributes": {
      "domain_arn": {
        "computed": true,
        "description": "The ARN of the domain being replicated.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The name of the domain to replicate, which becomes the primary. The domain must be in the same account and Region as the stack. After a switchover or failover this domain may no longer be the primary.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_domain_arn": {
        "computed": true,
        "description": "The ARN of the domain to replicate with, which becomes the secondary. The peer domain must be in the same account and can be in the same or another Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_role": {
        "computed": true,
        "description": "The replication role of the peer domain (PRIMARY or SECONDARY), as of the last time CloudFormation read the resource. Switchover and failover change it outside CloudFormation; call DescribeDomainReplication for the current role.",
        "description_kind": "plain",
        "type": "string"
      },
      "role": {
        "computed": true,
        "description": "The replication role of the domain (PRIMARY or SECONDARY), as of the last time CloudFormation read the resource. Switchover and failover change it outside CloudFormation; call DescribeDomainReplication for the current role.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The state of the replication relationship, as of the last time CloudFormation read it.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchService::ReplicationPeer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserviceReplicationPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserviceReplicationPeer), &result)
	return &result
}
