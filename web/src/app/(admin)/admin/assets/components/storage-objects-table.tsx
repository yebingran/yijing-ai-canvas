"use client";

import { useEffect, useState } from "react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { App, Button, Flex, Input, Modal, Select, Space, Table, Tag, Typography } from "antd";
import { AssetMedia } from "@/components/assets/asset-media";
import { useCopyText } from "@/hooks/use-copy-text";
import { fetchAdminStorageObjects, publishAdminStorageObject, type AdminAsset, type AdminStorageObject } from "@/services/api/admin";
import { useUserStore } from "@/stores/use-user-store";

export function StorageObjectsTable({ onPublished }: { onPublished: (asset: AdminAsset) => void }) {
    const token = useUserStore((state) => state.token);
    const { message } = App.useApp();
    const copyText = useCopyText();
    const queryClient = useQueryClient();
    const [keyword, setKeyword] = useState("");
    const [type, setType] = useState("");
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);
    const [preview, setPreview] = useState<AdminStorageObject | null>(null);
    const query = useQuery({
        queryKey: ["admin", "storage-objects", token, keyword, type, page, pageSize],
        queryFn: () => fetchAdminStorageObjects(token, { keyword, type, page, pageSize }),
        enabled: Boolean(token),
        retry: false,
    });
    const publish = useMutation({
        mutationFn: (id: string) => publishAdminStorageObject(token, id),
        onSuccess: async (asset) => {
            await Promise.all([
                queryClient.invalidateQueries({ queryKey: ["admin", "storage-objects"] }),
                queryClient.invalidateQueries({ queryKey: ["admin", "assets"] }),
                queryClient.invalidateQueries({ queryKey: ["asset-library"] }),
            ]);
            message.success("已设为公共素材，可继续编辑标题、封面和首页展示");
            onPublished(asset);
        },
        onError: (error) => message.error(error instanceof Error ? error.message : "发布失败"),
    });
    useEffect(() => {
        if (query.isError) message.error(query.error instanceof Error ? query.error.message : "获取同步记录失败");
    }, [message, query.error, query.isError]);

    return <Flex vertical gap={16}>
        <Typography.Text type="secondary">显示应用已同步到系统 OSS / S3 / WebDAV 的文件及上传用户。设为公共素材后可在素材库使用，再按需开启首页展示。未通过本应用同步的桶内文件不在此列表。</Typography.Text>
        <Space wrap>
            <Input.Search placeholder="搜索文件路径、用户名或用户 ID" allowClear style={{ width: 320 }} onSearch={(value) => { setKeyword(value.trim()); setPage(1); }} />
            <Select aria-label="同步文件类型" value={type} style={{ width: 120 }} onChange={(value) => { setType(value); setPage(1); }} options={[{ label: "全部类型", value: "" }, { label: "图片", value: "image" }, { label: "视频", value: "video" }, { label: "音频", value: "audio" }]} />
            <Button onClick={() => void query.refetch()}>刷新同步记录</Button>
        </Space>
        <Table<AdminStorageObject>
            rowKey="id" size="small" loading={query.isFetching} dataSource={query.data?.items || []} scroll={{ x: 1250 }}
            pagination={{ current: page, pageSize, total: query.data?.total || 0, showSizeChanger: true, showTotal: (total) => `共 ${total} 条`, onChange: (next, size) => { setPage(size === pageSize ? next : 1); setPageSize(size); } }}
            columns={[
                { title: "预览", width: 88, render: (_, item) => <button type="button" aria-label={`预览${item.objectKey}`} className="block h-12 w-16 cursor-pointer overflow-hidden rounded" onClick={() => setPreview(item)}><AssetMedia type={item.mimeType.split("/")[0]} url={item.url} title={item.objectKey} compact /></button> },
                { title: "文件 / 存储位置", width: 260, render: (_, item) => <><Typography.Text ellipsis={{ tooltip: item.objectKey }} style={{ display: "block", maxWidth: 250 }}>{item.objectKey}</Typography.Text><Typography.Text type="secondary">{item.providerName || "系统存储"}{item.bucket ? ` / ${item.bucket}` : ""}</Typography.Text></> },
                { title: "同步用户", width: 180, render: (_, item) => <><Typography.Text style={{ display: "block" }}>{item.userDisplayName || item.username || "未命名用户"}</Typography.Text><Typography.Text type="secondary" ellipsis={{ tooltip: item.createdBy }} style={{ display: "block", maxWidth: 170 }}>{item.username ? `${item.username} · ` : ""}{item.createdBy}</Typography.Text></> },
                { title: "格式 / 大小", width: 150, render: (_, item) => <><Tag>{item.mimeType}</Tag><Typography.Text type="secondary" style={{ display: "block" }}>{(item.bytes / 1024 / 1024).toFixed(2)} MB</Typography.Text></> },
                { title: "同步时间", width: 170, render: (_, item) => new Date(item.createdAt).toLocaleString("zh-CN", { hour12: false }) },
                { title: "链接", width: 140, render: (_, item) => <Space orientation="vertical" size={0}><Button size="small" type="link" onClick={() => copyText(new URL(item.url, window.location.origin).href)}>复制访问链接</Button>{item.publicUrl ? <Button size="small" type="link" onClick={() => copyText(item.publicUrl)}>复制 OSS 直链</Button> : null}</Space> },
                { title: "公共素材", width: 150, render: (_, item) => <Button size="small" disabled={Boolean(item.assetId) || !/^(image|video|audio)\//.test(item.mimeType)} loading={publish.isPending && publish.variables === item.id} onClick={() => publish.mutate(item.id)}>{item.assetId ? "已设为公共素材" : "设为公共素材"}</Button> },
            ]}
        />
        <Modal title="同步文件预览" open={Boolean(preview)} footer={null} width={800} destroyOnHidden onCancel={() => setPreview(null)}>
            {preview ? <Flex vertical gap={16}><AssetMedia type={preview.mimeType.split("/")[0]} url={preview.url} title={preview.objectKey} /><Typography.Paragraph copyable={{ text: new URL(preview.url, window.location.origin).href }}>{new URL(preview.url, window.location.origin).href}</Typography.Paragraph></Flex> : null}
        </Modal>
    </Flex>;
}
