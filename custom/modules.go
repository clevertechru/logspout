package main

import (
	_ "github.com/clevertechru/logspout/adapters/syslog"
	_ "github.com/clevertechru/logspout/transports/tcp"
	_ "github.com/clevertechru/logspout/transports/tls"
	_ "github.com/clevertechru/logspout/transports/udp"
	_ "github.com/looplab/logspout-logstash"
)
