package config

type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`                                                 // 数据库驱动
	Host                string `mapstructure:"host" json:"host" yaml:"host"`                                                       //数据库地址
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`                                                       //数据库端口
	Database            string `mapstructure:"database" json:"database" yaml:"database"`                                           //数据库名称
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`                                           //用户名
	Password            string `mapstructure:"password" json:"password" yaml:"password"`                                           //密码
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`                                              //编码格式
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`                         //空闲连接池中连接的最大数量
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`                         //打开数据库连接的最大数量
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                                           //日志级别
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"` //是否启用日志文件
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`                               //日志文件名称
}
