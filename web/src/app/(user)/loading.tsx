import { LoaderCircle } from "lucide-react";

export default function Loading() {
    return <div role="status" aria-busy="true" className="flex h-full min-h-48 flex-col items-center justify-center gap-3 bg-background text-foreground">
        <LoaderCircle className="size-7 animate-spin motion-reduce:animate-none" aria-hidden />
        <p className="text-sm opacity-65">页面加载中，请稍候…</p>
    </div>;
}
