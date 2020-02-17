# Gin Authentication Middleware

## Introduction

This example application has three types of users: basic, subscriber and admin. While it doesn’t have any special functionality, your actual business logic could have multiple user type definitions and various restrictions on endpoints based on what kind of user is hitting the endpoint. For example, a music application (like Spotify) could restrict download endpoints to paid subscriber groups only while playback could be available to everyone.

My simple application doesn’t have music playback or anything fancy though. It just echoes a message sent through one of the endpoints back to the user with admin and subscriber users getting their username and authentication type back in the message. Note that an admin can’t hit subscriber endpoints and vice-versa. Any other user can’t hit any of those endpoints. There’s a POST endpoint (which doesn’t do anything different) but can be hit by both admin and subscriber users. The users specify their authentication type during login.

> Note that this isn’t supposed to represent your business logic and the user authentication type should always be generated on the backend itself. You can store it in an encrypted session on client side for sure but it’s 2018 and storing claims in JSON web tokens is a thing and is the recommended way to do it.

## Router groups

Assuming you have already initialized your gin router, you can define the groups as follow:

```
api := router.Group("/api/v1")
// no authentication endpoints
{
	api.POST("/login", loginHandler)
	api.GET("/message/:msg", noAuthMessageHandler)
}
// basic authentication endpoints
{
	basicAuth := api.Group("/")
	basicAuth.Use(AuthenticationRequired())
	{
		basicAuth.GET("/logout", logoutHandler)
	}
}
// admin authentication endpoints
{
	adminAuth := api.Group("/admin")
	adminAuth.Use(AuthenticationRequired("admin"))
	{
		adminAuth.GET("/message/:msg", adminMessageHandler)
	}
}
// subscriber authentication endpoints
{
	subscriberAuth := api.Group("/")
	subscriberAuth.Use(AuthenticationRequired("subscriber"))
	{
		subscriberAuth.GET("/subscriber/message/:msg", subscriberMessageHandler)
	}
}

apiV2 := router.Group("/api/v2")
router.Use(AuthenticationRequired("admin", "subscriber"))
// admin and subscriber authentication endpoints
{
	apiV2.POST("/post/message/:msg", postMessageHandler)
}
```

Note how the groups themselves can have sub-groups with the relative link (to the parent group) and that’s what we use to apply middleware to the group specifically (the `AuthenticationRequired` function, more on that later). You can have sub-sub-groups too but it might mean a middleware encountered earlier aborting the request there on a user being unauthenticated etc. so it’s not the best practice. Also, using `routerGroup.METHOD().Use(middleware)`to define middleware on a route led to the middleware being applied on the whole group so not sure if that’s a good practice too.

## Authentication Required

```
package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		if len(auths) != 0 {
			authType := session.Get("authType")
			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
				c.Abort()
				return
			}
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
```

This is where you actually check the user’s claim from the session/JWT and decide if you will let the user access the endpoint or return a `StatusUnauthorized` (if user is not logged in) or `StatusForbidden` (invalid claim). The function itself returns a gin.HandlerFunc which is basically the middleware which would be called wherever you used it (`router.Use(AuthenticationRequired())`) before the request handler. The parameters define what the valid claims are and after extracting the user’s claim, you have to verify if the user has one of the valid claims.

Note that we call `c.Abort()` if the user is unauthenticated/unauthorized. This is because gin calls the next function in the chain even after you write the header (`c.JSON()`) using `c.Next()`. The next handler is likely your route handler which might try to rewrite the header which would result in an error. That also implies the last call to `c.Next()` is unnecessary.

You can also add another check to ensure the user is in your database (if somehow the secret key for the particular JWT/session) is leaked/found and the user is trying a malicious dictionary attack to find other valid user IDs/emails signed with that key (a highly unlikely made-up scenario). Use redis or something similar to cache the user check calls to the database in that case since this middleware might be called frequently.

## Test the Application

### Login with an admin user

```
POST localhost:9000/api/v1/login

body:
{
	"username": "akagi",
	"authType": "admin"
}

result:
{
    "message": "authentication successful"
}
```

### Login with a subcriber user

```
POST localhost:9000/api/v1/login

body:
{
	"username": "yushin",
	"authType": "subscriber"
}

result:
{
    "message": "authentication successful"
}
```

### Login with a normal user

```
POST localhost:9000/api/v1/login

body:
{
	"username": "mogami",
	"authType": "user"
}

result:
{
    "message": "authentication successful"
}
```

### Login with a user with empty username

```
POST localhost:9000/api/v1/login

body:
{
	"username": " ",
	"authType": "user"
}

result:
{
    "error": "username can't be empty"
}
```

### Logout

```
GET localhost:9000/api/v1/logout

result:
{
    "message": "successfully logged out"
}
```

### No auth message

```
GET localhost:9000/api/v1/message/hahaha

result:
{
    "message": "hahaha"
}
```

### Admin message

```
GET localhost:9000/api/v1/admin/message/hihihi

result without admin login:
{
    "error": "invalid request, restricted endpoint"
}
or
{
    "error": "user needs to be signed in to access this service"
}

result with admin login:
{
    "message": "Hello Admin akagi, here's your message: hihihi"
}
```

### Subscriber message

```
GET localhost:9000/api/v1/subscriber/message/hohoho

result without subscriber login:
{
    "error": "invalid request, restricted endpoint"
}
or
{
    "error": "user needs to be signed in to access this service"
}

result with subscriber login:
{
    "message": "Hello Subscriber yushin, here's your message: hohoho"
}
```

### Admin and Subscriber message

```
POST localhost:9000/api/v2/post/message/hahaha

result without admin/subscriber login:
{
    "error": "user needs to be signed in to access this service"
}
or
{
    "error": "invalid request, restricted endpoint"
}

result with admin/subscriber login:
{
    "message": "Hello Admin/Subscriber akagi, your message: hahaha will be posted"
}
```