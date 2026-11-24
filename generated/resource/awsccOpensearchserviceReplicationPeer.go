package resource

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
        "description": "The name of the domain to replicate, which becomes the primary. The domain must be in the same account and Region as the stack. After a switchover or failover this domain may no longer be the primary.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_domain_arn": {
        "description": "The ARN of the domain to replicate with, which becomes the secondary. The peer domain must be in the same account and can be in the same or another Region.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Resource Type definition for AWS::OpenSearchService::ReplicationPeer. Replicates an Amazon OpenSearch Service domain to a peer domain, in the same or another Region. The domain named in DomainName becomes the primary when the resource is created; switchover and failover are performed outside CloudFormation and do not affect this resource. Creating and deleting the resource requires es:AddReplicationPeer and es:RemoveReplicationPeer on both the domain and the peer domain. After a failover the relationship is EJECTED on both domains (on the former primary once its Region is reachable again) and cannot be resumed. To replicate again, delete this resource, which removes the ejected relationship from both domains, then create a new ReplicationPeer with the new primary as DomainName, in a stack in that domain's Region; the new relationship is set up from scratch. Changing DomainName or PeerDomainArn in a single update fails while either domain is still in the existing relationship, because the replacement is added before the existing relationship is removed.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpensearchserviceReplicationPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserviceReplicationPeer), &result)
	return &result
}
