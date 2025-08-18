<template>
  <section class="container mx-auto px-4 py-10 text-center">
    <h2 class="text-3xl font-bold text-gray-800 mb-4">Bem-vindo ao Short Url</h2>
    <p class="text-gray-600 mb-8">
      Encurte seus links de forma rápida, segura e eficiente.
    </p>

    <div class="max-w-xl mx-auto flex flex-col md:flex-row items-center gap-4">
      <input
        type="text"
        placeholder="Cole sua URL aqui..."
        v-model="url"
        ref="urlInput"
        class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
        @click="handlerClick"
        :disabled="loading"
        class="flex items-center justify-center gap-2 bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg transition-colors disabled:bg-gray-400 disabled:cursor-not-allowed"
      >
        <svg
          v-if="loading"
          class="animate-spin h-5 w-5 text-white"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8v4l3-3-3-3v4a8 8 0 00-8 8h4l-3 3 3 3H4z"
          ></path>
        </svg>
        {{ loading ? 'Encurtando...' : 'Encurtar' }}
      </button>
    </div>

    <!-- Exibir resultado da URL encurtada -->
    <div v-if="datas?.shortUrl" class="mt-8 text-center">
      <p class="text-gray-700 mb-2">Sua URL encurtada:</p>
      <a
        :href="datas.shortUrl"
        target="_blank"
        class="text-blue-600 hover:underline text-lg font-semibold"
      >
        {{ datas.shortUrl }}
      </a>

      <!-- Espaço para QR Code -->
      <div class="mt-6 flex justify-center">
        <img
          :src="datas.qrCode"
          alt="QR Code"
          class="w-[200px] h-[200px] border rounded-lg shadow-md"
        />
      </div>
    </div>

    <!-- Exibir erro -->
    <div v-if="error?.message" class="mt-4 text-red-600 text-center">
      {{ error.message }}
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUrlSavior } from '../composables/use.url.savior.ts';

const url = ref('');
const urlInput = ref<HTMLInputElement | null>(null);
const { datas, urlSavior, error, loading } = useUrlSavior();

const handlerClick = () => {
  urlSavior(url.value);
  url.value = '';
  urlInput.value?.focus();
};
</script>
