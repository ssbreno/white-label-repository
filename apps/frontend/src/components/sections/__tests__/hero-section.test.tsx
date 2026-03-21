import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import React from 'react';

import { HeroSection } from '../hero-section';

jest.mock('next/link', () => {
  return {
    __esModule: true,
    default: ({ children, href }: { children: React.ReactNode; href: string }) => {
      return React.createElement('a', { href }, children);
    },
  };
});

jest.mock('@/components/ui/button', () => ({
  Button: ({ children }: { children: React.ReactNode }) => {
    return React.createElement('button', null, children);
  },
}));

jest.mock('@/lib/content', () => ({
  landingContent: {
    hero: {
      title: 'Test Hero Title',
      subtitle: 'Test subtitle text',
      ctaPrimary: 'Get Started',
      ctaPrimaryHref: '/start',
      ctaSecondary: 'Docs',
      ctaSecondaryHref: '/docs',
    },
  },
}));

describe('HeroSection', () => {
  it('renders the hero heading as h1', () => {
    render(React.createElement(HeroSection));
    const heading = screen.getByRole('heading', { level: 1 });
    expect(heading).toHaveTextContent('Test Hero Title');
  });

  it('renders subtitle text', () => {
    render(React.createElement(HeroSection));
    expect(screen.getByText('Test subtitle text')).toBeInTheDocument();
  });

  it('renders CTA buttons', () => {
    render(React.createElement(HeroSection));
    expect(screen.getByText('Get Started')).toBeInTheDocument();
    expect(screen.getByText('Docs')).toBeInTheDocument();
  });
});
