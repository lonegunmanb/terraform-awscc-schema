package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorPrefetchSchedule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the prefetch schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "consumption": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "avail_matching_criteria": {
              "computed": true,
              "description": "If you only want MediaTailor to insert prefetched ads into avails that match specific dynamic variables, set the avail matching criteria.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dynamic_variable": {
                    "computed": true,
                    "description": "The dynamic variable(s) that MediaTailor should use as avail matching criteria.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "computed": true,
                    "description": "For the DynamicVariable specified in AvailMatchingCriteria, the Operator that is used for the comparison.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "end_time": {
              "computed": true,
              "description": "The time when MediaTailor no longer considers the prefetched ads for use in an ad break, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "The time when prefetched ads are considered for use in an ad break, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name to assign to the prefetch schedule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "playback_configuration_name": {
        "description": "The name of the playback configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recurring_prefetch_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_time": {
              "computed": true,
              "description": "The end time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recurring_consumption": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "avail_matching_criteria": {
                    "computed": true,
                    "description": "The configuration for the dynamic variables that determine which ad breaks that MediaTailor inserts prefetched ads in.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dynamic_variable": {
                          "computed": true,
                          "description": "The dynamic variable(s) that MediaTailor should use as avail matching criteria.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "operator": {
                          "computed": true,
                          "description": "For the DynamicVariable specified in AvailMatchingCriteria, the Operator that is used for the comparison.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "retrieved_ad_expiration_seconds": {
                    "computed": true,
                    "description": "The number of seconds that an ad is available for insertion after it was prefetched.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "recurring_retrieval": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delay_after_avail_end_seconds": {
                    "computed": true,
                    "description": "The number of seconds that MediaTailor waits after an ad avail before prefetching ads for the next avail.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "dynamic_variables": {
                    "computed": true,
                    "description": "The dynamic variables to use for substitution during prefetch requests to the ADS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "traffic_shaping_retrieval_window": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "retrieval_window_duration_seconds": {
                          "computed": true,
                          "description": "The amount of time, in seconds, that MediaTailor spreads prefetch requests to the ADS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "traffic_shaping_tps_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "peak_concurrent_users": {
                          "computed": true,
                          "description": "The expected peak number of concurrent viewers for your content.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "peak_tps": {
                          "computed": true,
                          "description": "The maximum number of transactions per second (TPS) that your ad decision server (ADS) can handle.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "traffic_shaping_type": {
                    "computed": true,
                    "description": "Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "start_time": {
              "computed": true,
              "description": "The start time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "retrieval": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dynamic_variables": {
              "computed": true,
              "description": "The dynamic variables to use for substitution during prefetch requests to the ad decision server (ADS).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "end_time": {
              "computed": true,
              "description": "The time when prefetch retrieval ends for the ad break, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "The time when prefetch retrievals can start for this break, as an ISO 8601 date-time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "traffic_shaping_retrieval_window": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "retrieval_window_duration_seconds": {
                    "computed": true,
                    "description": "The amount of time, in seconds, that MediaTailor spreads prefetch requests to the ADS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "traffic_shaping_tps_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "peak_concurrent_users": {
                    "computed": true,
                    "description": "The expected peak number of concurrent viewers for your content.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "peak_tps": {
                    "computed": true,
                    "description": "The maximum number of transactions per second (TPS) that your ad decision server (ADS) can handle.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "traffic_shaping_type": {
              "computed": true,
              "description": "Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "schedule_type": {
        "computed": true,
        "description": "The frequency that MediaTailor creates prefetch schedules.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stream_id": {
        "computed": true,
        "description": "An optional stream identifier that MediaTailor uses to prefetch ads for multiple streams that use the same playback configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the prefetch schedule.",
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
    "description": "Definition of AWS::MediaTailor::PrefetchSchedule Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorPrefetchScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorPrefetchSchedule), &result)
	return &result
}
