package middleware

import "github.com/go-chi/chi/middleware"

var Logging = middleware.Logger

// func Logging(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
//
// 		// Pass to the next handler
// 		next.ServeHTTP(w, r)
//
// 		duration := time.Since(start)
// 		logger.Sugar.Infow("request",
// 			"method", r.Method,
// 			"path", r.URL.Path,
// 			"duration", duration,
// 			"remote", r.RemoteAddr,
// 			"user_agent", r.UserAgent(),
// 		)
// 	})
// }
