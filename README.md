### idEngine
The idEngine came about when I was developing an IoT product, and needed an way to deploy millions of identities (user id, secret) to devices.
Prototyping with the Arduino Nano 33 IoT, it was clear there are significant logical security challenges in deploying credentials.

There is almost nothing secret on an IoT device, so factory installed credentials seems like fools gold.
The common case where a device is reset also raises the question, of whether a secret can be rotated safely?

In addition, there are significant scaling challenges to running a credential service running into the millions of accounts.
The keywords here, from an old infrstructure guy, are: simplicity, robustness, automation.

My approach is to treat an IoT device, and to a lesser extent a mobile app install, as simple a receptical that can hold a credential.
However, until a credential is deployed that device has no identity.
Resetting the device turns it back into a mere receptical.

### Project Goals
* an easy to deploy identity and authentication infrastructure service
* encourage dynamic credential automation for applications and infrastructure
* develop a simple and understandable implementation

### Audience
* IoT hackers and makers, idEngine easily scales to millions of credentials (or use idThings.io)
* Developers in a micro-service environment, or building mobile apps with minimal backend APIs
* DevOps folk who need to manage storage and services containing credentials
* People interested in improving the traditional user/password key-pair approach

---

### API Summary
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

Here are simple password authentication examples, using the idEngine service at idThings.io:
```
$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
HTTP/1.1 200 OK

$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: wrong-password"
HTTP/1.1 401 Unauthorized

$
```
For sentient devices it's a very short step to using HMAC digests to sign http requests to your API.
This means the device secret isn't transmitted with every call.
```
$ curl -I https://api.idthings.io/identities/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Digest: HMAC-SHA256,c7fc567324b236e...,1604573826351,my device data"
HTTP/1.1 200 OK
```
#### Rotate secrets: /identities/rotate/&lt;guid&gt;
Your remote device can rotate its own password, receiving a fresh one with this request:
```
$ curl https://api.idthings.io/identities/rotate/be39aaa1-3ab2-4855-9b13-d1bae9410baf \
    -H "X-idThings-Password: 03Yg@&8F0OJM*6*@MDO0"
{WEN*86I9t3OUq0#))D4T}

$
```
