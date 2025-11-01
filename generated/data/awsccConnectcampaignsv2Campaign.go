package data

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
        "computed": true,
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
                          "type": "string"
                        },
                        "source_email_address_display_name": {
                          "computed": true,
                          "description": "The name of the source email address display name",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "wisdom_template_arn": {
                          "computed": true,
                          "description": "Arn",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                          "type": "string"
                        },
                        "wisdom_template_arn": {
                          "computed": true,
                          "description": "Arn",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                    "type": "number"
                  },
                  "connect_queue_id": {
                    "computed": true,
                    "description": "The queue for the call",
                    "description_kind": "plain",
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
                                "type": "bool"
                              },
                              "enable_answer_machine_detection": {
                                "computed": true,
                                "description": "Flag to decided whether outbound calls should have answering machine detection enabled or not",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "connect_contact_flow_id": {
                          "computed": true,
                          "description": "The identifier of the contact flow for the outbound call",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "connect_source_phone_number": {
                          "computed": true,
                          "description": "The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
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
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "bandwidth_allocation": {
                                "computed": true,
                                "description": "The bandwidth allocation of a queue resource.",
                                "description_kind": "plain",
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
                                "type": "number"
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
            }
          },
          "nesting_mode": "single"
        }
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
                          "type": "number"
                        },
                        "max_count_per_recipient": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "unit": {
                          "computed": true,
                          "description": "The communication limit time unit",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "instance_limits_handling": {
              "computed": true,
              "description": "Enumeration of Instance Limits handling in a Campaign",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
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
                    "type": "string"
                  },
                  "local_time_zone_detection": {
                    "computed": true,
                    "description": "Local TimeZone Detection method list",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
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
                                      "type": "string"
                                    },
                                    "start_time": {
                                      "computed": true,
                                      "description": "Time in ISO 8601 format, e.g. T23:11",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "set"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
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
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of a restricted period",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "start_date": {
                                "computed": true,
                                "description": "Date in ISO 8601 format, e.g. 2024-01-01",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
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
      "connect_campaign_flow_arn": {
        "computed": true,
        "description": "Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "connect_instance_id": {
        "computed": true,
        "description": "Amazon Connect Instance Id",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Campaign name",
        "description_kind": "plain",
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
              "type": "string"
            },
            "refresh_frequency": {
              "computed": true,
              "description": "Time duration in ISO 8601 format",
              "description_kind": "plain",
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "Timestamp with no UTC offset or timezone",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::ConnectCampaignsV2::Campaign",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectcampaignsv2CampaignSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectcampaignsv2Campaign), &result)
	return &result
}
