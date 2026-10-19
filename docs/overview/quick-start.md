---
title: 快速开始
description: 启动翼境 AI · 无限画布
---

# 快速开始

演示地址：[wuxian.yunsuguan.com](https://wuxian.yunsuguan.com)。自行部署建议使用私有发布镜像，完整凭据与 1Panel 设置见 [Docker 部署](docker.md)。

```bash
git clone git@github.com:yebingran/yijing-ai-canvas.git
cd yijing-ai-canvas
cp .env.example .env
# 修改 ADMIN_PASSWORD，并完成 ghcr.io 登录
docker compose up -d
```

启动后访问 `http://localhost:3000`，管理员用户名为 `admin`，密码由 `.env` 中的 `ADMIN_PASSWORD` 设置。

## 本地构建

```bash
cp .env.example .env
docker compose -f docker-compose.local.yml up -d --build
```

## 首次使用

- 本地直连模式：打开右上角配置，填写 Base URL、API Key 和模型名。
- 后台渠道模式：进入管理后台配置系统模型和渠道。
- 提示词仓库：在 `/admin/prompts` 拉取或同步。
- 顶部文字公告与首页大图轮播：在「系统设置 → 公开配置 → 可视化编辑」分别配置，保存后回到首页查看。

未登录时，画布项目和“我的素材”保存在浏览器本地；登录且账号同步可用时，同步到账号/云端。本地直连模式的 AI API Key 保存在浏览器本地，由前端直接请求对应接口。
