<template>
  <footer class="bg-gray-900 text-white mt-8">
    <div class="container mx-auto flex flex-col md:flex-row items-center justify-between px-4 py-4">
      <!-- Texto -->
      <p class="text-sm">&copy; 2025 Short Url. Todos os direitos reservados.</p>

      <!-- Links -->
      <ul class="flex space-x-6 mt-2 md:mt-0">
        <li>
          <router-link to="/" class="hover:text-blue-400 text-xs transition-colors">Home</router-link>
        </li>
        <li>
          <router-link to="/short" class="hover:text-blue-400 text-xs transition-colors">Short</router-link>
        </li>
      </ul>
    </div>
  </footer>
</template>

