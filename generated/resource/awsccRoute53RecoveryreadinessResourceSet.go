package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoveryreadinessResourceSet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_set_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource set.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_set_name": {
        "computed": true,
        "description": "The name of the resource set to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_set_type": {
        "description": "The resource type of the resources in the resource set. Enter one of the following values for resource type: \n\nAWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resources": {
        "description": "A list of resource objects in the resource set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_id": {
              "computed": true,
              "description": "The component identifier of the resource, generated when DNS target resource is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dns_target_resource": {
              "computed": true,
              "description": "A component for DNS/routing control readiness checks.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_name": {
                    "computed": true,
                    "description": "The domain name that acts as an ingress point to a portion of the customer application.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hosted_zone_arn": {
                    "computed": true,
                    "description": "The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_set_id": {
                    "computed": true,
                    "description": "The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "record_type": {
                    "computed": true,
                    "description": "The type of DNS record of the target resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_resource": {
                    "computed": true,
                    "description": "The target resource that the Route 53 record points to.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "nlb_resource": {
                          "computed": true,
                          "description": "The Network Load Balancer resource that a DNS target resource points to.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "arn": {
                                "computed": true,
                                "description": "A Network Load Balancer resource Amazon Resource Name (ARN).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "r53_resource": {
                          "computed": true,
                          "description": "The Route 53 resource that a DNS target resource record points to.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "domain_name": {
                                "computed": true,
                                "description": "The DNS target domain name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "record_set_id": {
                                "computed": true,
                                "description": "The Resource Record set id.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "readiness_scopes": {
              "computed": true,
              "description": "A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the AWS resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "A tag to associate with the parameters for a resource set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoveryreadinessResourceSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoveryreadinessResourceSet), &result)
	return &result
}
