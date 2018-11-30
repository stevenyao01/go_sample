package sdk

/**
 * @Package Name: sdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-13 上午11:05
 * @Description:
 */

type MqSub struct {
	topic             string
	qos               byte
	callbackReceive   CbReceive
	callbackBroadCast CbBroadCast
}

func NewMqSub(t string, q byte, c CbReceive, b CbBroadCast) (*MqSub, error) {
	return &MqSub{
		topic:             t,
		qos:               q,
		callbackReceive:   c,
		callbackBroadCast: b,
	}, nil
}