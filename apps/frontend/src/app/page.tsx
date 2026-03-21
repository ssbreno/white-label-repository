import type { Metadata } from 'next';

import { ApiHealthCheck } from '@/components/api-health-check';
import { Footer } from '@/components/layout/footer';
import { Header } from '@/components/layout/header';
import { CTASection } from '@/components/sections/cta-section';
import { FAQSection } from '@/components/sections/faq-section';
import { FeaturesSection } from '@/components/sections/features-section';
import { HeroSection } from '@/components/sections/hero-section';
import {
  OrganizationJsonLd,
  SoftwareApplicationJsonLd,
  WebPageJsonLd,
  WebSiteJsonLd,
} from '@/components/structured-data';
import { siteConfig } from '@/lib/site-config';

export const metadata: Metadata = {
  title: 'Home',
  description: siteConfig.description,
  alternates: { canonical: '/' },
};

export default function HomePage() {
  return (
    <>
      <OrganizationJsonLd name={siteConfig.name} url={siteConfig.url} />
      <WebSiteJsonLd name={siteConfig.name} url={siteConfig.url} />
      <WebPageJsonLd
        name={`${siteConfig.name} - Home`}
        description={siteConfig.description}
        url={siteConfig.url}
      />
      <SoftwareApplicationJsonLd
        name={siteConfig.name}
        description={siteConfig.description}
        url={siteConfig.url}
      />

      <Header />

      <main className="min-h-screen bg-background">
        <div className="container mx-auto px-4 py-8 space-y-16">
          <HeroSection />
          <FeaturesSection />
          <ApiHealthCheck />
          <CTASection />
          <FAQSection />
        </div>
      </main>

      <Footer />
    </>
  );
}
