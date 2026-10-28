package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotOtaUpdate = `{
  "block": {
    "attributes": {
      "additional_parameters": {
        "computed": true,
        "description": "A list of additional OTA update parameters, which are name-value pairs.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "aws_job_abort_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "abort_criteria_list": {
              "computed": true,
              "description": "The list of criteria that determine when and how to abort the job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The type of job action to take to initiate the job abort.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "failure_type": {
                    "computed": true,
                    "description": "The type of job execution failures that can initiate a job abort.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "min_number_of_executed_things": {
                    "computed": true,
                    "description": "The minimum number of things which must receive job execution notifications before the job can be aborted.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "threshold_percentage": {
                    "computed": true,
                    "description": "The minimum percentage of job execution failures that must occur to initiate the job abort.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "aws_job_executions_rollout_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exponential_rate": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_rate_per_minute": {
                    "computed": true,
                    "description": "The minimum number of things that will be notified of a pending job, per minute, at the start of the job rollout.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "increment_factor": {
                    "computed": true,
                    "description": "The rate of increase for a job rollout.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "rate_increase_criteria": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "number_of_notified_things": {
                          "computed": true,
                          "description": "When this number of things have been notified, it will initiate an increase in the rollout rate.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "number_of_succeeded_things": {
                          "computed": true,
                          "description": "When this number of things have succeeded in their job execution, it will initiate an increase in the rollout rate.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maximum_per_minute": {
              "computed": true,
              "description": "The maximum number of OTA update job executions started per minute.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "aws_job_presigned_url_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expires_in_sec": {
              "computed": true,
              "description": "How long (in seconds) pre-signed URLs are valid. Valid values are 60 - 3600, the default value is 1800 seconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "aws_job_timeout_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "in_progress_timeout_in_minutes": {
              "computed": true,
              "description": "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "creation_date": {
        "computed": true,
        "description": "The date when the OTA update was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the OTA update.",
        "description_kind": "plain",
        "type": "string"
      },
      "files": {
        "computed": true,
        "description": "The files to be streamed by the OTA update.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "A list of name-attribute pairs.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "code_signing": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_signer_job_id": {
                    "computed": true,
                    "description": "The ID of the AWSSignerJob which was created to sign the file.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "custom_code_signing": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "certificate_chain": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "certificate_name": {
                                "computed": true,
                                "description": "The name of the certificate.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "inline_document": {
                                "computed": true,
                                "description": "A base64 encoded binary representation of the code signing certificate chain.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "hash_algorithm": {
                          "computed": true,
                          "description": "The hash algorithm used to code sign the file.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "signature": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "inline_document": {
                                "computed": true,
                                "description": "A base64 encoded binary representation of the code signing signature.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "signature_algorithm": {
                          "computed": true,
                          "description": "The signature algorithm used to code sign the file.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "start_signing_job_parameter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_destination": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "bucket": {
                                      "computed": true,
                                      "description": "The S3 bucket that contains the updated firmware.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prefix": {
                                      "computed": true,
                                      "description": "The S3 prefix.",
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
                        "signing_profile_name": {
                          "computed": true,
                          "description": "The code-signing profile name.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "signing_profile_parameter": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "certificate_arn": {
                                "computed": true,
                                "description": "Certificate ARN.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "certificate_path_on_device": {
                                "computed": true,
                                "description": "The location of the code-signing certificate on your device.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "platform": {
                                "computed": true,
                                "description": "The hardware platform of your device.",
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "file_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_location": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The S3 bucket.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description": "The S3 key.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description": "The S3 bucket version.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "stream": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_id": {
                          "computed": true,
                          "description": "The ID of a file associated with a stream.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stream_id": {
                          "computed": true,
                          "description": "The stream ID.",
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
            "file_name": {
              "computed": true,
              "description": "The name of the file.",
              "description_kind": "plain",
              "type": "string"
            },
            "file_type": {
              "computed": true,
              "description": "An integer value you can include in the job document to allow your devices to identify the type of file received from the cloud.",
              "description_kind": "plain",
              "type": "number"
            },
            "file_version": {
              "computed": true,
              "description": "The file version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description": "The date when the OTA update was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "ota_update_arn": {
        "computed": true,
        "description": "The OTA update ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "ota_update_id": {
        "computed": true,
        "description": "The ID of the OTA update to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "ota_update_status": {
        "computed": true,
        "description": "The status of the OTA update.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocols": {
        "computed": true,
        "description": "The protocol used to transfer the OTA update image.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "computed": true,
        "description": "The IAM role that grants AWS IoT Core access to the Amazon S3, IoT jobs and AWS Code Signing resources to create an OTA update job.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata which can be used to manage updates.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_selection": {
        "computed": true,
        "description": "Specifies whether the update will continue to run (CONTINUOUS), or will be complete after all things specified as targets have completed the update (SNAPSHOT).",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "The devices targeted to receive OTA updates.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::IoT::OTAUpdate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotOtaUpdateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotOtaUpdate), &result)
	return &result
}
