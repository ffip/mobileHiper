package mobile

// The above type represents a configuration structure with various fields for different settings.
// @property {Sync} Sync - The `Sync` property is of type `Sync` and is used for configuring
// synchronization settings.
// @property {PKI} PKI - PKI stands for Public Key Infrastructure. It is a set of roles, policies, and
// procedures needed to create, manage, distribute, use, store, and revoke digital certificates and
// manage public-key encryption. In the context of the `config` struct, the `PKI` property represents
// the configuration
// @property Points - The `Points` property is a map where the keys are strings and the values are
// slices of strings. It is used to store a collection of points, where each point is represented by a
// string.
// @property {Tower} Tower - The `Tower` property represents the configuration for the tower component.
// It may include settings such as the tower's address, port, and authentication credentials.
// @property {Listen} Listen - The `Listen` property is a configuration for specifying the listening
// address and port for the server. It typically includes the following fields:
// @property {Punchy} Punchy - The `Punchy` property is a struct that contains configuration options
// related to the Punchy service. It may include settings such as the Punchy server address, port, and
// other relevant parameters.
// @property {SSHD} SSHD - The `SSHD` property represents the configuration for the SSH daemon. It
// includes settings such as the SSH port, allowed users, and authentication methods.
// @property {Proxy} Proxy - The `Proxy` property in the `config` struct represents the configuration
// for the proxy settings. It includes settings such as the proxy type, address, and authentication
// credentials.
// @property {Tun} Tun - The `Tun` property in the `config` struct represents the configuration for the
// tunnel interface. It includes settings such as the IP address, subnet mask, and MTU (Maximum
// Transmission Unit) for the tunnel interface.
// @property {Logging} Logging - The `Logging` property is a struct that contains configuration options
// for logging. It may include properties such as log level, log file path, log rotation settings, etc.
// @property {Stats} Stats - The `Stats` property is a struct that contains configuration options
// related to statistics and monitoring. It may include properties such as:
// @property {Handshakes} Handshakes - The `Handshakes` property is a struct that contains
// configuration options related to handshakes. It may include properties such as `Timeout`,
// `MaxRetries`, `RetryInterval`, etc. These properties define how handshakes are handled in the
// application.
// @property {Timers} Timers - The `Timers` property is a struct that contains various timer
// configurations. It is used to define the timing settings for different operations within the
// application.
// @property {PSK} PSK - The `PSK` property in the `config` struct represents the Pre-Shared Key
// configuration. It is used for authentication and encryption purposes in a network communication. The
// `PSK` struct may contain additional fields that define the specific configuration for the Pre-Shared
// Key.
// @property {Firewall} Firewall - The `Firewall` property is a configuration for firewall settings. It
// specifies rules and settings related to network traffic filtering and security.
// @property {string} Cipher - The `Cipher` property is a string that specifies the encryption cipher
// to be used. It is used to encrypt and decrypt data during communication.
type config struct {
	Sync       Sync                `json:"sync,omitempty" yaml:"sync,omitempty"`
	PKI        PKI                 `json:"pki,omitempty" yaml:"pki,omitempty"`
	Points     map[string][]string `json:"points,omitempty" yaml:"points,omitempty"`
	Tower      Tower               `json:"tower,omitempty" yaml:"tower,omitempty"`
	Listen     Listen              `json:"listen,omitempty" yaml:"listen,omitempty"`
	Punchy     Punchy              `json:"punchy,omitempty" yaml:"punchy,omitempty"`
	SSHD       SSHD                `json:"sshd,omitempty" yaml:"sshd,omitempty"`
	Proxy      Proxy               `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	Tun        Tun                 `json:"tun,omitempty" yaml:"tun,omitempty"`
	Logging    Logging             `json:"logging,omitempty" yaml:"logging,omitempty"`
	Stats      Stats               `json:"stats,omitempty" yaml:"stats,omitempty"`
	Handshakes Handshakes          `json:"handshakes,omitempty" yaml:"handshakes,omitempty"`
	Timers     Timers              `json:"timers,omitempty" yaml:"timers,omitempty"`
	PSK        PSK                 `json:"psk,omitempty" yaml:"psk,omitempty"`
	Firewall   Firewall            `json:"firewall,omitempty" yaml:"firewall,omitempty"`
	Cipher     string              `json:"cipher,omitempty" yaml:"cipher,omitempty"`
}

// The Sync type represents a synchronization configuration with various properties.
// @property {bool} Enable - A boolean value indicating whether synchronization is enabled or not.
// @property {bool} Persistent - The `Persistent` property indicates whether the synchronization should
// be persistent or not. If set to `true`, the synchronization will persist even after the application
// restarts. If set to `false`, the synchronization will only occur during the current session and will
// not persist after the application restarts.
// @property {string} Interval - The `Interval` property specifies the time interval at which the
// synchronization should occur. It is a string value that represents a duration, such as "1h" for 1
// hour or "30s" for 30 seconds.
// @property {string} Source - The `Source` property represents the source from where the data will be
// synchronized. It could be a file path, a database connection string, or any other source of data.
// @property {string} Store - The `Store` property is used to specify the location where the
// synchronized data will be stored.
// @property {string} Addition - The "Addition" property is an optional field that can be used to
// provide additional information or configuration for the synchronization process.
type Sync struct {
	Enable     bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
	Persistent bool   `json:"persistent,omitempty" yaml:"persistent,omitempty"`
	Interval   string `json:"interval,omitempty" yaml:"interval,omitempty"`
	Source     string `json:"source,omitempty" yaml:"source,omitempty"`
	Store      string `json:"store,omitempty" yaml:"store,omitempty"`
	Addition   string `json:"addition,omitempty" yaml:"addition,omitempty"`
}
type ExpiryCheck struct {
	Enabled     bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	TimeLeft    string `json:"time_left,omitempty" yaml:"time_left,omitempty"`
	LogInterval string `json:"log_interval,omitempty" yaml:"log_interval,omitempty"`
}
type PKI struct {
	CA                string      `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert              string      `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key               string      `json:"key,omitempty" yaml:"key,omitempty"`
	Blocklist         []string    `json:"blocklist,omitempty" yaml:"blocklist,omitempty"`
	DisconnectInvalid bool        `json:"disconnect_invalid,omitempty" yaml:"disconnect_invalid,omitempty"`
	ExpiryCheck       ExpiryCheck `json:"expiry_check,omitempty" yaml:"expiry_check,omitempty"`
}
type DNS struct {
	Enable   bool              `json:"enable,omitempty" yaml:"enable,omitempty"`
	Addr     string            `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int               `json:"port,omitempty" yaml:"port,omitempty"`
	Interval int               `json:"interval,omitempty" yaml:"interval,omitempty"`
	Mirror   string            `json:"mirror,omitempty" yaml:"mirror,omitempty"`
	Records  map[string]string `json:"records,omitempty" yaml:"records,omitempty"`
}
type Tower struct {
	Service           bool                       `json:"service,omitempty" yaml:"service,omitempty"`
	DNS               DNS                        `json:"dns,omitempty" yaml:"dns,omitempty"`
	Interval          int                        `json:"interval,omitempty" yaml:"interval,omitempty"`
	DetectionPoint    map[string]map[string]any  `json:"detection_point,omitempty" yaml:"detection_point,omitempty"`
	RemoteAllowList   map[string]bool            `json:"remote_allow_list,omitempty" yaml:"remote_allow_list,omitempty"`
	RemoteAllowRanges map[string]map[string]bool `json:"remote_allow_ranges,omitempty" yaml:"remote_allow_ranges,omitempty"`
	LocalAllowList    map[string]any             `json:"local_allow_list,omitempty" yaml:"local_allow_list,omitempty"`
	AdvertiseAddrs    []string                   `json:"advertise_addrs,omitempty" yaml:"advertise_addrs,omitempty"`
}
type Listen struct {
	Addr          string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port          int    `json:"port,omitempty" yaml:"port,omitempty"`
	Batch         int    `json:"batch,omitempty" yaml:"batch,omitempty"`
	ReadBuffer    int    `json:"read_buffer,omitempty" yaml:"read_buffer,omitempty"`
	WriteBuffer   int    `json:"write_buffer,omitempty" yaml:"write_buffer,omitempty"`
	SendRecvError string `json:"send_recv_error,omitempty" yaml:"send_recv_error,omitempty"`
	Routines      int    `json:"routines,omitempty" yaml:"routines,omitempty"`
}
type Punchy struct {
	Enable          bool     `json:"enable,omitempty" yaml:"enable,omitempty"`
	Frequency       string   `json:"frequency,omitempty" yaml:"frequency,omitempty"`
	Respond         bool     `json:"respond,omitempty" yaml:"respond,omitempty"`
	Delay           string   `json:"delay,omitempty" yaml:"delay,omitempty"`
	RespondDelay    string   `json:"respond_delay,omitempty" yaml:"respond_delay,omitempty"`
	PreferredRanges []string `json:"preferred_ranges,omitempty" yaml:"preferred_ranges,omitempty"`
}
type Users struct {
	Name string   `json:"name,omitempty" yaml:"name,omitempty"`
	Keys []string `json:"keys,omitempty" yaml:"keys,omitempty"`
}
type SSHD struct {
	Enabled  bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Port     int     `json:"port,omitempty" yaml:"port,omitempty"`
	PointKey string  `json:"point_key,omitempty" yaml:"point_key,omitempty"`
	Users    []Users `json:"users,omitempty" yaml:"users,omitempty"`
}
type Socks5 struct {
	Addr     string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}
type Forward struct {
	Proto  string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Local  string `json:"local,omitempty" yaml:"local,omitempty"`
	Remote string `json:"remote,omitempty" yaml:"remote,omitempty"`
}
type Proxy struct {
	Socks5  []Socks5  `json:"socks5,omitempty" yaml:"socks5,omitempty"`
	Forward []Forward `json:"forward,omitempty" yaml:"forward,omitempty"`
}
type Routes struct {
	Mtu   int    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Route string `json:"route,omitempty" yaml:"route,omitempty"`
}
type RouteTable struct {
	Route  string `json:"route,omitempty" yaml:"route,omitempty"`
	Via    string `json:"via,omitempty" yaml:"via,omitempty"`
	Mtu    int    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Metric int    `json:"metric,omitempty" yaml:"metric,omitempty"`
	Enable bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
}
type Tun struct {
	Enable             bool         `json:"enable,omitempty" yaml:"enable,omitempty"`
	Dev                string       `json:"dev,omitempty" yaml:"dev,omitempty"`
	DropLocalBroadcast bool         `json:"drop_local_broadcast,omitempty" yaml:"drop_local_broadcast,omitempty"`
	DropMulticast      bool         `json:"drop_multicast,omitempty" yaml:"drop_multicast,omitempty"`
	TxQueue            int          `json:"tx_queue,omitempty" yaml:"tx_queue,omitempty"`
	Mtu                int          `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Routes             []Routes     `json:"routes,omitempty" yaml:"routes,omitempty"`
	RouteTable         []RouteTable `json:"route_table,omitempty" yaml:"route_table,omitempty"`
}
type Logging struct {
	Level      string `json:"level,omitempty" yaml:"level,omitempty"`
	Language   string `json:"lang,omitempty" yaml:"lang,omitempty"`
	Format     string `json:"format,omitempty" yaml:"format,omitempty"`
	FilePath   string `json:"file_path,omitempty" yaml:"file_path,omitempty"`
	MaxSize    int    `json:"max_size,omitempty" yaml:"max_size,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty" yaml:"max_backups,omitempty"`
	MaxAge     int    `json:"max_age,omitempty" yaml:"max_age,omitempty"`
}
type Stats struct {
	Type           string `json:"type,omitempty" yaml:"type,omitempty"`
	Listen         string `json:"listen,omitempty" yaml:"listen,omitempty"`
	Path           string `json:"path,omitempty" yaml:"path,omitempty"`
	NameSpace      string `json:"name_space,omitempty" yaml:"name_space,omitempty"`
	Extention      string `json:"extention,omitempty" yaml:"extention,omitempty"`
	Prefix         string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Protocol       string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Server         string `json:"server,omitempty" yaml:"server,omitempty"`
	MessageMetrics bool   `json:"message_metrics,omitempty" yaml:"message_metrics,omitempty"`
	TowerMetrics   bool   `json:"tower_metrics,omitempty" yaml:"tower_metrics,omitempty"`
}
type Handshakes struct {
	TryInterval      string `json:"try_interval,omitempty" yaml:"try_interval,omitempty"`
	Retries          int    `json:"retries,omitempty" yaml:"retries,omitempty"`
	TriggerBuffer    int    `json:"trigger_buffer,omitempty" yaml:"trigger_buffer,omitempty"`
	ChurnLimiting    bool   `json:"churn_limiting,omitempty" yaml:"churn_limiting,omitempty"`
	ChurnNumFailures int    `json:"churn_num_failures,omitempty" yaml:"churn_num_failures,omitempty"`
	ChurnPeriod      string `json:"churn_period,omitempty" yaml:"churn_period,omitempty"`
}
type Timers struct {
	ConnectionAliveInterval int `json:"connection_alive_interval,omitempty" yaml:"connection_alive_interval,omitempty"`
	PendingDeletionInterval int `json:"pending_deletion_interval,omitempty" yaml:"pending_deletion_interval,omitempty"`
}
type PSK struct {
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
	Keys any    `json:"keys,omitempty" yaml:"keys,omitempty"`
}
type Conntrack struct {
	TCPTimeout     string `json:"tcp_timeout,omitempty" yaml:"tcp_timeout,omitempty"`
	UDPTimeout     string `json:"udp_timeout,omitempty" yaml:"udp_timeout,omitempty"`
	DefaultTimeout string `json:"default_timeout,omitempty" yaml:"default_timeout,omitempty"`
}
type Outbound struct {
	Port  string `json:"port,omitempty" yaml:"port,omitempty"`
	Proto string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Point string `json:"point,omitempty" yaml:"point,omitempty"`
}
type Inbound struct {
	Port   string   `json:"port,omitempty" yaml:"port,omitempty"`
	Proto  string   `json:"proto,omitempty" yaml:"proto,omitempty"`
	Point  string   `json:"point,omitempty" yaml:"point,omitempty"`
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`
}
type Firewall struct {
	OutboundAction string     `json:"outbound_action,omitempty" yaml:"outbound_action,omitempty"`
	InboundAction  string     `json:"inbound_action,omitempty" yaml:"inbound_action,omitempty"`
	Conntrack      Conntrack  `json:"conntrack,omitempty" yaml:"conntrack,omitempty"`
	Outbound       []Outbound `json:"outbound,omitempty" yaml:"outbound,omitempty"`
	Inbound        []Inbound  `json:"inbound,omitempty" yaml:"inbound,omitempty"`
}

func newConfig() *config {
	mtu := 1300
	return &config{
		PKI: PKI{
			Blocklist: []string{},
		},
		Points: map[string][]string{},
		Tower: Tower{
			DNS: DNS{Interval: 60},
		},
		Listen: Listen{
			Addr:  "0.0.0.0",
			Port:  35533,
			Batch: 64,
		},
		Punchy: Punchy{
			Enable: true,
			Delay:  "1s",
		},
		Cipher: "aes",
		SSHD: SSHD{
			Users: []Users{},
		},
		Tun: Tun{
			Dev:                "vlan",
			DropLocalBroadcast: true,
			DropMulticast:      true,
			TxQueue:            500,
			Mtu:                mtu,
			Routes:             []Routes{},
			RouteTable:         []RouteTable{},
		},
		Logging: Logging{
			Level:  "info",
			Format: "text",
		},
		Stats: Stats{},
		Handshakes: Handshakes{
			TryInterval: "100ms",
			Retries:     20,
		},
		Firewall: Firewall{
			Conntrack: Conntrack{
				TCPTimeout:     "120h",
				UDPTimeout:     "3m",
				DefaultTimeout: "10m",
			},
			Outbound: []Outbound{
				{
					Port:  "any",
					Proto: "any",
					Point: "any",
				},
			},
			Inbound: []Inbound{},
		},
	}
}
