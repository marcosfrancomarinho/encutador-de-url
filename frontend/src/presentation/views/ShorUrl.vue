<template>
  <section class="container mx-auto px-4 py-10 text-center">
    <h2 class="text-3xl font-bold text-gray-800 mb-4">Bem-vindo ao Short Url</h2>
    <p class="text-gray-600 mb-8">
      Encurte seus links de forma rápida, segura e eficiente.
    </p>

    <!-- Campo de entrada + botão -->
    <div class="max-w-xl mx-auto flex flex-col md:flex-row items-center gap-4">
      <input type="text" placeholder="Cole sua URL aqui..." v-model="url"
        class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      <button @click="handlerCick"
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg transition-colors">
        Encurtar
      </button>
    </div>

    <!-- Exibir resultado -->
    <div v-if="datas?.shortUrl" class="mt-8 text-center">
      <p class="text-gray-700 mb-2">Sua URL encurtada:</p>
      <a :href="datas.shortUrl" target="_blank" class="text-blue-600 hover:underline text-lg font-semibold">
        {{ datas.shortUrl }}
      </a>

      <!-- Espaço para o QR Code -->
      <div class="mt-6 flex justify-center">
        <img :src="datas.qrCode" alt="QR Code" class="w-50 h-50 border rounded-lg shadow-md" />
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUrlSavior } from '../composables/use.url.savior.ts';

const url = ref('');
const { datas, urlSavior } = useUrlSavior();
const handlerCick =  () => urlSavior(url.value);

</script>
