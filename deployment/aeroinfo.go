package deployment

import (
	"fmt"

	"github.com/go-logr/logr"

	aero "github.com/ashishshinde/aerospike-client-go/v6"
)

// IsClusterAndStable returns true if the cluster formed by the set of hosts is stable.
func IsClusterAndStable(log logr.Logger, policy *aero.ClientPolicy, allHosts []*HostConn) (bool, error) {
	c, err := newCluster(log, policy, allHosts, allHosts)
	if err != nil {
		return false, fmt.Errorf("unable to create a cluster copy for running aeroinfo: %v", err)
	}

	return c.IsClusterAndStable(getHostIDsFromHostConns(allHosts))
}

// InfoQuiesce quiesce hosts.
func InfoQuiesce(log logr.Logger, policy *aero.ClientPolicy, allHosts, selectedHosts []*HostConn,
	removedNamespaces []string) error {
	c, err := newCluster(log, policy, allHosts, selectedHosts)
	if err != nil {
		return fmt.Errorf("unable to create a cluster copy for running aeroinfo: %v", err)
	}

	return c.InfoQuiesce(getHostIDsFromHostConns(selectedHosts), getHostIDsFromHostConns(allHosts), removedNamespaces)
}

// InfoQuiesceUndo revert the effects of quiesce on the next recluster event
func InfoQuiesceUndo(log logr.Logger, policy *aero.ClientPolicy, allHosts []*HostConn) error {
	c, err := newCluster(log, policy, allHosts, allHosts)
	if err != nil {
		return fmt.Errorf("unable to create a cluster copy for running aeroinfo: %v", err)
	}

	return c.InfoQuiesceUndo(getHostIDsFromHostConns(allHosts))
}

// SetMigrateFillDelay sets the given migrate-fill-delay on all the given cluster nodes
func SetMigrateFillDelay(log logr.Logger, policy *aero.ClientPolicy, allHosts []*HostConn, migrateFillDelay int) error {
	c, err := newCluster(log, policy, allHosts, allHosts)
	if err != nil {
		return fmt.Errorf("unable to create a cluster copy for running aeroinfo: %v", err)
	}

	return c.setMigrateFillDelay(migrateFillDelay, allHosts)
}

// GetClusterNamespaces gets the cluster namespaces
func GetClusterNamespaces(log logr.Logger, policy *aero.ClientPolicy,
	allHosts []*HostConn) (map[string][]string, error) {
	c, err := newCluster(log, policy, allHosts, allHosts)
	if err != nil {
		return nil, fmt.Errorf("unable to create a cluster copy for running aeroinfo: %v", err)
	}

	return c.getClusterNamespaces(getHostIDsFromHostConns(allHosts))
}
