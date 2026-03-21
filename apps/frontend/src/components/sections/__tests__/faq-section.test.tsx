/**
 * @jest-environment jsdom
 */
import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';

jest.mock('@/lib/content', () => ({
  landingContent: {
    faq: [
      { question: 'What is this?', answer: 'A template.' },
      { question: 'How to use?', answer: 'Fork and customize.' },
    ],
  },
}));

const { FAQSection } = require('../faq-section');

describe('FAQSection', () => {
  it('renders the section heading', () => {
    render(<FAQSection />);
    expect(screen.getByText('Frequently Asked Questions')).toBeInTheDocument();
  });

  it('renders all FAQ questions', () => {
    render(<FAQSection />);
    expect(screen.getByText('What is this?')).toBeInTheDocument();
    expect(screen.getByText('How to use?')).toBeInTheDocument();
  });

  it('renders FAQ answers', () => {
    render(<FAQSection />);
    expect(screen.getByText('A template.')).toBeInTheDocument();
    expect(screen.getByText('Fork and customize.')).toBeInTheDocument();
  });

  it('includes FAQ JSON-LD script', () => {
    render(<FAQSection />);
    const scripts = document.querySelectorAll('script[type="application/ld+json"]');
    expect(scripts.length).toBeGreaterThan(0);
    const jsonLd = JSON.parse(scripts[0].textContent || '');
    expect(jsonLd['@type']).toBe('FAQPage');
    expect(jsonLd.mainEntity).toHaveLength(2);
  });
});
