# INTRODUCTION

A JWT technically is a mechanism to verify the owner of some JSON data. It’s an encoded string, which is URL safe, that can contain an unlimited amount of data (unlike a cookie), and it’s cryptographically signed.

When a server receives a JWT, it can guarantee the data it contains can be trusted because it’s signed by the source. No middleman can modify a JWT once it’s sent.

It’s important to note that a JWT guarantees data ownership but not encryption; the JSON data you store into a JWT can be seen by anyone that intercepts the token, as it’s just serialized, not encrypted. For this reason, it’s highly recommended to use HTTPS with JWTs (and HTTPS in general, by the way).

# JWT STRUCTURE

JWT is composed of 3 parts which are separated by the dot sign "."

* Header
* Payload
* Signature

## Header

Header is representative for data type and the algorithm being used.

```
{
    "typ": "JWT",
    "alg": "HS256"
}
```

* “typ” (type) : datatype is JWT
* “alg” (algorithm) : encoding algorithm is HS256

## Payload

Payload will transfer the data which users want to "convey", such as username, userID, author, etc.

```
{
  "user_name": "admin",
  "user_id": "1513717410",
  "authorities": "ADMIN_USER",
  "jti": "474cb37f-2c9c-44e4-8f5c-1ea5e4cc4d18"
}
```

**Note:**

- Don't put too much data in the payload because it might affect the delay of server response to a long Token.

## Signature

Signature is created by the encryption of the **header**, **payload** and a **secret** string.

```
data = base64urlEncode( header ) + "." + base64urlEncode( payload )
signature = Hash( data, secret );
```

base64UrlEncoder: algorithm used to encode header and payload.

After encoding the header and payload, we will have a string like below:

```
// header
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
// payload
eyJhdWQiOlsidGVzdGp3dHJlc291cmNlaWQiXSwidXNlcl9uYW1lIjoiYWRtaW4iLCJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNTEzNzE
```

Then, encrypting above 2 strings together with the secret string will generate **signature** string as below:

```
9nRhBWiRoryc8fV5xRpTmw9iyJ6EM7WTGTjvCM1e36Q
```

Finally, combine 3 strings above, we will have a complete JWT token:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsidGVzdGp3dHJlc291cmNlaWQiXSwidXNlcl9uYW1lIjoiYWRtaW4iLCJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNTEzNzE.9nRhBWiRoryc8fV5xRpTmw9iyJ6EM7WTGTjvCM1e36Q
```

# When to use JWT

JWT is a great technology for API authentication and server-to-server authorization.

It’s not a good choice for sessions.

- Authentication

- Exchange info

## Using JWT for API authentication

A very common use of a JWT token, and the one you should probably only use JWT for, is as an API authentication mechanism.

![Secret token](./secret_token.png)

# Reference pages

- https://topdev.vn/blog/json-web-token-la-gi/

- https://topdev.vn/blog/authentication-nang-cao-trong-single-page-application-spa-reactvue-dung-jwt-ket-hop-cookie/

- https://topdev.vn/blog/tai-sao-phai-su-dung-json-web-token-jwt-de-bao-mat-api/

- https://blog.logrocket.com/jwt-authentication-best-practices/

- https://topdev.vn/blog/huong-dan-authorization-voi-jwt/

- https://levelup.gitconnected.com/using-jwt-in-your-react-redux-app-for-authorization-d31be51a50d2
