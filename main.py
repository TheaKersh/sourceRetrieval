from asyncio.windows_events import NULL
from pickle import APPEND
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





# Gets a bunch of cookies for future interactions
r = s.get(url)

print(r.status_code)

# Sends the login data
r = s.post(url, data=data )

f = open("test.html", "w")
st = r.content.decode("utf-8")
f.write(st)
f2 = open("links.txt", "r")
links = []
print(f2.readline())
for line in f2:
	links.append(line[1:(len(line)-2)])
print("https://ps.seattleschools.org/guardian/" + links[0])
r = s.get("https://ps.seattleschools.org/guardian/" + links[0])

#lol
f_asses = open("assignments.txt", "a")
for link in links:
	if ">" in link:
		continue
	else:
		r = s.get("https://ps.seattleschools.org/guardian/" + link)
		f_asses.write(r.content.decode("utf-8"))

url = "https://ps.seattleschools.org/ws/xte/assignment/lookup?_=1647657691326"
postData = {"section_ids":[515612],"student_ids":[194202],"start_date":"2022-1-28","end_date":"2022-4-7"}
s.get("https://ps.seattleschools.org/scripts/components/studentScores/views/studentScores.html")
s.get("https://ps.seattleschools.org/ws/i18n/messageKeys?_=1647657541047&keys=psx.js.jswidgets.gridWidget,psx.js.scripts_components_shared_interceptors")
r = s.post(url, json=postData)

print(r.content)

