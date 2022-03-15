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


# Gets a bunch of cookies for future interactions
r = s.get(url)

print(r.status_code)

# Sends the login data
r = s.post( url, data=data )

print(r.content)
f = open("test.html", "a")
str = r.content.decode("utf-8")
f.write(str)

