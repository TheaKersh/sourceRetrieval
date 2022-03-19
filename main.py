from asyncio.windows_events import NULL
from tkinter import N
from requests import Request, Session

s = Session()

url = "https://ps.seattleschools.org/guardian/home.html"

data = {
	"dbpw": "donaldtrump",
	"translator_username": "",
	"translator_password": "",
	"translator_ldappassword": "",
	"returnUrl": "https://ps.seattleschools.org/guardian/home.html",
	"serviceName": "PS+Parent+Portal",
	"serviceTicket": "",
	"pcasServerUrl": "/",
	"credentialType": "User+Id+and+Password+Credential",
	"ldappassword": "donaldtrump",
	"account": "1tswaldron",
	"pw": "donaldtrump",
	"translatorpw": ""
}

def readLine(readfrom):
	ret = ""
	for i, char in range(readfrom):
		if char == "\n":
			break
		else:
			ret.join(char)
	return ret





# Gets a bunch of cookies for future interactions
r = s.get(url)

print(r.status_code)

# Sends the login data
r = s.post( url, data=data )

f = open("test.html", "w")
str = r.content.decode("utf-8")
f.write(str)
f2 = open("links.txt", "r")
s = f2.read()
print(s)
links = [""]
for i in range(len(s)):
	if s[i] == "\n":
		break
	else :
		links[0] = links[0].join(s[i])
print(links[0])
	