package common

var ERU_AGENT_VERSION = "unknown"

const (
	DEFAULT_ETCD_PREFIX = "eru"
	DOCKER_CLI_VERSION  = "1.25"

	STATUS_DIE     = "die"
	STATUS_START   = "start"
	STATUS_DESTROY = "destroy"

	DATETIME_FORMAT = "2006-01-02 15:04:05.999999"
	CNAME_NUM       = 3

	CGROUP_BASE_PATH = "/sys/fs/cgroup/%s/docker/%s/%s"

	VLAN_PREFIX = "cali0"
	DEFAULT_BR  = "eth0"
	DOCKERIZED  = "AGENT_IN_DOCKER"
)
