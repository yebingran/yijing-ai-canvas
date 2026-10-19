---
title: 官方更新合并
description: 私有定制仓库的分支与上游维护方式
---

# 官方更新合并

独立仓库：[yebingran/yijing-ai-canvas](https://github.com/yebingran/yijing-ai-canvas)。它不属于 GitHub fork 网络，使用 Git 历史继续合并官方更新。

## 分支约定

- `main`：合入官方更新的翼境 AI 定制发布分支。
- `comemoe-video-adapter`：保留的定制分支，迁移发布时与 `main` 对齐。
- `codex/announcement-settings`、`codex/newapi-image-support`：保留迁移前的定制分支。
- `codex/pre-private-migration`：迁移前的完整定制版本。
- `codex/source-main`：迁移前公开仓库的主分支。
- `codex/upstream-migration`：本次合入的官方提交，保留官方历史。

本次官方基线为 `ccda172efba25cbee990a785bba715b90cc42b95`；迁移前定制基线为 `c10b7e902c14fa14b5994bf6d01e4a8b9da83d7e`。

## 后续更新

首次配置：

```bash
git remote add upstream https://github.com/tigerowo/infinite-canvas.git
```

在干净工作区创建更新分支并合并：

```bash
git fetch origin
git fetch upstream main --tags
git switch main
git pull --ff-only origin main
git switch -c codex/merge-upstream
git merge upstream/main
```

按实际冲突处理后提交，经人工验证再合并回 `main`。不要用官方分支强制覆盖定制分支，也不要把私有定制代码推到公开上游。

合并时保留：算力萌请求与轮询、524 任务恢复、生成结果先显示再同步 OSS、New API 与图片解析/恢复修复，以及 `public.announcement` 和 `public.homeBanners` 两套后台配置。README 赞助信息、私有镜像地址和自己的仓库链接也需要保留。

`main` 推送及版本标签触发私有 GHCR 镜像构建；更新镜像后由服务器重新拉取部署。原 `@tigerowo/canvas-agent` 依赖及原作者信息继续保留，私有仓库不会自动向公共 npm 发布该包。
