<template>
  <div>
    <nav
      class="bg-black flex items-center justify-between flex-wrap bg-teal p-6 shadow-md font-medium">
      <div class="flex items-center flex-no-shrink text-white mr-6">
        <OurLogo />
      </div>
      <div class="block lg:hidden">
        <button
          class="flex items-center px-3 py-2 border rounded text-teal-lighter border-teal-light hover:text-white hover:border-white">
          <svg class="h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
            <title>Menu</title>
            <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
          </svg>
        </button>
      </div>
      <div class="w-full block flex-grow lg:flex lg:w-auto font-gilroy font-bold">
        <div class="text-sm lg:flex-grow text-center">
          <a
            href="/"
            class="text-white block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-gray-500 mx-12"
          >Home</a>
          <a
            href="#responsive-header"
            class="text-white block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-gray-500 mx-12"
          >About us</a>
          <a
            href="#responsive-header"
            class="text-white block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-gray-500 mx-12"
          >contact us</a>
        </div>
      </div>
    </nav>

  </div>
</template>