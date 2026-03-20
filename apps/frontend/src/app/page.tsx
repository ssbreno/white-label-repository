'use client';

import { useEffect, useState } from 'react';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { ThemeToggle } from '@/components/theme-toggle';
import { Badge } from '@/components/ui/badge';
import { api } from '@/lib/api';

import type { HealthStatus } from '@/types';

export default function HomePage() {
  const [health, setHealth] = useState<HealthStatus | null>(null);
  const [loading, setLoading] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');

  const checkHealth = async () => {
    setLoading(true);
    try {
      const data = await api.checkHealth();
      setHealth(data);
    } catch {
      setHealth(null);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    checkHealth();
  }, []);

  return (
    <main className="min-h-screen bg-background">
      {/* Header */}
      <header className="border-b">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold text-primary">
              {process.env.NEXT_PUBLIC_APP_NAME || 'White Label App'}
            </h1>
            <p className="text-sm text-muted-foreground">
              Nx · Go · Next.js · shadcn/ui
            </p>
          </div>
          <ThemeToggle />
        </div>
      </header>

      <div className="container mx-auto px-4 py-8 space-y-8">
        {/* Hero section */}
        <section className="text-center space-y-4 py-12">
          <h2 className="text-4xl font-bold tracking-tight">
            White Label Monorepo
          </h2>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            A production-ready template with Go backend, Next.js frontend, and
            shared libraries — ready to customise for any project.
          </p>
          <div className="flex gap-3 justify-center">
            <Button size="lg">Get Started</Button>
            <Button size="lg" variant="outline">
              View Docs
            </Button>
          </div>
        </section>

        {/* Search example */}
        <section className="max-w-md mx-auto">
          <div className="flex gap-2">
            <Input
              placeholder="Search..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            <Button variant="secondary">Search</Button>
          </div>
        </section>

        {/* Feature cards */}
        <section className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Go Backend</CardTitle>
              <CardDescription>High-performance REST API</CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Built with Gin framework, PostgreSQL via pgx, structured
                logging, and graceful shutdown support.
              </p>
              <div className="mt-4 flex gap-2 flex-wrap">
                <Badge>Go 1.21+</Badge>
                <Badge variant="secondary">Gin</Badge>
                <Badge variant="secondary">PostgreSQL</Badge>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Next.js Frontend</CardTitle>
              <CardDescription>Modern React application</CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Next.js 14 with App Router, TypeScript, Tailwind CSS, and
                shadcn/ui component library with dark mode support.
              </p>
              <div className="mt-4 flex gap-2 flex-wrap">
                <Badge>Next.js 14</Badge>
                <Badge variant="secondary">TypeScript</Badge>
                <Badge variant="secondary">Tailwind</Badge>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Shared Libraries</CardTitle>
              <CardDescription>Reusable code across projects</CardDescription>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-muted-foreground">
                Shared TypeScript types, utility functions, and configuration
                that can be imported by any app in the monorepo.
              </p>
              <div className="mt-4 flex gap-2 flex-wrap">
                <Badge>Types</Badge>
                <Badge variant="secondary">Utils</Badge>
                <Badge variant="secondary">Config</Badge>
              </div>
            </CardContent>
          </Card>
        </section>

        {/* API Health Check */}
        <section>
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                API Health
                {health && (
                  <Badge
                    variant={
                      health.status === 'healthy' ? 'default' : 'destructive'
                    }
                  >
                    {health.status}
                  </Badge>
                )}
              </CardTitle>
              <CardDescription>
                Real-time backend connectivity status
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              {health ? (
                <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                  {Object.entries(health.services).map(([name, status]) => (
                    <div key={name} className="text-center">
                      <p className="text-sm font-medium capitalize">{name}</p>
                      <Badge
                        variant={
                          status === 'healthy' ? 'default' : 'destructive'
                        }
                        className="mt-1"
                      >
                        {status}
                      </Badge>
                    </div>
                  ))}
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  {loading
                    ? 'Checking API...'
                    : 'Backend unavailable. Start the backend server to see health status.'}
                </p>
              )}
              <Button
                onClick={checkHealth}
                disabled={loading}
                variant="outline"
                size="sm"
              >
                {loading ? 'Checking...' : 'Refresh'}
              </Button>
            </CardContent>
          </Card>
        </section>
      </div>
    </main>
  );
}
