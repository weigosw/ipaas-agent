#!/bin/bash

# 构建和运行脚本
# 演示多环境配置的使用

echo "=== 构建 ipaas-agent ==="
go build -o ipaas-agent .

echo ""
echo "=== 多环境配置演示 ==="
echo ""

echo "1. 默认环境（dev）:"
echo "   ./ipaas-agent"
echo ""

echo "2. 指定开发环境:"
echo "   ./ipaas-agent -env dev"
echo ""

echo "3. 指定生产环境:"
echo "   ./ipaas-agent -env prod"
echo ""

echo "4. 通过环境变量指定:"
echo "   IPAAS_ENV=prod ./ipaas-agent"
echo ""

echo "5. 使用环境变量覆盖配置:"
echo "   AUTH_CLIENTID=custom_id ./ipaas-agent -env dev"
echo ""

echo "=== 当前可用的配置文件 ==="
ls -la config/
