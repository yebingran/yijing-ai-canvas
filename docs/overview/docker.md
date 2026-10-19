---
title: Docker 与 1Panel 部署
description: 使用私有 GHCR 镜像部署翼境 AI
---

# Docker 与 1Panel 部署

私有镜像：`ghcr.io/yebingran/yijing-ai-canvas`。支持 Linux amd64、arm64。建议部署具体版本标签；`latest` 用于跟随主分支最新构建。

## 私有仓库认证

在 GitHub 账号中创建具有 `read:packages` 权限的 personal access token（classic）。持有者须有私有镜像读取权限。Token 仅在服务器或 1Panel 仓库凭据中填写，不要写入源码、Compose、截图或聊天。

终端使用：

```bash
docker login ghcr.io -u yebingran
```

在密码提示处输入 Token。1Panel 中添加容器镜像仓库：地址 `ghcr.io`，用户名 `yebingran`，密码填 Token。编排内将原镜像替换为 `ghcr.io/yebingran/yijing-ai-canvas:v0.8.2`，拉取后重新创建容器。

## Compose 启动

```bash
git clone git@github.com:yebingran/yijing-ai-canvas.git
cd yijing-ai-canvas
cp .env.example .env
# 修改管理员密码和其他环境变量
docker compose up -d
```

访问 `http://localhost:3000`。默认管理员用户名 `admin`，密码由 `.env` 中的 `ADMIN_PASSWORD` 设置。

更新镜像：

```bash
docker compose pull
docker compose up -d
```

迁换镜像时，保留原环境变量、端口和数据挂载，不要创建空的数据目录代替原目录。更新前备份数据库和素材。

## 本地构建

```bash
cp .env.example .env
docker compose -f docker-compose.local.yml up -d --build
```

## 数据目录

`docker-compose.yml` 把 `./data` 挂载到容器内 `/app/data`，保存 SQLite 数据库、提示词数据和上传素材。SQLite 配置示例：

```text
DATABASE_DSN=/app/data/infinite-canvas.db
```

对象存储等配置沿用原部署。静态资源路径的生产部署验证仍列为待办；本版本构建成功不等同于已完成服务器环境的全部验证。
