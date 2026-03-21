import { siteConfig } from '../site-config';

describe('siteConfig', () => {
  it('has default name', () => {
    expect(siteConfig.name).toBeTruthy();
    expect(typeof siteConfig.name).toBe('string');
  });

  it('has default description', () => {
    expect(siteConfig.description).toBeTruthy();
    expect(siteConfig.description.length).toBeGreaterThan(10);
  });

  it('has default url', () => {
    expect(siteConfig.url).toBeTruthy();
    expect(siteConfig.url).toContain('http');
  });

  it('has keywords as array', () => {
    expect(Array.isArray(siteConfig.keywords)).toBe(true);
    expect(siteConfig.keywords.length).toBeGreaterThan(0);
  });

  it('has locale', () => {
    expect(siteConfig.locale).toMatch(/^[a-z]{2}_[A-Z]{2}$/);
  });

  it('has themeColor as hex', () => {
    expect(siteConfig.themeColor).toMatch(/^#[0-9a-fA-F]{6}$/);
  });
});
