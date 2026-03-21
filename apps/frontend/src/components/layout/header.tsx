import Link from 'next/link';

import { ThemeToggle } from '@/components/theme-toggle';
import { siteConfig } from '@/lib/site-config';

export function Header() {
  return (
    <header className="border-b">
      <div className="container mx-auto px-4 py-4 flex items-center justify-between">
        <Link href="/" className="flex flex-col">
          <span className="text-2xl font-bold text-primary">{siteConfig.name}</span>
          <span className="text-sm text-muted-foreground">Nx · Go · Next.js · shadcn/ui</span>
        </Link>
        <nav aria-label="Main navigation" className="flex items-center gap-4">
          <Link
            href={`https://github.com/ssbreno/white-label-repository`}
            className="text-sm text-muted-foreground hover:text-foreground transition-colors"
            target="_blank"
            rel="noopener noreferrer"
          >
            GitHub
          </Link>
          <ThemeToggle />
        </nav>
      </div>
    </header>
  );
}
