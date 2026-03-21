import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { landingContent } from '@/lib/content';

export function FeaturesSection() {
  const { features } = landingContent;

  return (
    <section aria-labelledby="features-heading">
      <h2 id="features-heading" className="text-3xl font-bold tracking-tight text-center mb-8">
        Features
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {features.map((feature) => (
          <article key={feature.title}>
            <Card className="h-full">
              <CardHeader>
                <CardTitle className="text-lg">{feature.title}</CardTitle>
                <CardDescription>{feature.description}</CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground">{feature.details}</p>
                <div className="mt-4 flex gap-2 flex-wrap">
                  {feature.badges.map((badge) => (
                    <Badge key={badge} variant="secondary">
                      {badge}
                    </Badge>
                  ))}
                </div>
              </CardContent>
            </Card>
          </article>
        ))}
      </div>
    </section>
  );
}
