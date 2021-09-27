# go-ip-mmdb package

This Go package is used to enrich GeoIP2 Maxmind DB, followed [this tutorial](https://blog.maxmind.com/2020/09/01/enriching-mmdb-files-with-your-own-data-using-go/) and [this repository](https://github.com/maxmind/mmdb-from-go-blogpost)

Data struct is followed `geoip2-golang`'s [reader](https://github.com/oschwald/geoip2-golang/blob/main/reader.go)

## Requirements

1. Go 1.16 or later.
1. [`mmdbinspect`](https://github.com/maxmind/mmdbinspect)

## Using `mmdbinspect`

City: `mmdbinspect -db GeoIP2-City.mmdb 8.8.8.8`

```json
[
    {
        "Database": "GeoIP2-City_20210122/GeoIP2-City.mmdb",
        "Records": [
            {
                "Network": "116.108.94.53/24",
                "Record": {
                    "city": {
                        "geoname_id": 1566083,
                        "names": {
                            "de": "Ho-Chi-Minh-Stadt",
                            "en": "Ho Chi Minh City",
                            "es": "Ciudad Ho Chi Minh",
                            "fr": "HCMV",
                            "ja": "ホーチミン市",
                            "pt-BR": "Cidade de Ho Chi Minh",
                            "ru": "Хошимин",
                            "zh-CN": "胡志明市"
                        }
                    },
                    "continent": {
                        "code": "AS",
                        "geoname_id": 6255147,
                        "names": {
                            "de": "Asien",
                            "en": "Asia",
                            "es": "Asia",
                            "fr": "Asie",
                            "ja": "アジア",
                            "pt-BR": "Ásia",
                            "ru": "Азия",
                            "zh-CN": "亚洲"
                        }
                    },
                    "country": {
                        "geoname_id": 1562822,
                        "iso_code": "VN",
                        "names": {
                            "de": "Vietnam",
                            "en": "Vietnam",
                            "es": "Vietnam",
                            "fr": "Vietnam",
                            "ja": "ベトナム",
                            "pt-BR": "Vietnã",
                            "ru": "Вьетнам",
                            "zh-CN": "越南"
                        }
                    },
                    "location": {
                        "accuracy_radius": 10,
                        "latitude": 10.8104,
                        "longitude": 106.6444,
                        "time_zone": "Asia/Ho_Chi_Minh"
                    },
                    "registered_country": {
                        "geoname_id": 1562822,
                        "iso_code": "VN",
                        "names": {
                            "de": "Vietnam",
                            "en": "Vietnam",
                            "es": "Vietnam",
                            "fr": "Vietnam",
                            "ja": "ベトナム",
                            "pt-BR": "Vietnã",
                            "ru": "Вьетнам",
                            "zh-CN": "越南"
                        }
                    },
                    "subdivisions": [
                        {
                            "geoname_id": 1580578,
                            "iso_code": "SG",
                            "names": {
                                "de": "Ho-Chi-Minh-Stadt",
                                "en": "Ho Chi Minh",
                                "es": "Ciudad Ho Chi Minh",
                                "fr": "Saigon",
                                "ja": "ホーチミン市",
                                "pt-BR": "Cidade de Ho Chi Minh",
                                "ru": "Хо Ши Мин",
                                "zh-CN": "胡志明市"
                            }
                        }
                    ]
                }
            }
        ],
        "Lookup": "116.108.94.53"
    }
]
```

ISP: `mmdbinspect -db GeoIP2-ISP.mmdb 8.8.8.8`

```json
[
    {
        "Database": "GeoIP2-ISP.mmdb",
        "Records": [
            {
                "Network": "8.8.8.8/32",
                "Record": {
                    "autonomous_system_number": 15169,
                    "autonomous_system_organization": "GOOGLE",
                    "isp": "Google",
                    "organization": "Google"
                }
            }
        ],
        "Lookup": "8.8.8.8"
    }
]
```