export class URL {
  private longUrl: string;
  public constructor(longUrl: string) {
    this.longUrl = longUrl;
  }
  public static create(longUrl: string): URL {
    const trimmed = this.validate(longUrl);
    return new URL(trimmed);
  }
  public getLongUrl(): string {
    return this.longUrl;
  }

  private static validate(longUrl: string): string {
    if (!longUrl) throw new Error('URL não foi informada.');
    if (typeof longUrl !== 'string') throw new Error('URL precisa ser string.');
    const trimmed: string = longUrl.trim();
    if (trimmed.length === 0) throw new Error('URL está vazia.');
    try {
      new globalThis.URL(trimmed);
    } catch {
      throw new Error('URL com formato inválido.');
    }
    return trimmed;
  }
}
