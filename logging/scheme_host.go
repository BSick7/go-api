package logging

import (
	"net/http"
	"strings"
)

// ExtractScheme returns the protocol used by the client
// This considers reverse proxy headers and falls back to request characteristics
func ExtractScheme(r *http.Request) string {
	// Check for WebSocket upgrade request
	if r.Header.Get("Connection") == "Upgrade" && r.Header.Get("Upgrade") == "websocket" {
		if r.TLS != nil {
			return "wss"
		}
		return "ws"
	}

	// Prefer Forwarded if available
	if fwd := r.Header.Get("Forwarded"); fwd != "" {
		for _, part := range strings.Split(fwd, ";") {
			if strings.HasPrefix(strings.TrimSpace(part), "proto=") {
				return strings.TrimPrefix(strings.TrimSpace(part), "proto=")
			}
		}
	}

	// Fallback to X-Forwarded-Proto
	if xfproto := r.Header.Get("X-Forwarded-Proto"); xfproto != "" {
		return xfproto
	}

	// Fallback: detect based on TLS (useful if directly exposed)
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func ExtractHost(r *http.Request) string {
	// 1. Standard Forwarded header (RFC 7239)
	if forwarded := r.Header.Get("Forwarded"); forwarded != "" {
		parts := strings.Split(forwarded, ";")
		for _, part := range parts {
			if strings.HasPrefix(strings.TrimSpace(part), "host=") {
				return strings.TrimPrefix(strings.TrimSpace(part), "host=")
			}
		}
	}

	// 2. X-Forwarded-Host (de facto standard)
	if xfHost := r.Header.Get("X-Forwarded-Host"); xfHost != "" {
		// Some proxies may append multiple hosts: "client, proxy1, proxy2"
		// We take the first one, which is usually the client
		return strings.TrimSpace(strings.Split(xfHost, ",")[0])
	}

	// 3. Fallback to r.Host (value from the Host header)
	return r.Host
}
