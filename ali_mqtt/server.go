package ali_mqtt

import "errors"

var aliServer *AliConfig

// build ali connect config
func BuildAli(productKey, deviceName, deviceSecret string, qos byte) {
	a := &AliConfig{}
	a.qos = qos
	a.productKey = productKey
	a.deviceName = deviceName
	a.deviceSecret = deviceSecret
	aliServer = a
}

// to link aliyun mqtt
func ToConn() error {
	return aliServer.conn()
}

// to send msg
func ToSend(msg []byte) error {
	if aliServer != nil {
		return aliServer.sendMsg(msg)
	}
	return errors.New("aliServer is nil")
}

// jude aliyun mqtt is connected
func IsConnected() bool {
	if client == nil {
		return false
	}
	return client.IsConnected()
}
