### idEngine
The idEngine came about when I was developing an IoT product, and needed an way to deploy millions of identities (user id, secret) to devices.
Prototyping with the Arduino Nano 33 IoT, it was clear there are significant logical security challenges in deploying credentials.

There is almost nothing secret on an IoT device, so factory installed credentials seems like the wrong approach.
The common case where a device is reset also raises the question, of whether a secret can be rotated safely?

In addition, there are significant scaling challenges to running a credential service running into the millions of accounts.
The keywords here, from an old infrastructure guy, are: simplicity, robustness, automation.

My approach is to treat an IoT device, and to a lesser extent a mobile app install, as simple a receptacle that can hold a credential.
However, until a credential is deployed that device has no identity.
Resetting the device turns it back into a mere receptacle.

### Project Goals
* an easy to deploy identity and authentication container service
* encourage dynamic credential automation for applications and infrastructure
* develop a simple and understandable implementation

### Audience
* IoT hackers and makers, who needs easy to test and deploy credentials
* Developers in a micro-service environment, or building mobile apps with minimal backend APIs
* DevOps folk who need to manage storage and services containing credentials
* People interested in improving on the traditional user/password key-pair approach

---

### Quick Start
To get up and running quickly, use docker-compose to deploy idengine and redis containers.
```
$ cat docker-compose.yml
version: '3'
services:
    idengine:
        image: thisdougb/idengine:latest
        ports:
            - "8000:8000"
    redis:
        image: redis:alpine
        ports:
            - "6379:6379"
```
Bring up the containers:
```
$ docker-compose up -d
Creating network "root_default" with the default driver
Pulling idengine (thisdougb/idengine:latest)...
latest: Pulling from thisdougb/idengine
188c0c94c7c5: Pull complete
0ef7d3d256c8: Pull complete
eb1a45c31d19: Pull complete
53441b811652: Pull complete
Pulling redis (redis:alpine)...
alpine: Pulling from library/redis
67b3665cee45: Pull complete
0705890dd1f7: Pull complete
Creating root_idengine_1 ... done
Creating root_redis_1    ... done
```
And quick test:
```
$ curl localhost:8000/identities/new/
{3bbaa59a-a78c-4665-8b86-c7f31276bd83,D4K80A2NN@C@YjQ%66*n}
```

---

### API Summary
#### Get an identity: /identities/new/
Your remote device or mobile app can obtain an identity, like this:
```
$ curl https://api.idthings.io/identities/new/
{be39aaa1-3ab2-4855-9b13-d1bae9410baf,03Yg@&8F0OJM*6*@MDO0}
$
```
#### Authenticate a request: /identities/&lt;guid&gt;
When remote devices send requests to your own API, they include an auth header (password or digest).
Your API simply proxies this header to the idengine service for authentication.

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
{be39aaa1-3ab2-4855-9b13-d1bae9410baf,WEN*86I9t3OUq0#))D4T}

$
```
#### Output formatting
The default format is intended to be easy to consume on lower powered IoT devices, running C-type languages.
Typically these languages are string-challenged, so we try to make it as easy as possible.
Search your response stream for curly braces, and that's your data.

However, when requesting new identities or rotating secrets, you can a;sp specify the response in JSON format.
If that's what's easiest for your code.
```
curl "https://api.idthings.io/identities/new/?format=json"
{"id":"18896661-e861-47a2-b724-629a07a4c67d","secret":"#*P3ZO9F941L4C&L#s%C"}

curl https://api.idthings.io/identities/rotate/18896661-e861-47a2-b724-629a07a4c67d?format=json \
    -H "x-idthings-password: #*P3ZO9F941L4C&L#s%C"
{"id":"18896661-e861-47a2-b724-629a07a4c67d","secret":"m3GH7X5KCC#)0i(&CaIO"}
```
