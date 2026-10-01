package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccScnDataIntegrationFlow = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The creation time of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The Amazon Web Services Supply Chain instance identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The last modified time of the DataIntegrationFlow.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the DataIntegrationFlow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sources": {
        "description": "The source configurations for the DataIntegrationFlow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dataset_source": {
              "computed": true,
              "description": "The dataset source configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_identifier": {
                    "computed": true,
                    "description": "The ARN of the dataset.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The dataset options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dedupe_records": {
                          "computed": true,
                          "description": "The option to perform deduplication on data records sharing same primary key values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "dedupe_strategy": {
                          "computed": true,
                          "description": "The deduplication strategy.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_priority": {
                                "computed": true,
                                "description": "The field priority deduplication strategy configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "The list of field names and their sort order for deduplication.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the deduplication field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "sort_order": {
                                            "computed": true,
                                            "description": "The sort order.",
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
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "type": {
                                "computed": true,
                                "description": "The deduplication strategy type.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "load_type": {
                          "computed": true,
                          "description": "The load type.",
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
            "s3_source": {
              "computed": true,
              "description": "The S3 source configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description": "The S3 bucket name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The Amazon S3 options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_type": {
                          "computed": true,
                          "description": "The file type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The S3 prefix.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "source_name": {
              "description": "The source name that can be used as table alias in SQL transformation query.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_type": {
              "description": "The source type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "The tags for the DataIntegrationFlow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target": {
        "description": "The DataIntegrationFlow target parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dataset_target": {
              "computed": true,
              "description": "The dataset target configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataset_identifier": {
                    "computed": true,
                    "description": "The dataset ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The dataset options.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dedupe_records": {
                          "computed": true,
                          "description": "The option to perform deduplication on data records sharing same primary key values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "dedupe_strategy": {
                          "computed": true,
                          "description": "The deduplication strategy.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_priority": {
                                "computed": true,
                                "description": "The field priority deduplication strategy configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "fields": {
                                      "computed": true,
                                      "description": "The list of field names and their sort order for deduplication.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the deduplication field.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "sort_order": {
                                            "computed": true,
                                            "description": "The sort order.",
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
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "type": {
                                "computed": true,
                                "description": "The deduplication strategy type.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "load_type": {
                          "computed": true,
                          "description": "The load type.",
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
            "target_type": {
              "description": "The target type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "transformation": {
        "description": "The DataIntegrationFlow transformation parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sql_transformation": {
              "computed": true,
              "description": "The SQL transformation configuration parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query": {
                    "computed": true,
                    "description": "The transformation SQL query body based on SparkSQL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "transformation_type": {
              "description": "The transformation type.",
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
    "description": "Represents an AWS Supply Chain data integration flow.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccScnDataIntegrationFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccScnDataIntegrationFlow), &result)
	return &result
}
