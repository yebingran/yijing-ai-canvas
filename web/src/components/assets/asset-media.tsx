import { FileText, Music2 } from "lucide-react";

export function AssetMedia({ type, url, coverUrl, title, compact = false }: { type: string; url?: string; coverUrl?: string; title: string; compact?: boolean }) {
    const className = compact ? "h-full w-full object-cover" : "max-h-[65vh] w-full object-contain";
    if (type === "video") return <video src={url} poster={coverUrl || undefined} controls={!compact} muted={compact} playsInline preload="metadata" aria-label={title} className={className} />;
    if (type === "audio") return compact ? <Music2 aria-label={title} className="mx-auto h-full w-8 text-current opacity-50" /> : <audio src={url} controls preload="none" aria-label={title} className="w-full" />;
    if (type === "image" || coverUrl) return <img src={compact ? coverUrl || url : type === "image" ? url : coverUrl} alt={title} loading="lazy" className={className} />;
    return <FileText aria-label={title} className="mx-auto h-full w-8 text-current opacity-50" />;
}
