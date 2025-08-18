export class URL {
  private longUrl: string;
  public constructor(longUrl: string) {
    const trimmed = this.validate(longUrl);
    this.longUrl = trimmed;
  }

  public getLongUrl(): string {
    return this.longUrl;
  }

  private validate(longUrl: string): string {
    if (!longUrl) throw new Error('url longa não foi informada');
    if (typeof longUrl !== 'string') throw new Error('url precisa ser string');
    const trimmed: string = longUrl.trim();
    if (trimmed.length === 0) throw new Error('url longa vazia');
    new globalThis.URL(trimmed)
    return trimmed;
  }
}
