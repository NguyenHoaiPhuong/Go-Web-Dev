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

# How the JWT signature works

So if the header and payload of a JWT can be read and written to by anyone, what actually makes a JWT secure? The answer lies in how the last part (the signature) is generated.

Let’s pretend that your application wants to issue a JWT to a user (for example, user1) that has successfully signed in.

Making the header and payload are pretty straightforward: The header is more or less fixed, and the payload JSON object is formed by setting the user ID and the expiry time in unix milliseconds.

The application issuing the token will also have a key, which is a secret value, and known only to the application itself. The base64 representations of the header and payload are then combined with the secret key and then passed through a hashing algorithm (in this case its HS256, as mentioned in the header)

![How signature work](./jwt-how-signature-work.svg)

# Verifying a JWT

In order to verify an incoming JWT, a signature is once again generated using the header and payload from the incoming JWT, and the secret key. If the signature matches the one on the JWT, then the JWT is considered valid.

Now let’s pretend that you’re a hacker trying to issue a fake token. You can easily generate the header and payload, but without knowing the key, there is no way to generate a valid signature. If you try to tamper with the existing payload of a valid JWT, the signatures will no longer match.

![Verifying a JWT](./jwt-verification.svg)

In this way, the JWT acts as a way to authorize users in a secure manner, without actually storing any information (besides the key) on the application server.

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

- https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
