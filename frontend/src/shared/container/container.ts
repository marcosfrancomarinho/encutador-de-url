import { UrlSaviorUseCase } from '../../application/usecase/url.savior.usecase';
import { AxiosHttpPostClient } from '../../infrastructure/http/axios.http.post.client.ts';

export class Container {
  public static instance: Container;
  
  public static getInstance(): Container {
    if (!this.instance) {
      this.instance = new Container();
    }
    return this.instance;
  }

  public buildurlSaviorUseCase() {
    const httpPostClient = new AxiosHttpPostClient('http://localhost:3030');
    return new UrlSaviorUseCase(httpPostClient);
  }
}
