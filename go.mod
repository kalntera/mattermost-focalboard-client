module github.com/kalntera/mattermost-focalboard-client

go 1.20

replace github.com/francoispqt/gojay v1.2.13 => github.com/kalntera/gojay v0.0.0-20240119150629-2858d078a506

replace golang.org/x/oauth2 v0.16.0 => github.com/kalntera/oauth2 v0.0.0-20240122224933-40a291d11bc8

replace github.com/mattermost/mattermost/server/public v0.0.10-0.20231116111926-0bc542620ce2 => github.com/mattermost/mattermost/server/public v0.0.12

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattermost/mattermost/server/public v0.0.12
	github.com/rivo/uniseg v0.4.4
)

require (
	github.com/fatih/color v1.16.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/go-hclog v1.6.2 // indirect
	github.com/hashicorp/go-plugin v1.6.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/grpc v1.60.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

require (
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/dyatlov/go-opengraph/opengraph v0.0.0-20220524092352-606d7b1e5f8a // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.5 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/mattermost/go-i18n v1.11.1-0.20211013152124-5c415071e404 // indirect
	github.com/mattermost/ldap v0.0.0-20231116144001-0f480c025956 // indirect
	github.com/mattermost/logr/v2 v2.0.21 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/wiggin77/merror v1.0.5 // indirect
	github.com/wiggin77/srslog v1.0.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
