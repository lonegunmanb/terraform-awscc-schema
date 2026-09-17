package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53TrafficPolicyInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the traffic policy instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "description": "The ID of the hosted zone that Amazon Route 53 creates the resource record sets in. The bare ID as Route 53 returns it, without a /hostedzone/ prefix.",
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
      "name": {
        "description": "The domain name, or subdomain name, for which Amazon Route 53 answers DNS queries by using the resource record sets it creates for this traffic policy instance. Must be lower-case and end with a trailing dot, which is the form Route 53 returns: Route 53 normalizes DNS names, so admitting another form would neither round-trip through Read nor guarantee that a change to this property changes the resource's identity.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the traffic policy instance. A steady-state instance is Applied.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_policy_id": {
        "description": "The ID of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone. Lower-case, as Route 53 returns it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_policy_instance_id": {
        "computed": true,
        "description": "The ID that Amazon Route 53 assigned to the traffic policy instance when it was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_policy_type": {
        "computed": true,
        "description": "The DNS type that Amazon Route 53 assigned to all of the resource record sets that it created for this traffic policy instance. Route 53 derives it from the referenced traffic policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_policy_version": {
        "description": "The version of the traffic policy that Amazon Route 53 uses to create resource record sets in the specified hosted zone.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "ttl": {
        "description": "The TTL that Amazon Route 53 assigns to all of the resource record sets that it creates in the specified hosted zone.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::Route53::TrafficPolicyInstance. Applies a version of a Route 53 traffic policy to a DNS name in a public hosted zone, and Route 53 creates and owns the resource record sets that answer queries for that name.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53TrafficPolicyInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53TrafficPolicyInstance), &result)
	return &result
}
