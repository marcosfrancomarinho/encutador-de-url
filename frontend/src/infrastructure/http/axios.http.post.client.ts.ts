import type { AxiosInstance } from 'axios';
import type { URL } from '../../domain/entities/url';
import type { HttpPostClient } from '../../domain/gateway/http.post.client';
import axios from 'axios';

export class AxiosHttpPostClient implements HttpPostClient {
  private instance: AxiosInstance;
  private path: string;
  public constructor(path: string) {
    this.instance = axios;
    this.path = path;
  }
  public async execute<T>(url: URL): Promise<T> {
    const headers = { 'Content-Type': 'application/json' };
    const body = { long_url: url.getLongUrl() };
    try {
      const response = await this.instance.post(this.path, body, { headers });
      return { shortUrl: response.data.url_short, qrCode: response.data.qr_code } as T;
    } catch (error: any) {
      throw new Error(error.response.data.error);
    }
  }
}
