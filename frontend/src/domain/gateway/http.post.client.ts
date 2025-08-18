import type { URL } from '../entities/url';

export interface HttpPostClient {
  execute<T>(url: URL): Promise<T>;
}
