/**
 * @jest-environment jsdom
 */
import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';

jest.mock('@/components/ui/card', () => ({
  Card: ({ children, className }: { children: React.ReactNode; className?: string }) => (
    <div className={className}>{children}</div>
  ),
  CardHeader: ({ children }: { children: React.ReactNode }) => <div>{children}</div>,
  CardTitle: ({ children, className }: { children: React.ReactNode; className?: string }) => (
    <h3 className={className}>{children}</h3>
  ),
  CardDescription: ({ children }: { children: React.ReactNode }) => <p>{children}</p>,
  CardContent: ({ children }: { children: React.ReactNode }) => <div>{children}</div>,
}));

jest.mock('@/components/ui/badge', () => ({
  Badge: ({ children }: { children: React.ReactNode }) => <span>{children}</span>,
}));

jest.mock('@/lib/content', () => ({
  landingContent: {
    features: [
      {
        title: 'Feature One',
        description: 'Description one',
        details: 'Details one',
        badges: ['Badge1', 'Badge2'],
      },
      {
        title: 'Feature Two',
        description: 'Description two',
        details: 'Details two',
        badges: ['Badge3'],
      },
    ],
  },
}));

const { FeaturesSection } = require('../features-section');

describe('FeaturesSection', () => {
  it('renders the section heading as h2', () => {
    render(<FeaturesSection />);
    const heading = screen.getByRole('heading', { level: 2 });
    expect(heading).toHaveTextContent('Features');
  });

  it('renders all feature cards', () => {
    render(<FeaturesSection />);
    expect(screen.getByText('Feature One')).toBeInTheDocument();
    expect(screen.getByText('Feature Two')).toBeInTheDocument();
  });

  it('renders feature badges', () => {
    render(<FeaturesSection />);
    expect(screen.getByText('Badge1')).toBeInTheDocument();
    expect(screen.getByText('Badge3')).toBeInTheDocument();
  });

  it('wraps feature cards in article elements', () => {
    render(<FeaturesSection />);
    const articles = document.querySelectorAll('article');
    expect(articles).toHaveLength(2);
  });
});
