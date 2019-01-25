package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
)

type key int
const requestIDKey key = 0

func generateRandomID() string {
	return "users-" + strconv.Itoa(rand.Intn(100))
}

// Pull the context from the X-REQUEST-ID and store it in the context
func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = generateRandomID()
	}

	return context.WithValue(ctx, requestIDKey, reqID)
}

// Pull the request ID out of the context
func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

/* 	Middleware to put a context with a request ID into the request.
	It is possible to make specific requests have unique context per route since we can attack Middleware
	on any number of routes we want. We can generalize it into a group of routes that take Use a specific middleware
*/
func RequestIDMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}