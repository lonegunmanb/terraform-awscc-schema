package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptunegraphGraph = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "computed": true,
        "description": "Value that indicates whether the Graph has deletion protection enabled. The graph can't be deleted when deletion protection is enabled.\n\n_Default_: If not specified, the default value is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the graph. For example: ` + "`" + `g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_arn": {
        "computed": true,
        "description": "Graph resource ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_id": {
        "computed": true,
        "description": "The auto-generated id assigned by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_name": {
        "computed": true,
        "description": "Contains a user-supplied name for the Graph. \n\nIf you don't specify a name, we generate a unique Graph Name using a combination of Stack Name and a UUID comprising of 4 characters.\n\n_Important_: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
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
      "provisioned_memory": {
        "description": "Memory for the Graph.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "public_connectivity": {
        "computed": true,
        "description": "Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.\n\nWhen the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.\n\nWhen the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.\n\n_Default_: If not specified, the default value is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replica_count": {
        "computed": true,
        "description": "Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.\n\nReplica Count should always be less than or equal to 2.\n\n_Default_: If not specified, the default value is 1.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with this graph.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vector_search_configuration": {
        "computed": true,
        "description": "Vector Search Configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vector_search_dimension": {
              "description": "The vector search dimension",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "The AWS::NeptuneGraph::Graph resource creates an Amazon NeptuneGraph Graph.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNeptunegraphGraphSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptunegraphGraph), &result)
	return &result
}
