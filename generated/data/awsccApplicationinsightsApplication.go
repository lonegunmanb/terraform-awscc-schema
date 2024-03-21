package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationinsightsApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The ARN of the ApplicationInsights application.",
        "description_kind": "plain",
        "type": "string"
      },
      "attach_missing_permission": {
        "computed": true,
        "description": "If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing",
        "description_kind": "plain",
        "type": "bool"
      },
      "auto_configuration_enabled": {
        "computed": true,
        "description": "If set to true, application will be configured with recommended monitoring configuration.",
        "description_kind": "plain",
        "type": "bool"
      },
      "component_monitoring_settings": {
        "computed": true,
        "description": "The monitoring settings of the components.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_arn": {
              "computed": true,
              "description": "The ARN of the compnonent.",
              "description_kind": "plain",
              "type": "string"
            },
            "component_configuration_mode": {
              "computed": true,
              "description": "The component monitoring configuration mode.",
              "description_kind": "plain",
              "type": "string"
            },
            "component_name": {
              "computed": true,
              "description": "The name of the component.",
              "description_kind": "plain",
              "type": "string"
            },
            "custom_component_configuration": {
              "computed": true,
              "description": "The monitoring configuration of the component.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration_details": {
                    "computed": true,
                    "description": "The configuration settings",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_metrics": {
                          "computed": true,
                          "description": "A list of metrics to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metric_name": {
                                "computed": true,
                                "description": "The name of the metric to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "alarms": {
                          "computed": true,
                          "description": "A list of alarms to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_name": {
                                "computed": true,
                                "description": "The name of the CloudWatch alarm to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "severity": {
                                "computed": true,
                                "description": "Indicates the degree of outage when the alarm goes off.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "ha_cluster_prometheus_exporter": {
                          "computed": true,
                          "description": "The HA cluster Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "hana_prometheus_exporter": {
                          "computed": true,
                          "description": "The HANA DB Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agree_to_install_hanadb_client": {
                                "computed": true,
                                "description": "A flag which indicates agreeing to install SAP HANA DB client.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "hana_port": {
                                "computed": true,
                                "description": "The HANA DB port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "hana_secret_name": {
                                "computed": true,
                                "description": "The secret name which manages the HANA DB credentials e.g. {\n  \"username\": \"\u003c\u003e\",\n  \"password\": \"\u003c\u003e\"\n}.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "hanasid": {
                                "computed": true,
                                "description": "HANA DB SID.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "jmx_prometheus_exporter": {
                          "computed": true,
                          "description": "The JMX Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "host_port": {
                                "computed": true,
                                "description": "Java agent host port",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "jmxurl": {
                                "computed": true,
                                "description": "JMX service URL.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "logs": {
                          "computed": true,
                          "description": "A list of logs to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "encoding": {
                                "computed": true,
                                "description": "The type of encoding of the logs to be monitored.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_group_name": {
                                "computed": true,
                                "description": "The CloudWatch log group name to be associated to the monitored log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_path": {
                                "computed": true,
                                "description": "The path of the logs to be monitored.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_type": {
                                "computed": true,
                                "description": "The log type decides the log patterns against which Application Insights analyzes the log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pattern_set": {
                                "computed": true,
                                "description": "The name of the log pattern set.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "net_weaver_prometheus_exporter": {
                          "computed": true,
                          "description": "The NetWeaver Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "instance_numbers": {
                                "computed": true,
                                "description": "SAP instance numbers for ASCS, ERS, and App Servers.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "sapsid": {
                                "computed": true,
                                "description": "SAP NetWeaver SID.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "processes": {
                          "computed": true,
                          "description": "A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metrics": {
                                "computed": true,
                                "description": "A list of metrics to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "process_name": {
                                "computed": true,
                                "description": "The name of the process to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "sql_server_prometheus_exporter": {
                          "computed": true,
                          "description": "The SQL Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "sql_secret_name": {
                                "computed": true,
                                "description": "Secret name which managers SQL exporter connection. e.g. {\"data_source_name\": \"sqlserver://\u003cUSERNAME\u003e:\u003cPASSWORD\u003e@localhost:1433\"}",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "windows_events": {
                          "computed": true,
                          "description": "A list of Windows Events to log.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "event_levels": {
                                "computed": true,
                                "description": "The levels of event to log. ",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The type of Windows Events to log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_group_name": {
                                "computed": true,
                                "description": "The CloudWatch log group name to be associated to the monitored log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pattern_set": {
                                "computed": true,
                                "description": "The name of the log pattern set.",
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
                  "sub_component_type_configurations": {
                    "computed": true,
                    "description": "Sub component configurations of the component.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sub_component_configuration_details": {
                          "computed": true,
                          "description": "The configuration settings of sub components.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metrics": {
                                "computed": true,
                                "description": "A list of metrics to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "logs": {
                                "computed": true,
                                "description": "A list of logs to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encoding": {
                                      "computed": true,
                                      "description": "The type of encoding of the logs to be monitored.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_group_name": {
                                      "computed": true,
                                      "description": "The CloudWatch log group name to be associated to the monitored log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_path": {
                                      "computed": true,
                                      "description": "The path of the logs to be monitored.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_type": {
                                      "computed": true,
                                      "description": "The log type decides the log patterns against which Application Insights analyzes the log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "pattern_set": {
                                      "computed": true,
                                      "description": "The name of the log pattern set.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "processes": {
                                "computed": true,
                                "description": "A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metrics": {
                                      "computed": true,
                                      "description": "A list of metrics to monitor for the component.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "alarm_metric_name": {
                                            "computed": true,
                                            "description": "The name of the metric to be monitored for the component.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "process_name": {
                                      "computed": true,
                                      "description": "The name of the process to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "windows_events": {
                                "computed": true,
                                "description": "A list of Windows Events to log.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "event_levels": {
                                      "computed": true,
                                      "description": "The levels of event to log. ",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "event_name": {
                                      "computed": true,
                                      "description": "The type of Windows Events to log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_group_name": {
                                      "computed": true,
                                      "description": "The CloudWatch log group name to be associated to the monitored log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "pattern_set": {
                                      "computed": true,
                                      "description": "The name of the log pattern set.",
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
                        "sub_component_type": {
                          "computed": true,
                          "description": "The sub component type.",
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
            "default_overwrite_component_configuration": {
              "computed": true,
              "description": "The overwritten settings on default component monitoring configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration_details": {
                    "computed": true,
                    "description": "The configuration settings",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_metrics": {
                          "computed": true,
                          "description": "A list of metrics to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metric_name": {
                                "computed": true,
                                "description": "The name of the metric to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "alarms": {
                          "computed": true,
                          "description": "A list of alarms to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_name": {
                                "computed": true,
                                "description": "The name of the CloudWatch alarm to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "severity": {
                                "computed": true,
                                "description": "Indicates the degree of outage when the alarm goes off.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "ha_cluster_prometheus_exporter": {
                          "computed": true,
                          "description": "The HA cluster Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "hana_prometheus_exporter": {
                          "computed": true,
                          "description": "The HANA DB Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agree_to_install_hanadb_client": {
                                "computed": true,
                                "description": "A flag which indicates agreeing to install SAP HANA DB client.",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "hana_port": {
                                "computed": true,
                                "description": "The HANA DB port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "hana_secret_name": {
                                "computed": true,
                                "description": "The secret name which manages the HANA DB credentials e.g. {\n  \"username\": \"\u003c\u003e\",\n  \"password\": \"\u003c\u003e\"\n}.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "hanasid": {
                                "computed": true,
                                "description": "HANA DB SID.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "jmx_prometheus_exporter": {
                          "computed": true,
                          "description": "The JMX Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "host_port": {
                                "computed": true,
                                "description": "Java agent host port",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "jmxurl": {
                                "computed": true,
                                "description": "JMX service URL.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "logs": {
                          "computed": true,
                          "description": "A list of logs to monitor for the component.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "encoding": {
                                "computed": true,
                                "description": "The type of encoding of the logs to be monitored.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_group_name": {
                                "computed": true,
                                "description": "The CloudWatch log group name to be associated to the monitored log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_path": {
                                "computed": true,
                                "description": "The path of the logs to be monitored.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_type": {
                                "computed": true,
                                "description": "The log type decides the log patterns against which Application Insights analyzes the log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pattern_set": {
                                "computed": true,
                                "description": "The name of the log pattern set.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "net_weaver_prometheus_exporter": {
                          "computed": true,
                          "description": "The NetWeaver Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "instance_numbers": {
                                "computed": true,
                                "description": "SAP instance numbers for ASCS, ERS, and App Servers.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "sapsid": {
                                "computed": true,
                                "description": "SAP NetWeaver SID.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "processes": {
                          "computed": true,
                          "description": "A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metrics": {
                                "computed": true,
                                "description": "A list of metrics to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "process_name": {
                                "computed": true,
                                "description": "The name of the process to be monitored for the component.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "sql_server_prometheus_exporter": {
                          "computed": true,
                          "description": "The SQL Prometheus Exporter settings.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prometheus_port": {
                                "computed": true,
                                "description": "Prometheus exporter port.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "sql_secret_name": {
                                "computed": true,
                                "description": "Secret name which managers SQL exporter connection. e.g. {\"data_source_name\": \"sqlserver://\u003cUSERNAME\u003e:\u003cPASSWORD\u003e@localhost:1433\"}",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "windows_events": {
                          "computed": true,
                          "description": "A list of Windows Events to log.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "event_levels": {
                                "computed": true,
                                "description": "The levels of event to log. ",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The type of Windows Events to log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_group_name": {
                                "computed": true,
                                "description": "The CloudWatch log group name to be associated to the monitored log.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pattern_set": {
                                "computed": true,
                                "description": "The name of the log pattern set.",
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
                  "sub_component_type_configurations": {
                    "computed": true,
                    "description": "Sub component configurations of the component.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "sub_component_configuration_details": {
                          "computed": true,
                          "description": "The configuration settings of sub components.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_metrics": {
                                "computed": true,
                                "description": "A list of metrics to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "logs": {
                                "computed": true,
                                "description": "A list of logs to monitor for the component.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encoding": {
                                      "computed": true,
                                      "description": "The type of encoding of the logs to be monitored.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_group_name": {
                                      "computed": true,
                                      "description": "The CloudWatch log group name to be associated to the monitored log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_path": {
                                      "computed": true,
                                      "description": "The path of the logs to be monitored.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_type": {
                                      "computed": true,
                                      "description": "The log type decides the log patterns against which Application Insights analyzes the log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "pattern_set": {
                                      "computed": true,
                                      "description": "The name of the log pattern set.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "processes": {
                                "computed": true,
                                "description": "A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "alarm_metrics": {
                                      "computed": true,
                                      "description": "A list of metrics to monitor for the component.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "alarm_metric_name": {
                                            "computed": true,
                                            "description": "The name of the metric to be monitored for the component.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "process_name": {
                                      "computed": true,
                                      "description": "The name of the process to be monitored for the component.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "windows_events": {
                                "computed": true,
                                "description": "A list of Windows Events to log.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "event_levels": {
                                      "computed": true,
                                      "description": "The levels of event to log. ",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "event_name": {
                                      "computed": true,
                                      "description": "The type of Windows Events to log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "log_group_name": {
                                      "computed": true,
                                      "description": "The CloudWatch log group name to be associated to the monitored log.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "pattern_set": {
                                      "computed": true,
                                      "description": "The name of the log pattern set.",
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
                        "sub_component_type": {
                          "computed": true,
                          "description": "The sub component type.",
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
            "tier": {
              "computed": true,
              "description": "The tier of the application component.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "custom_components": {
        "computed": true,
        "description": "The custom grouped components.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_name": {
              "computed": true,
              "description": "The name of the component.",
              "description_kind": "plain",
              "type": "string"
            },
            "resource_list": {
              "computed": true,
              "description": "The list of resource ARNs that belong to the component.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "cwe_monitor_enabled": {
        "computed": true,
        "description": "Indicates whether Application Insights can listen to CloudWatch events for the application resources.",
        "description_kind": "plain",
        "type": "bool"
      },
      "grouping_type": {
        "computed": true,
        "description": "The grouping type of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_pattern_sets": {
        "computed": true,
        "description": "The log pattern sets.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_patterns": {
              "computed": true,
              "description": "The log patterns of a set.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pattern": {
                    "computed": true,
                    "description": "The log pattern.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "pattern_name": {
                    "computed": true,
                    "description": "The name of the log pattern.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rank": {
                    "computed": true,
                    "description": "Rank of the log pattern.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "pattern_set_name": {
              "computed": true,
              "description": "The name of the log pattern set.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "ops_center_enabled": {
        "computed": true,
        "description": "When set to true, creates opsItems for any problems detected on an application.",
        "description_kind": "plain",
        "type": "bool"
      },
      "ops_item_sns_topic_arn": {
        "computed": true,
        "description": "The SNS topic provided to Application Insights that is associated to the created opsItem.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "computed": true,
        "description": "The name of the resource group.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of Application Insights application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ApplicationInsights::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApplicationinsightsApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationinsightsApplication), &result)
	return &result
}
