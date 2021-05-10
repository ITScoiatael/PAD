<template>
  <!-- <nav class="bg-primary py-2 md:py-4">
        <div class="container px-4 mx-auto md:flex md:items-center">
            <OurLogo/>
            <div class="hidden md:flex flex-col md:flex-row md:ml-auto mt-3 md:mt-0" id="navbar-collapse">
                <div class="content-around ">
                    <NuxtLink  to="./"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">Home</NuxtLink >
                    <a href="https://github.com/ITScoiatael/PAD" target="_blank"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">About
                        us</a>
                    <NuxtLink  to="Catalog"
                        class="p-2 lg:px-4 md:mx-2 text-black-900 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300">Catalog</NuxtLink>
                </div>
            </div>


        </div>
  </nav>-->

  <nav class="flex items-center justify-between flex-wrap bg-teal p-6 bg-primary shadow-md font-medium">
    <div class="flex items-center flex-no-shrink text-white mr-6">
      <OurLogo />
    </div>
    <div class="block lg:hidden">
      <button
        class="flex items-center px-3 py-2 border rounded text-teal-lighter border-teal-light hover:text-white hover:border-white"
      >
        <svg class="h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
          <title>Menu</title>
          <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
        </svg>
      </button>
    </div>
    <div class="w-full block flex-grow lg:flex lg:w-auto">
      <div class="text-sm lg:flex-grow text-center">
        <a
          href="#responsive-header"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-white mr-4"
        >Home</a>
        <a
          href="#responsive-header"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-white mr-4"
        >About us</a>
        <a
          href="#responsive-header"
          class="block mt-4 lg:inline-block lg:mt-0 text-teal-lighter uppercase hover:text-white"
        >contact us</a>
      </div>
    </div>
  </nav>
</template>