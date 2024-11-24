package middleware

import (
    "net/http"
    "github.com/sirupsen/logrus"
)

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        next.ServeHTTP(w, r)
    })
}

func RequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logrus.Infof("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
