package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppfabricIngestionDestination = `{
  "block": {
    "attributes": {
      "app_bundle_arn": {
        "description": "The Amazon Resource Name (ARN) of the app bundle.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ingestion destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the ingestion destination was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_configuration": {
        "description": "Contains information about the destination of ingested data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audit_log": {
              "description": "Contains information about an audit log destination configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "description": "Contains information about an audit log destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "firehose_stream": {
                          "computed": true,
                          "description": "Contains information about an Amazon Kinesis Data Firehose delivery stream.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "stream_name": {
                                "computed": true,
                                "description": "The name of the Amazon Kinesis Data Firehose delivery stream.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "s3_bucket": {
                          "computed": true,
                          "description": "Contains information about an Amazon S3 bucket.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description": "The name of the Amazon S3 bucket.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "prefix": {
                                "computed": true,
                                "description": "The object key to use.",
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
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ingestion_arn": {
        "description": "The Amazon Resource Name (ARN) of the ingestion.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "processing_configuration": {
        "description": "Contains information about how ingested data is processed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audit_log": {
              "description": "Contains information about an audit log processing configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "format": {
                    "description": "The format in which the audit logs need to be formatted.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "schema": {
                    "description": "The event schema in which the audit logs need to be formatted.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The state of the ingestion destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of the key-value pairs of the tag or tags to assign to the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when the ingestion destination was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates an ingestion destination, which specifies how an application's ingested data is processed by Amazon Web Services AppFabric and where it's delivered.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppfabricIngestionDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppfabricIngestionDestination), &result)
	return &result
}
