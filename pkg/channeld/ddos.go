package channeld

import (
	"sync"
	"time"

	"github.com/channeldorg/channeld/pkg/channeldpb"
	"go.uber.org/zap"
)

var unauthenticatedConnections sync.Map

var failedAuthCounters = make(map[string]int)
var ipBlacklist = make(map[string]time.Time)
var pitBlacklist = make(map[string]time.Time)

func InitAntiDDoS() {
	Event_AuthComplete.Listen(func(data AuthEventData) {
		if data.Connection.GetConnectionType() == channeldpb.ConnectionType_SERVER {
			return
		}

		if data.AuthResult == channeldpb.AuthResultMessage_INVALID_LT {
			// Invalid access token - record the PIT
			failedAuthCounters[data.PlayerIdentifierToken]++
			if GlobalSettings.MaxFailedAuthAttempts > 0 && failedAuthCounters[data.PlayerIdentifierToken] >= GlobalSettings.MaxFailedAuthAttempts {
				pitBlacklist[data.PlayerIdentifierToken] = time.Now()
				securityLogger.Info("blacklisted PIT due to too many failed auth attempts", zap.String("pit", data.PlayerIdentifierToken))
				data.Connection.Close(nil)
			}
		} else if data.AuthResult == channeldpb.AuthResultMessage_INVALID_PIT {
			// Invalid username token - record the IP
			addr := data.Connection.RemoteAddr()
			if addr == nil {
				return
			}
			ip := GetIP(addr)
			failedAuthCounters[ip]++
			if GlobalSettings.MaxFailedAuthAttempts > 0 && failedAuthCounters[ip] >= GlobalSettings.MaxFailedAuthAttempts {
				ipBlacklist[ip] = time.Now()
				securityLogger.Info("blacklisted IP due to too many failed auth attempts", zap.String("ip", ip))
				data.Connection.Close(nil)
			}
		}
	})

	// FSM disallowed - record the PIT
	Event_FsmDisallowed.Listen(func(c *Connection) {
		if c.GetConnectionType() == channeldpb.ConnectionType_SERVER {
			return
		}

		c.fsmDisallowedCounter++
		if GlobalSettings.MaxFsmDisallowed > 0 && c.fsmDisallowedCounter >= GlobalSettings.MaxFsmDisallowed {
			pitBlacklist[c.pit] = time.Now()
			securityLogger.Info("blacklisted PIT due to too many FSM disallowed", zap.String("pit", c.pit))
			c.Close(nil)
		}
	})

	go checkUnauthConns()
}

// Disconnection unauthenticated connections after ConnectionAuthTimeoutMs.
func checkUnauthConns() {
	for {
		unauthenticatedConnections.Range(func(_, v interface{}) bool {
			conn := v.(*Connection)
			if conn.IsClosing() {
				return true
			}
			if conn.state == ConnectionState_UNAUTHENTICATED && time.Since(conn.connTime).Milliseconds() >= GlobalSettings.ConnectionAuthTimeoutMs {
				ipBlacklist[GetIP(conn.RemoteAddr())] = time.Now()
				conn.Close(nil)
				securityLogger.Info("closed and blacklisted unauthenticated connection due to timeout", zap.String("ip", conn.conn.RemoteAddr().String()))
			}
			return true
		})

		time.Sleep(time.Millisecond * 500)
	}
}
