# Cyber Threat Intelligence API

Lightweight Cyber Threat Intelligence Platform over HTTP, built as a prototype product to be used in practise DevSecOps deployments.

The platform is Full Trust (as opposed to Zero Trust) and as long as the correct parameters are there, it will process the data given to it.

*NOTE VALUES SPECIFIED BETWEEN <> INDICATE URL PARAMETERS*

## GET ENDPOINTS
*ALL GET ENDPOINTS WILL RETURN A 200 STATUS CODE IF SUCCESFUL, ELSE 500*
| Endpoints | Parameters | Request Type | Details |
|-----------|------------|--------------|---------|
| /health   | N/A        | GET          | 200 Status Code if database connection is available, else 500.|
| /indicator/\<Indicator ID>| N/A  | GET | JSON of the specified indicator of compromise's information|
| /actor/\<Actor ID> | N/A | GET | JSON of the specified Threat Actor's information|

## POST ENDPOINTS
*ALL POST ENDPOINTS WILL RETURN A 202 ACCEPTED STATUS CODE IF SUCCESSFUL ELSE 500*  
*DATA is to be sent as form values*

/new_indicator - Add a new indicator to the Threat Intel Database.

| Parameter | Required? | Data Type | Comment |
|-----------|-----------|-----------|---------|
| type| YES |  string| The type of IOC being added, this must be one of "filehash", "ipaddress", "tactic", "cve", "email", "username", "hostname", in lower case.|
| value | YES | string | the value of the IOC being added, for example an IP address, "192.168.0.1"|
| comment | No | string | A Comment about the IOC being added, perhaps why it has been added to the Threat Intel DB |
| actor | No | integer | An integer ID (the ID the threat actor is stored as in the Threat Intel DB) of a Threat Actor, it is recommened to add any threat actors to the Database before IOCs so they can have IOCs mapped against them, otherwise will be mapped as unknown |

/new_actor - Add a new threat actor to the Threat Intel Database.

| Parameter | Required? | Data Type | Comment |
|-----------|-----------|-----------|---------|
| name | YES | string | The name of the new threat actor to be added to the database|

/new_alias - Add an Alias to a threat actor in the Threat Intel Database

| Parameter | Required? | Data Type | Comment |
|-----------|-----------|-----------|---------|
| actor | YES | integer | The Database ID for the threat actor the alias is to be added to |
| alias | YES | string | The alias to be added for the threat actor| 
