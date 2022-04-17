#Middleware

HTTP middleware adaptors and utilities based on Mat Ryer's post [Writing middleware in #golang and how Go makes it so much fun.](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81)

#Types/Interfaces

- Adapter interface : func(h http.Handler) http.Handler
- HttpLog interface : contains 2 functions. Request() to log the request and Response to log the response
- LogResponseWriter : internal struct to hold the http writer and record the status changes

#Functions

- Adapt(h http.Handler, as ...Adapter) http.Handler
- AdaptArray(h http.Handler, a []Adapter) http.Handler

#MiddlewareAdapters

- EmbedRequestId : takes a requestid generator and embeds a requestId in the request's context
- EmbedGormDb : takes a [GORM](https://gorm.io/index.html) db and embeds it in the request's context
- LimitToMethods : takes one or more http verbs and limits requests to those methods
- CheckNotFound : usually used on the last "/" route to ensure a 404 response on unknown URI paths. will force a 404 response if the request path does not match the supplied path

#Notes:

- as this follows the basic interface requirements as defined in the interface Adapter. Any similarly intefaced adapters may be used such as [cloudlena/adapters](https://github.com/cloudlena/adapters) and Gorilla middleware
- external middleware does not necessarily embody the same logging
