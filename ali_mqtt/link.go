package ali_mqtt

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliLog struct{}

func (a AliLog) Printf(format string, v ...interface{}) {
	g.Log().Printf(context.Background(), format, v...)
}
func (a AliLog) Println(v ...interface{}) {
	g.Log().Print(context.Background(), v...)
}

var (
	client  MQTT.Client
	GetMsgs = make(chan []byte, 10)
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) { //收到从阿里云下发的消息
	GetMsgs <- msg.Payload()
}

type AliConfig struct {
	subTopic     string //
	pubTopic     string //
	qos          byte
	productKey   string
	deviceName   string
	deviceSecret string
}

func (a *AliConfig) newTLSConfig() *tls.Config {
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM([]byte(RootPem))
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: false,
	}
}
func (a *AliConfig) conn() error {
	MQTT.ERROR = AliLog{}

	a.subTopic = "/" + a.productKey + "/" + a.deviceName + "/user/get"    //下行
	a.pubTopic = "/" + a.productKey + "/" + a.deviceName + "/user/update" //上行
	var raw_broker bytes.Buffer
	raw_broker.WriteString("tcp://")
	raw_broker.WriteString(a.productKey)
	raw_broker.WriteString(".iot-as-mqtt.cn-shanghai.aliyuncs.com:1883")
	opts := MQTT.NewClientOptions().AddBroker(raw_broker.String())

	auth := calculate_sign(a.deviceName, a.productKey, a.deviceName, a.deviceSecret, gconv.String(time.Now().Unix()*1000))

	opts.SetClientID(auth.mqttClientId)
	opts.SetUsername(auth.username)
	opts.SetPassword(auth.password)
	opts.SetKeepAlive(120 * time.Second)
	opts.SetDefaultPublishHandler(f)

	tlsconfig := a.newTLSConfig()
	opts.SetTLSConfig(tlsconfig)

	client = MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		g.Log().Debug(context.Background(), "connect true")
	}
	if token := client.Subscribe(a.subTopic, a.qos, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		g.Log().Debug(context.Background(), "sub true")
	}
	return nil
}

func (a *AliConfig) sendMsg(msg []byte) error {
	if IsConnected() {
		if token := client.Publish(a.pubTopic, a.qos, false, msg); token.Wait() && token.Error() != nil {
			g.Log().Debug(context.Background(), "to ali cloud offline error")
			return token.Error()
		}
	}
	return errors.New("not conn")
}
