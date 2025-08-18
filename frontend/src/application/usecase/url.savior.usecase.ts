import { URL } from '../../domain/entities/url';
import type { HttpPostClient } from '../../domain/gateway/http.post.client';
import type { RequestUrlSaviorDTO } from '../dto/request.url.savior.dto';
import type { ResponseUrlSaviorDTO } from '../dto/response.url.savior.dto';

export class UrlSaviorUseCase {
  private httpPostClient: HttpPostClient;
  public constructor(httpPostClient: HttpPostClient) {
    this.httpPostClient = httpPostClient;
  }

  public async save(payload: RequestUrlSaviorDTO): Promise<ResponseUrlSaviorDTO> {
    const url: URL = new URL(payload.longUrl);
    console.log(url)
    const response = await this.httpPostClient.execute<ResponseUrlSaviorDTO>(url);
    return response;
  }
}
