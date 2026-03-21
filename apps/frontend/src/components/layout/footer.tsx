import Link from 'next/link';

import { siteConfig } from '@/lib/site-config';

export function Footer() {
  return (
    <footer className="border-t mt-16" role="contentinfo">
      <div className="container mx-auto px-4 py-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <p className="font-semibold text-foreground">{siteConfig.name}</p>
            <p className="text-sm text-muted-foreground mt-2">{siteConfig.description}</p>
          </div>
          <nav aria-label="Documentation links">
            <p className="font-semibold text-foreground">Documentation</p>
            <ul className="mt-2 space-y-1">
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository/blob/master/docs/ARCHITECTURE.md"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Architecture
                </Link>
              </li>
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository/blob/master/docs/CUSTOMIZATION.md"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Customization
                </Link>
              </li>
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository/blob/master/docs/DEPLOYMENT.md"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Deployment
                </Link>
              </li>
            </ul>
          </nav>
          <nav aria-label="Project links">
            <p className="font-semibold text-foreground">Project</p>
            <ul className="mt-2 space-y-1">
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  GitHub
                </Link>
              </li>
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository/blob/master/CONTRIBUTING.md"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Contributing
                </Link>
              </li>
              <li>
                <Link
                  href="https://github.com/ssbreno/white-label-repository/blob/master/LICENSE"
                  className="text-sm text-muted-foreground hover:text-foreground transition-colors"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  License (GPL-3.0)
                </Link>
              </li>
            </ul>
          </nav>
        </div>
        <div className="border-t mt-8 pt-4 text-center">
          <p className="text-sm text-muted-foreground">
            &copy; {new Date().getFullYear()} {siteConfig.name}. Open source under GPL-3.0.
          </p>
        </div>
      </div>
    </footer>
  );
}
