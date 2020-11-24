# idengine

The goal is an easy to use containerised identity and authentication service, mainly for IoT and mobile app developers.

### Curl examples
Your remote device or mobile app can obtain an identity dynamically, like this:
```
$ curl https://api.idthings.io/identities/new/
{be39aaa1-3ab2-4855-9b13-d1bae9410baf,03Yg@&8F0OJM*6*@MDO0}
$
```
When your remote device makes requests to your API, simply proxy authentication to the idengine service (here running at idThings):
```
$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
-H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
HTTP/1.1 200 OK

$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
-H "X-idThings-Password: wrong-password"
HTTP/1.1 401 Unauthorized

$
```
Your remote device can rotate its own password, receiving a fresh one with this request:
```
$ curl https://api.idthings.io/identities/rotate/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
-H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
{WEN*86I9t3OUq0#))D4T}

$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
-H "X-idThings-Password: WEN*86I9t3OUq0#))D4T"
HTTP/1.1 200 OK

$
```
