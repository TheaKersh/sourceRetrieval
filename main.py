from asyncio.windows_events import NULL
from pickle import APPEND
from tkinter import N
from requests import Request, Session
import time
import sched
import constant
sche = sched.scheduler(time.time, time.sleep)
print("a")
def check_Grades(sc):
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
	print("exec")

	# Gets a bunch of cookies for future interactions
	r = s.get(url)

	print(r.status_code)

	# Sends the login data
	r = s.post(url, data=data )

	f = open("test.html", "w")
	st = r.content.decode("utf-8")
	f.write(st)
	sche.enter(3600, 1, check_Grades, (sc,))
sche.enter(0, 1, check_Grades, (sche,))
sche.run()


