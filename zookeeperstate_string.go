// generated by stringer -type=ZookeeperState; DO NOT EDIT

package gozoo

import "fmt"

const _ZookeeperState_name = "ZooSessionStateZooAuthFailedStateZooConnectingStateZooAssociatingStateZooConnectedStateZooUnknownState"

var _ZookeeperState_index = [...]uint8{0, 15, 33, 51, 70, 87, 102}

func (i ZookeeperState) String() string {
	if i < 0 || i+1 >= ZookeeperState(len(_ZookeeperState_index)) {
		return fmt.Sprintf("ZookeeperState(%d)", i)
	}
	return _ZookeeperState_name[_ZookeeperState_index[i]:_ZookeeperState_index[i+1]]
}
