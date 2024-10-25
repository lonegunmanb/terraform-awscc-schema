package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecordSet = `{
  "block": {
    "attributes": {
      "alias_target": {
        "computed": true,
        "description": "Alias resource record sets only: Information about the AWS resource, such as a CloudFront distribution or an Amazon S3 bucket, that you want to route traffic to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_name": {
              "computed": true,
              "description": "The value that you specify depends on where you want to route queries.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "evaluate_target_health": {
              "computed": true,
              "description": "When EvaluateTargetHealth is true, an alias resource record set inherits the health of the referenced AWS resource, such as an ELB load balancer or another resource record set in the hosted zone.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hosted_zone_id": {
              "computed": true,
              "description": "The value used depends on where you want to route traffic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cidr_routing_config": {
        "computed": true,
        "description": "The object that is specified in resource record set object when you are linking a resource record set to a CIDR location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "collection_id": {
              "computed": true,
              "description": "The CIDR collection ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location_name": {
              "computed": true,
              "description": "The CIDR collection location name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "comment": {
        "computed": true,
        "description": "Optional: Any comments you want to include about a change batch request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failover": {
        "computed": true,
        "description": "To configure failover, you add the Failover element to two resource record sets. For one resource record set, you specify PRIMARY as the value for Failover; for the other resource record set, you specify SECONDARY. In addition, you include the HealthCheckId element and specify the health check that you want Amazon Route 53 to perform for each resource record set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "geo_location": {
        "computed": true,
        "description": "A complex type that lets you control how Amazon Route 53 responds to DNS queries based on the geographic origin of the query.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "continent_code": {
              "computed": true,
              "description": "For geolocation resource record sets, a two-letter abbreviation that identifies a continent.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "country_code": {
              "computed": true,
              "description": "For geolocation resource record sets, the two-letter code for a country.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subdivision_code": {
              "computed": true,
              "description": "For geolocation resource record sets, the two-letter code for a state of the United States.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "health_check_id": {
        "computed": true,
        "description": "If you want Amazon Route 53 to return this resource record set in response to a DNS query only when the status of a health check is healthy, include the HealthCheckId element and specify the ID of the applicable health check.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description": "The ID of the hosted zone that you want to create records in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hosted_zone_name": {
        "computed": true,
        "description": "The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, www.example.com.) as part of the HostedZoneName.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_value_answer": {
        "computed": true,
        "description": "To route traffic approximately randomly to multiple resources, such as web servers, create one multivalue answer record for each resource and specify true for MultiValueAnswer.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The name of the record that you want to create, update, or delete.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Amazon EC2 Region where you created the resource that this resource record set refers to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_records": {
        "computed": true,
        "description": "One or more values that correspond with the value that you specified for the Type property.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "set_identifier": {
        "computed": true,
        "description": "An identifier that differentiates among multiple resource record sets that have the same combination of name and type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ttl": {
        "computed": true,
        "description": "The resource record cache time to live (TTL), in seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The DNS record type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "weight": {
        "computed": true,
        "description": "Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set. Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type. Route 53 then responds to queries based on the ratio of a resource's weight to the total.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::Route53::RecordSet.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecordSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecordSet), &result)
	return &result
}
