# 多环境配置使用说明

本项目现在支持多环境配置文件切换，您可以通过以下方式来指定不同的运行环境：

## 配置文件结构

项目支持以下配置文件：

- `config/config.dev.yml` - 开发环境配置
- `config/config.prod.yml` - 生产环境配置  
- `config/test.yml` - 测试环境配置
- `config.yaml` - 默认配置文件（如果没有指定环境）

## 使用方法

### 1. 通过命令行参数指定环境

```bash
# 启动开发环境
./ipaas-agent -env dev

# 启动生产环境  
./ipaas-agent -env prod

# 启动测试环境
./ipaas-agent -env test
```

### 2. 通过环境变量指定

```bash
# 设置环境变量
export IPAAS_ENV=prod
./ipaas-agent

# 或者一行命令
IPAAS_ENV=dev ./ipaas-agent
```

### 3. 优先级

配置的优先级为：命令行参数 > 环境变量 > 默认值(dev)

如果没有指定任何环境，默认使用 `dev` 环境。

## 环境变量覆盖

除了切换配置文件，您还可以使用环境变量来覆盖配置文件中的特定值：

```bash
# 覆盖认证相关配置
export AUTH_CLIENTID="your_client_id"
export AUTH_CLIENTSECRET="your_client_secret" 
export IPAAS_AGENT_AUTH_OPEN_API_HOST="https://custom.api.host.com"

./ipaas-agent -env prod
```

## 配置文件示例

### config.dev.yml

```yaml
auth:
  clientID: "dev_client_id"
  clientSecret: "dev_client_secret"
  openAPIHost: "https://dev.api.dingtalk.com"

plugins:
  - type: "mysql"
    host: "localhost"
    port: 3306
    username: "dev_user"
    password: "dev_password"
    database: "dev_database"
```

### config.prod.yml

```yaml
auth:
  clientID: "prod_client_id"
  clientSecret: "prod_client_secret"
  openAPIHost: "https://api.dingtalk.com"

plugins:
  - type: "mysql"
    host: "prod-db.example.com"
    port: 3306
    username: "prod_user"
    password: "prod_password"
    database: "prod_database"
```

## 日志输出

程序启动时会显示当前使用的环境和配置文件：

```shell
INFO[0000] 启动程序                  buildTime=DEV environment=dev gitCommit=DEV version=DEV
INFO[0000] 加载配置文件...             environment=dev
INFO[0000] 成功加载配置文件            configFile=/path/to/config/config.dev.yml
```
