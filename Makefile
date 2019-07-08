VERSION := 0.0.3
TARGET := rsyslog_exporter
GOFLAGS := -ldflags "-X main.Version=$(VERSION)"
ROOTPKG := github.com/digitalocean/$(TARGET)
GO_VERSION := 1.12.6

include Makefile.COMMON
