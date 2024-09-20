package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaWorkGroup = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The date and time the workgroup was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The workgroup description.",
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
      "name": {
        "description": "The workGroup name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recursive_delete_option": {
        "computed": true,
        "description": "The option to delete the workgroup and its contents even if the workgroup contains any named queries.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The state of the workgroup: ENABLED or DISABLED.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags, separated by commas, that you want to attach to the workgroup as you create it",
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
      },
      "work_group_configuration": {
        "computed": true,
        "description": "The workgroup configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_configuration": {
              "computed": true,
              "description": "Additional Configuration that are passed to Athena Spark Calculations running in this workgroup",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bytes_scanned_cutoff_per_query": {
              "computed": true,
              "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "customer_content_encryption_configuration": {
              "computed": true,
              "description": "Indicates the KMS key for encrypting notebook content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key": {
                    "computed": true,
                    "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enforce_work_group_configuration": {
              "computed": true,
              "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "engine_version": {
              "computed": true,
              "description": "The Athena engine version for running queries.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "effective_engine_version": {
                    "computed": true,
                    "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "selected_engine_version": {
                    "computed": true,
                    "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "execution_role": {
              "computed": true,
              "description": "Execution Role ARN required to run Athena Spark Calculations",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_cloudwatch_metrics_enabled": {
              "computed": true,
              "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "requester_pays_enabled": {
              "computed": true,
              "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "result_configuration": {
              "computed": true,
              "description": "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acl_configuration": {
                    "computed": true,
                    "description": "Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_acl_option": {
                          "computed": true,
                          "description": "The Amazon S3 canned ACL that Athena should specify when storing query results. Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "encryption_configuration": {
                    "computed": true,
                    "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_option": {
                          "computed": true,
                          "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key": {
                          "computed": true,
                          "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the owner of S3 bucket where query results are stored",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_location": {
                    "computed": true,
                    "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
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
      },
      "work_group_configuration_updates": {
        "computed": true,
        "description": "The workgroup configuration update object",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_configuration": {
              "computed": true,
              "description": "Additional Configuration that are passed to Athena Spark Calculations running in this workgroup",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bytes_scanned_cutoff_per_query": {
              "computed": true,
              "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "customer_content_encryption_configuration": {
              "computed": true,
              "description": "Indicates the KMS key for encrypting notebook content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key": {
                    "computed": true,
                    "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enforce_work_group_configuration": {
              "computed": true,
              "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "engine_version": {
              "computed": true,
              "description": "The Athena engine version for running queries.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "effective_engine_version": {
                    "computed": true,
                    "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "selected_engine_version": {
                    "computed": true,
                    "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "execution_role": {
              "computed": true,
              "description": "Execution Role ARN required to run Athena Spark Calculations",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "publish_cloudwatch_metrics_enabled": {
              "computed": true,
              "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "remove_bytes_scanned_cutoff_per_query": {
              "computed": true,
              "description": "Indicates that the data usage control limit per query is removed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "remove_customer_content_encryption_configuration": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "requester_pays_enabled": {
              "computed": true,
              "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "result_configuration_updates": {
              "computed": true,
              "description": "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acl_configuration": {
                    "computed": true,
                    "description": "Indicates that an Amazon S3 canned ACL should be set to control ownership of stored query results",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_acl_option": {
                          "computed": true,
                          "description": "The Amazon S3 canned ACL that Athena should specify when storing query results. Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "encryption_configuration": {
                    "computed": true,
                    "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_option": {
                          "computed": true,
                          "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "kms_key": {
                          "computed": true,
                          "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "expected_bucket_owner": {
                    "computed": true,
                    "description": "The AWS account ID of the owner of S3 bucket where query results are stored",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_location": {
                    "computed": true,
                    "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "remove_acl_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "remove_encryption_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "remove_expected_bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "remove_output_location": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
    "description": "Resource schema for AWS::Athena::WorkGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAthenaWorkGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaWorkGroup), &result)
	return &result
}
