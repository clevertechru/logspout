package main

import (
	_ "github.com/clevertechru/logspout/healthcheck"
	_ "github.com/clevertechru/logspout/adapters/raw"
	_ "github.com/clevertechru/logspout/adapters/syslog"
	_ "github.com/clevertechru/logspout/adapters/multiline"
	_ "github.com/clevertechru/logspout/adapters/file"
	_ "github.com/clevertechru/logspout/httpstream"
	_ "github.com/clevertechru/logspout/routesapi"
	_ "github.com/clevertechru/logspout/transports/tcp"
	_ "github.com/clevertechru/logspout/transports/udp"
	_ "github.com/clevertechru/logspout/transports/tls"
)
