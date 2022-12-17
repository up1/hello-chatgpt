try to coding with go

POST /v1/register

Body Request in JSON format
{
"user": "demo",
"password": "pass"
}

Return respose with HTTP Code = 200 and JSON data
{
	"message": "Register success"
}

and Error with HTTP code = 500
{
	"message": "Internal server error"
}
