import { createApp } from 'vue';
import './presentation/styles/global.css';
import App from './presentation/App.vue';
import { router } from './presentation/router/index';
import { Container } from './shared/container/container';

function main() {
  const { buildurlSaviorUseCase } = Container.getInstance();
  const app = createApp(App);
  app.provide('url_savior', buildurlSaviorUseCase());
  app.use(router);
  app.mount('#app');
}

main();
