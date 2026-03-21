/**
 * @jest-environment jsdom
 */
import '@testing-library/jest-dom';
import { render } from '@testing-library/react';

import {
  FAQJsonLd,
  OrganizationJsonLd,
  SoftwareApplicationJsonLd,
  WebPageJsonLd,
  WebSiteJsonLd,
} from '../structured-data';

function getJsonLd(): Record<string, unknown> {
  const script = document.querySelector('script[type="application/ld+json"]');
  return JSON.parse(script?.textContent || '{}');
}

describe('Structured Data Components', () => {
  afterEach(() => {
    document.head.innerHTML = '';
    document.body.innerHTML = '';
  });

  it('OrganizationJsonLd renders correct schema', () => {
    render(<OrganizationJsonLd name="Test Org" url="https://example.com" />);
    const data = getJsonLd();
    expect(data['@context']).toBe('https://schema.org');
    expect(data['@type']).toBe('Organization');
    expect(data.name).toBe('Test Org');
  });

  it('WebSiteJsonLd renders with SearchAction', () => {
    render(<WebSiteJsonLd name="Test Site" url="https://example.com" />);
    const data = getJsonLd();
    expect(data['@type']).toBe('WebSite');
    expect(data.potentialAction).toBeDefined();
  });

  it('WebPageJsonLd renders correct schema', () => {
    render(<WebPageJsonLd name="Home" description="A test page" url="https://example.com" />);
    const data = getJsonLd();
    expect(data['@type']).toBe('WebPage');
    expect(data.description).toBe('A test page');
  });

  it('FAQJsonLd renders question-answer pairs', () => {
    render(
      <FAQJsonLd
        items={[
          { question: 'Q1?', answer: 'A1' },
          { question: 'Q2?', answer: 'A2' },
        ]}
      />
    );
    const data = getJsonLd();
    expect(data['@type']).toBe('FAQPage');
    const entities = data.mainEntity as Array<Record<string, unknown>>;
    expect(entities).toHaveLength(2);
  });

  it('SoftwareApplicationJsonLd renders correct schema', () => {
    render(
      <SoftwareApplicationJsonLd
        name="Test App"
        description="A test app"
        url="https://example.com"
      />
    );
    const data = getJsonLd();
    expect(data['@type']).toBe('SoftwareApplication');
  });
});
