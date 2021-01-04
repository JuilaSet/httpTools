package exception

import "errors"

var ConnectTimeoutException = errors.New("connect failed")