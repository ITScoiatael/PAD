<template>
  <header
    class="bg-black flex items-center justify-between flex-wrap bg-teal p-2 shadow-md font-medium"
  >
    <OurLogo />
    <nav class>
      <div class="block lg:hidden">
        <div
          class="flex items-center px-3 py-2 border rounded text-teal-lighter border-teal-light hover:text-white hover:border-white"
        >
          <svg
            class="h-3 w-3"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <title>Menu</title>
            <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
          </svg>
        </div>
      </div>
      <div
        class="w-full block flex-grow lg:flex lg:w-auto font-gilroy font-medium"
      >
        <div class="text-sm lg:flex-grow text-center text-white uppercase">
          <NuxtLink
            to="/"
            class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter hover:text-gray-500 mx-12"
            >Home</NuxtLink
          >
          <NuxtLink
            to="/about"
            class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter hover:text-gray-500 mx-12"
            >About us</NuxtLink
          >
          <NuxtLink
            to="/contact"
            class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter hover:text-gray-500 mx-12"
            >contact us</NuxtLink
          >
          <NuxtLink
            to="/login"
            class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter hover:text-gray-500 mx-12"
            >login</NuxtLink
          >
        </div>
      </div>
    </nav>
  </header>
</template>