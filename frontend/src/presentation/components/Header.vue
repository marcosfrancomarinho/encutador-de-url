<template>
  <header class="bg-gray-900 text-white shadow-md">
    <div class="container mx-auto flex items-center justify-between px-4 py-3">
      <!-- Logo / Nome -->
      <h1 class="text-xl font-bold">Short Url</h1>

      <!-- Menu -->
      <nav>
        <ul class="flex space-x-6">
          <li>
            <router-link to="/" class="hover:text-blue-400 cursor-pointer transition-colors">
              Home
            </router-link>
          </li>
          <li>
            <router-link  to="short" class="hover:text-blue-400 cursor-pointer transition-colors">
              Short
            </router-link>
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>
