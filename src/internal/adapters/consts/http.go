package consts

import "time"

const (
	// HTTP server default values.
	HTTPServerPortDefault = "8081"
	HTTPServerHostDefault = "0.0.0.0"

	/* HTTP Server values:
	ReadTimeout - protects against slow-loris / hanging clients. Time from first byte until the entire request is read.
	ReadHeaderTimeout - prevents header-based DoS (clients sending headers at 1 byte/second).
	WriteTimeout - cuts long-running handlers (DB queries, external calls) that exceed this.
		Must be ≥ your p99 request latency + buffer.
	IdleTimeout - closes idle keep-alive connections to free file descriptors.
		Cloud load-balancers usually recycle after 60–120 s.
	MaxHeaderBytes - rejects oversized headers (HTTP smuggling, memory abuse).
		99 % of REST APIs fit in 64 KiB; 1 MiB is safe. */
	ReadTimeout       = 10 * time.Second
	ReadHeaderTimeout = 5 * time.Second
	WriteTimeout      = 30 * time.Second
	IdleTimeout       = 120 * time.Second
	MaxHeaderBytes    = 1 << 20

	// HTTP headers.
	ContentType = "Content-Type"

	// MIME types.
	ApplicationJSON = "application/json"
)
