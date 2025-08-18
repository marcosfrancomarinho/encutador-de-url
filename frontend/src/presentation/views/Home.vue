<template>
    <section class="bg-gray-50 min-h-[80vh] flex items-center">
        <div class="container mx-auto px-4 text-center">
            <!-- Título -->
            <h2 class="text-4xl md:text-5xl font-bold text-gray-800 mb-4">
                Bem-vindo ao <span class="text-blue-600">Short Url</span>
            </h2>

            <!-- Subtítulo -->
            <p class="text-lg text-gray-600 mb-8 max-w-2xl mx-auto">
                Encurte links longos de forma rápida, segura e elegante.
                Uma solução simples para compartilhar URLs sem complicação.
            </p>

            <!-- Botões -->
            <div class="flex flex-col sm:flex-row gap-4 justify-center">
                <router-link to="/short"
                    class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg text-lg transition">
                    Começar Agora
                </router-link>
                <a target="_blank" href="https://github.com/marcosfrancomarinho/encutador-de-url"
                    class="border border-blue-600 text-blue-600 hover:bg-blue-50 px-6 py-3 rounded-lg text-lg transition">
                    Saiba Mais
                </a>
            </div>

        </div>
    </section>
</template>


