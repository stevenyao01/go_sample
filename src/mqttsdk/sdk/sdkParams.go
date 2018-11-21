package sdk

/**
 * @Package Name: sdk
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-11-12 下午5:08
 * @Description:
 */

type SdkParams struct {
	server    string
	device_sk string
	device_id string
}

func NewSdkParams(s string, ds string, di string) (*SdkParams, error) {
	return &SdkParams{
		server:s,
		device_sk:ds,
		device_id:di,
	},nil
}