"use client";

import { LoaderCircle, type LucideIcon } from "lucide-react";
import Link, { useLinkStatus } from "next/link";
import { usePathname } from "next/navigation";
import { useRef, useState, type ReactNode } from "react";
import { createPortal } from "react-dom";

export function NavigationLink({ href, icon, children, className, onNavigate }: { href: string; icon?: LucideIcon; children: ReactNode; className?: string; onNavigate?: () => void }) {
    const pathname = usePathname();
    const linkRef = useRef<HTMLAnchorElement>(null);
    const [prefetch, setPrefetch] = useState<boolean | null>(null);

    return <Link
        ref={linkRef}
        href={href}
        className={className}
        aria-current={pathname === href ? "page" : undefined}
        prefetch={prefetch}
        onMouseEnter={() => setPrefetch(true)}
        onFocus={() => setPrefetch(true)}
        onTouchStart={() => setPrefetch(true)}
        onNavigate={(event) => {
            if (linkRef.current?.querySelector('[data-navigation-pending="true"]')) {
                event.preventDefault();
                return;
            }
            if (pathname === href) event.preventDefault();
            onNavigate?.();
        }}
    >
        <NavigationContent icon={icon}>{children}</NavigationContent>
    </Link>;
}

function NavigationContent({ icon: Icon, children }: { icon?: LucideIcon; children: ReactNode }) {
    const { pending } = useLinkStatus();
    return <span className="contents" data-navigation-pending={pending} aria-busy={pending}>
        {Icon ? pending ? <LoaderCircle className="size-4 shrink-0 animate-spin motion-reduce:animate-none" aria-hidden /> : <Icon className="size-4 shrink-0" aria-hidden /> : null}
        {children}
        {pending ? createPortal(
            <div role="status" className="pointer-events-none fixed inset-x-0 top-20 z-[1100] flex justify-center">
                <span className="flex items-center gap-2 rounded-md bg-background/95 px-4 py-2 text-sm text-foreground">
                    <LoaderCircle className="size-4 animate-spin motion-reduce:animate-none" aria-hidden />正在打开页面…
                </span>
            </div>, document.body,
        ) : null}
    </span>;
}
