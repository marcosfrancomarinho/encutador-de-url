import { ref, inject } from 'vue';
import type { UrlSaviorUseCase } from '../../application/usecase/url.savior.usecase';

export type Payload = {
  shortUrl: string;
  qrCode: string;
};

export const useUrlSavior = () => {
  const urlSaviorUseCase = inject('url_savior') as UrlSaviorUseCase 

  const error = ref<Error | null>(null);
  const loading = ref(false);
  const datas = ref<Payload | null>(null);

  const urlSavior = async (longUrl: string) => {
    try {
      loading.value = true;
      error.value = null;
      datas.value = null;
    
      const response = await urlSaviorUseCase.save({ longUrl });
      datas.value = { shortUrl: response.shortUrl, qrCode: response.qrCode };
    } catch (err) {
      if (err instanceof Error) error.value = err;
    } finally {
      loading.value = false;
    }
  };

  return { error, datas, loading, urlSavior };
};
