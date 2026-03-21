'use client';

import { useEffect, useState } from 'react';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { api } from '@/lib/api';

import type { HealthStatus } from '@/types';

export function ApiHealthCheck() {
  const [health, setHealth] = useState<HealthStatus | null>(null);
  const [loading, setLoading] = useState(false);

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
    <aside aria-label="API Health Status">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            API Health
            {health && (
              <Badge variant={health.status === 'healthy' ? 'default' : 'destructive'}>
                {health.status}
              </Badge>
            )}
          </CardTitle>
          <CardDescription>Real-time backend connectivity status</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          {health ? (
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
              {Object.entries(health.services).map(([name, status]) => (
                <div key={name} className="text-center">
                  <p className="text-sm font-medium capitalize">{name}</p>
                  <Badge
                    variant={status === 'healthy' ? 'default' : 'destructive'}
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
          <Button onClick={checkHealth} disabled={loading} variant="outline" size="sm">
            {loading ? 'Checking...' : 'Refresh'}
          </Button>
        </CardContent>
      </Card>
    </aside>
  );
}
