### Project Goals
* an easy to deploy identity and authentication infrastructure service
* pragmatic approach focused on improving many rather than perfecting the few
* develop a simple yet robust set of methods

#### Get an identity: /identities/new/
Your remote device or mobile app can obtain an identity dynamically, like this:
```
$ curl https://api.idthings.io/identities/new/
{be39aaa1-3ab2-4855-9b13-d1bae9410baf,03Yg@&8F0OJM*6*@MDO0}
$
```
#### Authenticate a request: /identities/&lt;guid&gt;
Your remote devices will send requests to your own API, such as a data ping of a temperature monitor.
The device sends a header with either their password, or a computed HMAC digest.
Your API simply proxies this header in an authentiction request to the idengine service.

Here are simple password authentication examples, using the idengine service at idThings.io:
```
$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
HTTP/1.1 200 OK

$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: wrong-password"
HTTP/1.1 401 Unauthorized

$
```
#### Rotate secrets: /identities/rotate/&lt;guid&gt;
Your remote device can rotate its own password, receiving a fresh one with this request:
```
$ curl https://api.idthings.io/identities/rotate/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
{WEN*86I9t3OUq0#))D4T}

$
```
