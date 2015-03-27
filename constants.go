package gozoo

// #include <zookeeper/zookeeper.h>
import "C"

type ZookeeperEvent int
type ZookeeperError int
type ZookeeperState int
type ZookeeperCreateFlag int

const (
	ZooCreatedEvent ZookeeperEvent = iota
	ZooDeletedEvent
	ZooChangedEvent
	ZooChildEvent
	ZooSessionEvent
	ZooNotWatchingEvent
	ZooUnknownEvent
)

const (
	ZooOk ZookeeperError = iota
	ZooSystemError
	ZooRuntimeInconsistencyError
	ZooDataInconsistencyError
	ZooConnectionLossError
	ZooMarshallingError
	ZooUnimplementedError
	ZooOperationTimeoutError
	ZooBadArgumentsError
	ZooInvalidStateError
	ZooApiError
	ZooNoNodeError
	ZooNoAuthError
	ZooBadVersionError
	ZooNoChildrenForEphemeralsError
	ZooNodeExistsError
	ZooNotEmptyError
	ZooSessionExpiredError
	ZooInvalidCallbackError
	ZooInvalidAclError
	ZooAuthFailedError
	ZooClosingError
	ZooNothingError
	ZooSessionMovedError
	ZooUnknownError
)

const (
	ZooSessionState ZookeeperState = iota
	ZooAuthFailedState
	ZooConnectingState
	ZooAssociatingState
	ZooConnectedState
	ZooUnknownState
)

const (
	ZooEphemeralCreateFlag ZookeeperCreateFlag = iota
	ZooSequenceCreateFlag
)

func convertZookeeperEvent(value C.int) ZookeeperEvent {
	switch value {
	case C.ZOO_CREATED_EVENT:
		return ZooCreatedEvent
	case C.ZOO_DELETED_EVENT:
		return ZooDeletedEvent
	case C.ZOO_CHANGED_EVENT:
		return ZooChangedEvent
	case C.ZOO_CHILD_EVENT:
		return ZooChildEvent
	case C.ZOO_SESSION_EVENT:
		return ZooSessionEvent
	case C.ZOO_NOTWATCHING_EVENT:
		return ZooNotWatchingEvent
	default:
		return ZooUnknownEvent
	}
}

func convertZookeeperState(value C.int) ZookeeperState {
	switch value {
	case C.ZOO_EXPIRED_SESSION_STATE:
		return ZooSessionState
	case C.ZOO_AUTH_FAILED_STATE:
		return ZooAuthFailedState
	case C.ZOO_CONNECTING_STATE:
		return ZooConnectingState
	case C.ZOO_ASSOCIATING_STATE:
		return ZooAssociatingState
	case C.ZOO_CONNECTED_STATE:
		return ZooConnectedState
	default:
		return ZooUnknownState
	}
}

func convertZookeeperError(value C.int) ZookeeperError {
	switch value {
	case C.ZOK:
		return ZooOk
	case C.ZSYSTEMERROR:
		return ZooSystemError
	case C.ZRUNTIMEINCONSISTENCY:
		return ZooRuntimeInconsistencyError
	case C.ZDATAINCONSISTENCY:
		return ZooDataInconsistencyError
	case C.ZCONNECTIONLOSS:
		return ZooConnectionLossError
	case C.ZMARSHALLINGERROR:
		return ZooMarshallingError
	case C.ZUNIMPLEMENTED:
		return ZooUnimplementedError
	case C.ZOPERATIONTIMEOUT:
		return ZooOperationTimeoutError
	case C.ZBADARGUMENTS:
		return ZooBadArgumentsError
	case C.ZINVALIDSTATE:
		return ZooInvalidStateError
	case C.ZAPIERROR:
		return ZooApiError
	case C.ZNONODE:
		return ZooNoNodeError
	case C.ZNOAUTH:
		return ZooNoAuthError
	case C.ZBADVERSION:
		return ZooBadVersionError
	case C.ZNOCHILDRENFOREPHEMERALS:
		return ZooNoChildrenForEphemeralsError
	case C.ZNODEEXISTS:
		return ZooNodeExistsError
	case C.ZNOTEMPTY:
		return ZooNotEmptyError
	case C.ZSESSIONEXPIRED:
		return ZooSessionExpiredError
	case C.ZINVALIDCALLBACK:
		return ZooInvalidCallbackError
	case C.ZINVALIDACL:
		return ZooInvalidAclError
	case C.ZAUTHFAILED:
		return ZooAuthFailedError
	case C.ZCLOSING:
		return ZooClosingError
	case C.ZNOTHING:
		return ZooNothingError
	case C.ZSESSIONMOVED:
		return ZooSessionMovedError
	default:
		return ZooUnknownError
	}
}

func convertToCreateFlag(flag ZookeeperCreateFlag) C.int {
	switch flag {
	case ZooEphemeralCreateFlag:
		return C.ZOO_EPHEMERAL
	case ZooSequenceCreateFlag:
		return C.ZOO_SEQUENCE
	default:
		return -1
	}
}
