package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_task": {
        "computed": true,
        "description": "The details of the import task to use to create the graph. When specified, the graph is created using CreateGraphUsingImportTask and data is imported from the supplied source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "blank_node_handling": {
              "computed": true,
              "description": "The method to handle blank nodes in the dataset. Currently, only convertToIri is supported, meaning blank nodes are converted to unique IRIs at load time. Must be provided when format is NTRIPLES",
              "description_kind": "plain",
              "type": "string"
            },
            "fail_on_error": {
              "computed": true,
              "description": "If set to true, the task halts when an import error is encountered. If set to false, the task skips the data that caused the error and continues if possible.",
              "description_kind": "plain",
              "type": "bool"
            },
            "format": {
              "computed": true,
              "description": "Specifies the format of S3 data to be imported. Valid values are CSV, which identifies the Gremlin CSV format, OPEN_CYPHER, which identifies the openCypher load format, or NTRIPLES, which identifies the RDF n-triples format.",
              "description_kind": "plain",
              "type": "string"
            },
            "import_options": {
              "computed": true,
              "description": "Contains options for controlling the import process. For example, if the failOnError key is set to false, the import skips the data that caused the error and continues if possible (whereas if set to true, the default, or if omitted, the import operation halts immediately when an error is encountered).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "neptune": {
                    "computed": true,
                    "description": "Options for importing data from a Neptune database.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "preserve_default_vertex_labels": {
                          "computed": true,
                          "description": "Neptune Analytics supports label-less vertices and no labels are assigned unless one is explicitly provided. Neptune assigns default labels when none is explicitly provided. When importing the data into Neptune Analytics, the default vertex labels can be omitted by setting preserveDefaultVertexLabels to false. Note that if the vertex only has default labels, and has no other properties or edges, then the vertex will effectively not get imported into Neptune Analytics when preserveDefaultVertexLabels is set to false.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "preserve_edge_ids": {
                          "computed": true,
                          "description": "Neptune Analytics currently does not support user defined edge ids. The edge ids are not imported by default. They are imported if preserveEdgeIds is set to true, and ids are stored as properties on the relationships with the property name neptuneEdgeId.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "s3_export_kms_key_id": {
                          "computed": true,
                          "description": "The KMS key to use to encrypt data in the S3 bucket where the graph data is exported.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_export_path": {
                          "computed": true,
                          "description": "The path to an S3 bucket from which to import data.",
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
            "max_provisioned_memory": {
              "computed": true,
              "description": "The maximum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph. Default: 1024, or the approved upper limit for your account. If both the minimum and maximum values are specified, the final provisioned-memory will be chosen per the actual size of your imported data. If neither value is specified, 128 m-NCUs are used.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_provisioned_memory": {
              "computed": true,
              "description": "The minimum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph. Default: 16",
              "description_kind": "plain",
              "type": "number"
            },
            "parquet_type": {
              "computed": true,
              "description": "The parquet type of the import task. Required when Format is PARQUET.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that will allow access to the data that is to be imported.",
              "description_kind": "plain",
              "type": "string"
            },
            "source": {
              "computed": true,
              "description": "A URL identifying to the location of the data to be imported. This can be an Amazon S3 path, or can point to a Neptune database endpoint or snapshot.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "The ARN of the KMS key used to encrypt data in the Neptune Analytics graph. If not specified, the graph is encrypted with an AWS managed key.",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_memory": {
        "computed": true,
        "description": "Memory for the Graph.",
        "description_kind": "plain",
        "type": "number"
      },
      "public_connectivity": {
        "computed": true,
        "description": "Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.\n\nWhen the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.\n\nWhen the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.\n\n_Default_: If not specified, the default value is false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "replica_count": {
        "computed": true,
        "description": "Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.\n\nReplica Count should always be less than or equal to 2.\n\n_Default_: If not specified, the default value is 1.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with this graph.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vector_search_configuration": {
        "computed": true,
        "description": "Vector Search Configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vector_search_dimension": {
              "computed": true,
              "description": "The vector search dimension",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::NeptuneGraph::Graph",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNeptunegraphGraphSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptunegraphGraph), &result)
	return &result
}
