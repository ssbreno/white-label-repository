import Link from 'next/link';

import { Button } from '@/components/ui/button';
import { landingContent } from '@/lib/content';

export function HeroSection() {
  const { hero } = landingContent;

  return (
    <section className="text-center space-y-4 py-16" aria-labelledby="hero-heading">
      <h1 id="hero-heading" className="text-4xl md:text-5xl font-bold tracking-tight">
        {hero.title}
      </h1>
      <p className="text-xl text-muted-foreground max-w-2xl mx-auto">{hero.subtitle}</p>
      <div className="flex gap-3 justify-center pt-4">
        <Button size="lg" asChild>
          <Link href={hero.ctaPrimaryHref} target="_blank" rel="noopener noreferrer">
            {hero.ctaPrimary}
          </Link>
        </Button>
        <Button size="lg" variant="outline" asChild>
          <Link href={hero.ctaSecondaryHref}>{hero.ctaSecondary}</Link>
        </Button>
      </div>
    </section>
  );
}
