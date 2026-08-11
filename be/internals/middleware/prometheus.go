package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	requestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Histogram of request duration in seconds",
			Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "endpoint"},
	)

	requestsInProgress = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Number of HTTP requests currently in progress",
		},
		[]string{"method", "endpoint"},
	)

	connectedDevices = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "connected_devices",
			Help: "Number of currently connected devices, tracked by unique user+device combinations with active tokens",
		},
		[]string{"user_id", "role", "device"},
	)

	loginCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "login_total",
			Help: "Total number of login attempts",
		},
		[]string{"status"},
	)
)

type deviceSession struct {
	UserID    string
	Role      string
	Device    string
	LastSeen  time.Time
	ExpiresAt time.Time
}

var (
	deviceStore   = make(map[string]deviceSession)
	deviceStoreMu sync.RWMutex
)

func TrackConnectedDevice(userID, role, device string, expiresAt time.Time) {
	deviceStoreMu.Lock()
	defer deviceStoreMu.Unlock()

	key := userID + ":" + device
	deviceStore[key] = deviceSession{
		UserID:    userID,
		Role:      role,
		Device:    device,
		LastSeen:  time.Now(),
		ExpiresAt: expiresAt,
	}
}

func RecordLoginAttempt(success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	loginCounter.WithLabelValues(status).Inc()
}

func cleanupExpiredDevices() {
	now := time.Now()
	for key, session := range deviceStore {
		if now.After(session.ExpiresAt) {
			connectedDevices.WithLabelValues(session.UserID, session.Role, session.Device).Set(0)
			delete(deviceStore, key)
		}
	}
}

func init() {
	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "build_info",
			Help: "Build information with version as value and labels for details",
			ConstLabels: prometheus.Labels{
				"version": "1.0.0",
				"module":  "be",
			},
		},
		func() float64 { return 1 },
	)

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			deviceStoreMu.Lock()
			cleanupExpiredDevices()
			deviceStoreMu.Unlock()
		}
	}()
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		requestsInProgress.WithLabelValues(method, path).Inc()

		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		requestCounter.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
		requestDuration.WithLabelValues(method, path).Observe(duration)
		requestsInProgress.WithLabelValues(method, path).Dec()

		if userID, exists := c.Get("userID"); exists {
			if role, roleExists := c.Get("role"); roleExists {
				device := c.Request.UserAgent()
				uid, _ := userID.(string)
				r, _ := role.(string)

				deviceStoreMu.Lock()
				key := uid + ":" + device
				session, found := deviceStore[key]
				if found {
					session.LastSeen = time.Now()
					deviceStore[key] = session
				}
				deviceStoreMu.Unlock()

				if found {
					connectedDevices.WithLabelValues(uid, r, device).Set(1)
				}
			}
		}
	}
}

func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
