<p align="center">
  <img src="web/public/logo.svg" width="96" alt="Infinite Canvas logo">
</p>

<h1 align="center">翼镜</h1>

让灵感在画布上成境。面向图片、视频与音频创作的 AI 工作台，在一张无限画布中组织脚本、参考图、模型调用和生成结果。

项目演示：[wuxian.yunsuguan.com](https://wuxian.yunsuguan.com) · 定制仓库：[yebingran/yijing-ai-canvas](https://github.com/yebingran/yijing-ai-canvas)

## 核心功能

- 无限画布：多项目、节点连线、参考素材、撤销重做和导入导出。
- AI 创作：图片、视频、音频生成，支持多个模型渠道、算力萌视频和 New API。
- 创作辅助：Agent、提示词库、素材库、工作流与 3D 导演台。
- 后端管理：账号与云同步、任务恢复、AI 日志、对象存储和模型渠道配置。
- 页面配置：管理员可分别维护顶部文字公告和首页大图轮播，支持排序与停用。

未登录时，画布项目和“我的素材”保存在浏览器本地；登录且账号同步可用时，会同步到账号/云端。本地直连模式的 AI API Key 保存在浏览器本地，由前端直接请求对应接口。

## 快速开始

仓库和发布镜像均为私有，需先获得对应访问权限。

```bash
git clone git@github.com:yebingran/yijing-ai-canvas.git
cd yijing-ai-canvas
cp .env.example .env
# 修改管理员密码和部署配置；先登录 ghcr.io，再启动
docker compose up -d
```

默认访问 `http://localhost:3000`。镜像地址为 `ghcr.io/yebingran/yijing-ai-canvas`，1Panel 配置私有仓库凭据后即可拉取。详细步骤见 [Docker 部署](docs/overview/docker.md)。

## 赞助支持

**[翼龙云 · www.yilongcloud.com](https://www.yilongcloud.com)** 为本项目提供赞助。提供阿里云国际、腾讯云国际、AWS、微软云 Azure、谷歌云 Google Cloud 等国际云厂商账号服务，支持代充值，面向开发者和企业的国际云资源使用需求。

**[云速管 · www.yunsuguan.com](https://www.yunsuguan.com)** 是自研云账号管理平台，围绕云账号管理与日常运维提供统一入口，帮助团队管理云资源与账号。

**[云速管 AI 中转平台 · ai.yunsuguan.com](https://ai.yunsuguan.com)** 提供 GPT、Gemini 等模型调用入口，主打“原厂对接，100% 保真”，便于开发者与应用接入 AI 能力。

## 项目来源与协议

本项目基于 [basketikun(纯前端)](https://github.com/basketikun/infinite-canvas) 为底，合并 [HuFakai](https://github.com/HuFakai/infinite-canvas) 生图增强版基础上，针对视频和视频生成逻辑配置更加完善，完善后端云同步机制，不再依赖纯前端。

翼镜 定制版继续基于 [TigerOWO / Infinite Canvas](https://github.com/tigerowo/infinite-canvas) 迭代，保留原作者信息、前端页面标识和 Git 历史。感谢所有原作者与贡献者。

> [!CAUTION]
> 项目目前处于开发阶段，不保证历史数据兼容。各种数据库结构和存储格式都可能直接调整，欢迎关注后续更新。
>
> 如果你需要稳定维护自己的分支，建议自行 fork 后独立开发。二次开发与 PR 请保留原作者信息和前端页面标识。

本项目继续采用 GNU Affero General Public License v3.0（AGPL-3.0-only），见 [LICENSE](LICENSE) 与 [协议说明](docs/business/license.md)。仓库访问权限不改变原开源许可及对应义务。

## 文档

- [快速开始](docs/overview/quick-start.md) · [功能介绍](docs/overview/features.md) · [Docker 与 1Panel 部署](docs/overview/docker.md)
- [画布操作手册](docs/canvas/canvas-node-manual.md) · [快捷键](docs/canvas/canvas-shortcuts.md)
- [官方更新合并](docs/backend/upstream-updates.md) · [待测试](docs/progress/pending-test.md)
- [文档索引](docs/index.md)
