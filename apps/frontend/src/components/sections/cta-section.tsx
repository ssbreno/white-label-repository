import Link from 'next/link';

import { Button } from '@/components/ui/button';

export function CTASection() {
  return (
    <section
      className="text-center py-12 bg-muted/50 rounded-lg px-8"
      aria-labelledby="cta-heading"
    >
      <h2 id="cta-heading" className="text-3xl font-bold tracking-tight">
        Ready to Build?
      </h2>
      <p className="text-lg text-muted-foreground mt-4 max-w-xl mx-auto">
        Fork the repository, customize the branding, and start building your product in minutes.
      </p>
      <div className="flex gap-3 justify-center mt-6">
        <Button size="lg" asChild>
          <Link
            href="https://github.com/ssbreno/white-label-repository/fork"
            target="_blank"
            rel="noopener noreferrer"
          >
            Fork on GitHub
          </Link>
        </Button>
        <Button size="lg" variant="outline" asChild>
          <Link
            href="https://github.com/ssbreno/white-label-repository/blob/master/docs/CUSTOMIZATION.md"
            target="_blank"
            rel="noopener noreferrer"
          >
            Customization Guide
          </Link>
        </Button>
      </div>
    </section>
  );
}
