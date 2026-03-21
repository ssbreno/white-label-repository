import { FAQJsonLd } from '@/components/structured-data';
import { landingContent } from '@/lib/content';

export function FAQSection() {
  const { faq } = landingContent;

  return (
    <section aria-labelledby="faq-heading">
      <FAQJsonLd items={faq} />
      <h2 id="faq-heading" className="text-3xl font-bold tracking-tight text-center mb-8">
        Frequently Asked Questions
      </h2>
      <div className="max-w-3xl mx-auto space-y-4">
        {faq.map((item) => (
          <details key={item.question} className="group border rounded-lg">
            <summary className="flex cursor-pointer items-center justify-between p-4 font-medium hover:bg-muted/50 transition-colors">
              <h3 className="text-left">{item.question}</h3>
              <span className="ml-4 shrink-0 text-muted-foreground group-open:rotate-180 transition-transform">
                &#9660;
              </span>
            </summary>
            <div className="px-4 pb-4">
              <p className="text-muted-foreground">{item.answer}</p>
            </div>
          </details>
        ))}
      </div>
    </section>
  );
}
