package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeMediaInsightsPipelineConfiguration = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The time at which the configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "elements": {
        "description": "The elements in the configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_transcribe_call_analytics_processor_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "call_analytics_stream_categories": {
                    "computed": true,
                    "description": "The categories to send to the insights target.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "content_identification_type": {
                    "computed": true,
                    "description": "Labels all PII identified in the transcript.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_redaction_type": {
                    "computed": true,
                    "description": "Redacts all PII identified in the transcript.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_partial_results_stabilization": {
                    "computed": true,
                    "description": "Enables partial result stabilization.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "filter_partial_results": {
                    "computed": true,
                    "description": "If true, partial results are filtered out.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "language_code": {
                    "computed": true,
                    "description": "The language code in the configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "language_model_name": {
                    "computed": true,
                    "description": "The name of the custom language model.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partial_results_stability": {
                    "computed": true,
                    "description": "The level of stability for partial results.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pii_entity_types": {
                    "computed": true,
                    "description": "The types of PII to redact.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "post_call_analytics_settings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_redaction_output": {
                          "computed": true,
                          "description": "The content redaction output settings.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_access_role_arn": {
                          "computed": true,
                          "description": "The ARN of the role used by Transcribe to upload post-call analysis.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "output_encryption_kms_key_id": {
                          "computed": true,
                          "description": "The ID of the KMS key used to encrypt the output.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "output_location": {
                          "computed": true,
                          "description": "The URL of the Amazon S3 bucket for post-call data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "vocabulary_filter_method": {
                    "computed": true,
                    "description": "The vocabulary filtering method.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_name": {
                    "computed": true,
                    "description": "The name of the custom vocabulary filter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_name": {
                    "computed": true,
                    "description": "The name of the custom vocabulary.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "amazon_transcribe_processor_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_identification_type": {
                    "computed": true,
                    "description": "Labels all PII identified in the transcript.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_redaction_type": {
                    "computed": true,
                    "description": "Redacts all PII identified in the transcript.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_partial_results_stabilization": {
                    "computed": true,
                    "description": "Enables partial result stabilization.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "filter_partial_results": {
                    "computed": true,
                    "description": "If true, partial results are filtered out.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "identify_language": {
                    "computed": true,
                    "description": "Turns language identification on or off.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "identify_multiple_languages": {
                    "computed": true,
                    "description": "Turns multiple language identification on or off.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "language_code": {
                    "computed": true,
                    "description": "The language code.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "language_model_name": {
                    "computed": true,
                    "description": "The name of the custom language model.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "language_options": {
                    "computed": true,
                    "description": "The language options for transcription.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "partial_results_stability": {
                    "computed": true,
                    "description": "The level of stability for partial results.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pii_entity_types": {
                    "computed": true,
                    "description": "The types of PII to redact.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "preferred_language": {
                    "computed": true,
                    "description": "The preferred language for transcription.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "show_speaker_label": {
                    "computed": true,
                    "description": "Enables speaker partitioning.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "vocabulary_filter_method": {
                    "computed": true,
                    "description": "The vocabulary filtering method.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_name": {
                    "computed": true,
                    "description": "The name of the custom vocabulary filter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_filter_names": {
                    "computed": true,
                    "description": "The names of the custom vocabulary filters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_name": {
                    "computed": true,
                    "description": "The name of the custom vocabulary.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vocabulary_names": {
                    "computed": true,
                    "description": "The names of the custom vocabularies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "kinesis_data_stream_sink_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "insights_target": {
                    "computed": true,
                    "description": "The ARN of the Kinesis Data Stream sink.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_recording_sink_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "The default URI of the Amazon S3 bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "recording_file_format": {
                    "computed": true,
                    "description": "The recording file format.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "description": "The element type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_insights_pipeline_configuration_arn": {
        "computed": true,
        "description": "The ARN of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_insights_pipeline_configuration_id": {
        "computed": true,
        "description": "The unique identifier of the configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_insights_pipeline_configuration_name": {
        "description": "The name of the media insights pipeline configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "real_time_alert_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disabled": {
              "computed": true,
              "description": "Turns off real-time alerts.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rules": {
              "computed": true,
              "description": "The rules in the alert.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "issue_detection_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rule_name": {
                          "computed": true,
                          "description": "The name of the issue detection rule.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "keyword_match_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "keywords": {
                          "computed": true,
                          "description": "The keywords or phrases to match.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "negate": {
                          "computed": true,
                          "description": "Matches keywords on their presence or absence.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "rule_name": {
                          "computed": true,
                          "description": "The name of the keyword match rule.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sentiment_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rule_name": {
                          "computed": true,
                          "description": "The name of the sentiment rule.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sentiment_type": {
                          "computed": true,
                          "description": "The type of sentiment.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "time_period": {
                          "computed": true,
                          "description": "The analysis interval in seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of alert rule.",
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
      "resource_access_role_arn": {
        "description": "The ARN of the role used by the service to access Amazon Web Services resources.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the configuration.",
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
      "updated_timestamp": {
        "computed": true,
        "description": "The time at which the configuration was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for an Amazon Chime SDK Media Insights Pipeline Configuration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeMediaInsightsPipelineConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeMediaInsightsPipelineConfiguration), &result)
	return &result
}
