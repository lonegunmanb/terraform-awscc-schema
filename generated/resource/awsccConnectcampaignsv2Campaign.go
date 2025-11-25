package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectcampaignsv2Campaign = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Connect Campaign Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_subtype_config": {
        "description": "The possible types of channel subtype config parameters",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "Email Channel Subtype config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity": {
                    "computed": true,
                    "description": "Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "default_outbound_config": {
                    "computed": true,
                    "description": "Default SMS outbound config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "connect_source_email_address": {
                          "computed": true,
                          "description": "Email address used for Email messages",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "source_email_address_display_name": {
                          "computed": true,
                          "description": "The name of the source email address display name",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "wisdom_template_arn": {
                          "computed": true,
                          "description": "Arn",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "outbound_mode": {
                    "computed": true,
                    "description": "Email Outbound Mode",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agentless_config": {
                          "computed": true,
                          "description": "Agentless config",
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
            "sms": {
              "computed": true,
              "description": "SMS Channel Subtype config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity": {
                    "computed": true,
                    "description": "Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "default_outbound_config": {
                    "computed": true,
                    "description": "Default SMS outbound config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "connect_source_phone_number_arn": {
                          "computed": true,
                          "description": "Arn",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "wisdom_template_arn": {
                          "computed": true,
                          "description": "Arn",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "outbound_mode": {
                    "computed": true,
                    "description": "SMS Outbound Mode",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agentless_config": {
                          "computed": true,
                          "description": "Agentless config",
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
            "telephony": {
              "computed": true,
              "description": "Telephony Channel Subtype config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity": {
                    "computed": true,
                    "description": "Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "connect_queue_id": {
                    "computed": true,
                    "description": "The queue for the call",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "default_outbound_config": {
                    "computed": true,
                    "description": "Default Telephone Outbound config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "answer_machine_detection_config": {
                          "computed": true,
                          "description": "The configuration used for answering machine detection during outbound calls",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "await_answer_machine_prompt": {
                                "computed": true,
                                "description": "Enables detection of prompts (e.g., beep after after a voicemail greeting)",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "enable_answer_machine_detection": {
                                "computed": true,
                                "description": "Flag to decided whether outbound calls should have answering machine detection enabled or not",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "connect_contact_flow_id": {
                          "computed": true,
                          "description": "The identifier of the contact flow for the outbound call",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "connect_source_phone_number": {
                          "computed": true,
                          "description": "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "ring_timeout": {
                          "computed": true,
                          "description": "Maximum ring time for outbound calls in seconds",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "outbound_mode": {
                    "computed": true,
                    "description": "Telephony Outbound Mode",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agentless_config": {
                          "computed": true,
                          "description": "Agentless config",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "predictive_config": {
                          "computed": true,
                          "description": "Predictive config",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bandwidth_allocation": {
                                "computed": true,
                                "description": "The bandwidth allocation of a queue resource.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "preview_config": {
                          "computed": true,
                          "description": "Preview config",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_actions": {
                                "computed": true,
                                "description": "Actions that can be performed by agent during preview phase",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "bandwidth_allocation": {
                                "computed": true,
                                "description": "The bandwidth allocation of a queue resource.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "timeout_config": {
                                "computed": true,
                                "description": "Timeout Config for preview contacts",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "duration_in_seconds": {
                                      "computed": true,
                                      "description": "Timeout duration for a preview contact in seconds",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
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
                        "progressive_config": {
                          "computed": true,
                          "description": "Progressive config",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bandwidth_allocation": {
                                "computed": true,
                                "description": "The bandwidth allocation of a queue resource.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "communication_limits_override": {
        "computed": true,
        "description": "Communication limits config",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "all_channels_subtypes": {
              "computed": true,
              "description": "Communication limits",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "communication_limit_list": {
                    "computed": true,
                    "description": "List of communication limit",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "frequency": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_count_per_recipient": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "unit": {
                          "computed": true,
                          "description": "The communication limit time unit",
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
            "instance_limits_handling": {
              "computed": true,
              "description": "Enumeration of Instance Limits handling in a Campaign",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "communication_time_config": {
        "computed": true,
        "description": "Campaign communication time config",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "Time window config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "open_hours": {
                    "computed": true,
                    "description": "Open Hours config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "daily_hours": {
                          "computed": true,
                          "description": "Daily Hours map",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "Day of week",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "List of time range",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "end_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
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
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "restricted_periods": {
                    "computed": true,
                    "description": "Restricted period config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "restricted_period_list": {
                          "computed": true,
                          "description": "List of restricted period",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "end_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "local_time_zone_config": {
              "computed": true,
              "description": "Local time zone config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_time_zone": {
                    "computed": true,
                    "description": "Time Zone Id in the IANA format",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "local_time_zone_detection": {
                    "computed": true,
                    "description": "Local TimeZone Detection method list",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sms": {
              "computed": true,
              "description": "Time window config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "open_hours": {
                    "computed": true,
                    "description": "Open Hours config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "daily_hours": {
                          "computed": true,
                          "description": "Daily Hours map",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "Day of week",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "List of time range",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "end_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
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
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "restricted_periods": {
                    "computed": true,
                    "description": "Restricted period config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "restricted_period_list": {
                          "computed": true,
                          "description": "List of restricted period",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "end_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "telephony": {
              "computed": true,
              "description": "Time window config",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "open_hours": {
                    "computed": true,
                    "description": "Open Hours config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "daily_hours": {
                          "computed": true,
                          "description": "Daily Hours map",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "Day of week",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "List of time range",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "end_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
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
                            "nesting_mode": "set"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "restricted_periods": {
                    "computed": true,
                    "description": "Restricted period config",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "restricted_period_list": {
                          "computed": true,
                          "description": "List of restricted period",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "end_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
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
      "connect_campaign_flow_arn": {
        "computed": true,
        "description": "Arn",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connect_instance_id": {
        "description": "Amazon Connect Instance Id",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Campaign name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "Campaign schedule",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_time": {
              "computed": true,
              "description": "Timestamp with no UTC offset or timezone",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refresh_frequency": {
              "computed": true,
              "description": "Time duration in ISO 8601 format",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "Timestamp with no UTC offset or timezone",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source": {
        "computed": true,
        "description": "The possible source of the campaign",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_profiles_segment_arn": {
              "computed": true,
              "description": "Arn",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_trigger": {
              "computed": true,
              "description": "The event trigger of the campaign",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "customer_profiles_domain_arn": {
                    "computed": true,
                    "description": "Arn",
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
      "tags": {
        "computed": true,
        "description": "One or more tags.",
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
      }
    },
    "description": "Definition of AWS::ConnectCampaignsV2::Campaign Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectcampaignsv2CampaignSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectcampaignsv2Campaign), &result)
	return &result
}
