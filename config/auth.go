/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// JWTConfig 配置信息
type JWTConfig struct {
	SigningKey    string `mapstructure:"signingKey" json:"signingKey"`
	TokenKey      string `mapstructure:"tokenKey" json:"tokenKey"`
	EffectiveTime int    `mapstructure:"effectiveTime" json:"effectiveTime"`
}

// AuthConfig Auth配置
type AuthConfig struct {
	JWTInfo JWTConfig `mapstructure:"jwt" json:"jwt"`
}
