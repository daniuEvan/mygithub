/**
 * @date: 2022/3/19
 * @desc: ...
 */

package config

// MysqlConfig mysql 配置
type MysqlConfig struct {
	DBName          string `mapstructure:"dbname" json:"dbname"`
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"`
}

// PostgresConfig pg 配置
type PostgresConfig struct {
	DBName          string `mapstructure:"dbname" json:"dbname"`
	Schema          string `mapstructure:"schema" json:"schema"`
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	Username        string `mapstructure:"username" json:"username"`
	Password        string `mapstructure:"password" json:"password"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime"`
}

type DatabaseConfig struct {
	DBType       string         `mapstructure:"dbType" json:"dbType"`
	TablePrefix  string         `mapstructure:"tablePrefix" json:"tablePrefix"`
	MysqlInfo    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	PostgresInfo PostgresConfig `mapstructure:"postgres" json:"postgres"`
}
